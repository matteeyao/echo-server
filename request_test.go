package main

import (
	"bufio"
	"fmt"
	"go/token"
	"net/url"
	"reflect"
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
		"GET / HTTP/1.1\r\n" +
			"Host: localhost.com\r\n" +
			"Accept: text/html\r\n" +
			"Content-Length: 7\r\n\r\n" +
			"abcdef\n",

		&Request{
			Method: "GET",
			URL: &url.URL{
				Path:   "/",
			},
			Proto:      "HTTP/1.1",
			Header: Header{
				"Accept":           {"text/html"},
				"Content-Length":   {"7"},
				"Host":				{"localhost.com"},
			},
			Body: "abcdef\n",
			RequestURI:    "/",
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

func TestConvertRawRequestStringToRequestStruct(t *testing.T) {
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
		testComparisonOfRequests(t, testName, req, tt.Req)
		if req.Body != tt.Req.Body {
			t.Errorf("%s: Body = %q want %q", testName, req.Body, tt.Req.Body)
		}
	}
}

func testComparisonOfRequests(t *testing.T, prefix string, have, want interface{}) {
	t.Helper()
	hv := reflect.ValueOf(have).Elem()
	wv := reflect.ValueOf(want).Elem()
	if hv.Type() != wv.Type() {
		t.Errorf("%s: type mismatch %v want %v", prefix, hv.Type(), wv.Type())
	}
	for i := 0; i < hv.NumField(); i++ {
		name := hv.Type().Field(i).Name
		if !token.IsExported(name) {
			continue
		}
		hf := hv.Field(i).Interface()
		wf := wv.Field(i).Interface()
		if !reflect.DeepEqual(hf, wf) {
			t.Errorf("%s:\n\n%s (Actual) = %v\r\n%s (Expected) = %v", prefix, name, hf, name, wf)
		}
	}
}
