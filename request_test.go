package main

import (
	"bufio"
	"fmt"
	"net/url"
	"strings"
	"testing"
)

func TestParseRequestLine(t *testing.T) {
	requestLine := "GET /todos HTTP/1.1"
	expectedMethod := "GET"
	expectedRequestURI := "/todos"
	expectedProto := "HTTP/1.1"
	if actualMethod, actualRequestURI, actualProto, _ := parseRequestLine(requestLine);
		!(actualMethod == expectedMethod && actualRequestURI == expectedRequestURI && expectedProto == actualProto) {
		t.Fatalf("Unexpected result:\nGot:\t%s, %s, %s\nExpected:\t%s, %s, %s\n",
			actualMethod, actualRequestURI, actualProto,
			expectedMethod, expectedRequestURI, expectedProto)
	}
	return
}

type reqTest struct {
	Raw     string
	Req     *Request
	Error   string
}

var noError = ""

var reqTests = []reqTest{
	// Baseline test; All Request fields included for template use
	{
		"GET http://www.techcrunch.com/ HTTP/1.1\r\n" +
			"Host: www.techcrunch.com\r\n" +
			"User-Agent: Fake\r\n" +
			"Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8\r\n" +
			"Accept-Language: en-us,en;q=0.5\r\n" +
			"Accept-Encoding: gzip,deflate\r\n" +
			"Accept-Charset: ISO-8859-1,utf-8;q=0.7,*;q=0.7\r\n" +
			"Keep-Alive: 300\r\n" +
			"Content-Length: 7\r\n" +
			"Proxy-Connection: keep-alive\r\n\r\n" +
			"abcdef\n",

		&Request{
			Method: "GET",
			URL: &url.URL{
				Scheme: "http",
				Host:   "www.techcrunch.com",
				Path:   "/",
			},
			Proto:      "HTTP/1.1",
			Header: Header{
				"Accept":           {"text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8"},
				"Accept-Language":  {"en-us,en;q=0.5"},
				"Accept-Encoding":  {"gzip,deflate"},
				"Accept-Charset":   {"ISO-8859-1,utf-8;q=0.7,*;q=0.7"},
				"Keep-Alive":       {"300"},
				"Proxy-Connection": {"keep-alive"},
				"Content-Length":   {"7"},
				"Host":				{"www.techcrunch.com"},
				"User-Agent":       {"Fake"},
			},
			Body: "abcdef\n",
			RequestURI:    "http://www.techcrunch.com/",
		},
		noError,
	},

	// GET request with no body (the normal case)
	{
		"GET / HTTP/1.1\r\n" +
			"Host: foo.com\r\n\r\n",

		&Request{
			Method: "GET",
			URL: &url.URL{
				Path: "/",
			},
			Proto:         "HTTP/1.1",
			Header:        Header{
				"Host":				{"foo.com"},
			},
			Body:		   "",
			RequestURI:    "/",
		},
		noError,
	},

	// Tests that we don't parse a path that looks like a
	// scheme-relative URI as a scheme-relative URI.
	{
		"GET //user@host/is/actually/a/path/ HTTP/1.1\r\n" +
			"Host: test\r\n\r\n",

		&Request{
			Method: "GET",
			URL: &url.URL{
				Path: "//user@host/is/actually/a/path/",
			},
			Proto:         "HTTP/1.1",
			Header:        Header{
				"Host":				{"test"},
			},
			Body:		   "",
			RequestURI:    "//user@host/is/actually/a/path/",
		},
		noError,
	},

	// Tests a bogus absolute-path on the Request-Line (RFC 7230 section 5.3.1)
	{
		"GET ../../../../etc/passwd HTTP/1.1\r\n" +
			"Host: test\r\n\r\n",
		nil,
		`parse "../../../../etc/passwd": invalid URI for request`,
	},

	// Tests missing URL:
	{
		"GET  HTTP/1.1\r\n" +
			"Host: test\r\n\r\n",
		nil,
		`parse "": empty url`,
	},
}

func TestReadRequest(t *testing.T) {
	for i := range reqTests {
		tt := &reqTests[i]
		req, err := readRequest(bufio.NewReader(strings.NewReader(tt.Raw)))
		if err != nil {
			if err.Error() != tt.Error {
				t.Errorf("#%d: error %q, want error %q", i, err.Error(), tt.Error)
			}
			continue
		}
		testName := fmt.Sprintf("Test %d (%q)", i, tt.Raw)
		diff(t, testName, req, tt.Req)
		if req.Body != tt.Req.Body {
			t.Errorf("%s: Body = %q want %q", testName, req.Body, tt.Req.Body)
		}
	}
}
