package main

import (
	"fmt"
	"regexp"
	"strings"
)

// NoBody is an empty string body
var NoBody = ""

type myMux struct {}

func (p *myMux) ServeHTTP(w *Response, r *Request) string {
	readTransfer(w, *r)

	matched, _ := regexp.MatchString(`todo/\d`, r.URL.Path)
	if matched {
		handleTodoResponse(w, r)
		return serveHTTP(w)
	}

	switch r.URL.Path {
	case "/head_request":
		headRequest(w, r)
	case "/simple_get":
		simpleGet(w, r)
	case "/simple_get_with_body":
		simpleGetWithBody(w, r)
	case "/simple_head":
		simpleHead(w, r)
	case "/method_options":
		methodOptions(w, r)
	case "/method_options2":
		methodOptions2(w, r)
	case "/redirect":
		redirect(w, r)
	case "/echo_body":
		echoBody(w, r)
	case "/text_response":
		textResponse(w, r)
	case "/html_response":
		htmlResponse(w, r)
	case "/json_response":
		jsonResponse(w, r)
	case "/xml_response":
		xmlResponse(w, r)
	case "/kitteh.jpg":
		kittehResponse(w, r)
	case "/doggo.png":
		doggoResponse(w, r)
	case "/kisses.gif":
		kissesResponse(w, r)
	case "/health-check.html":
		healthCheckResponse(w, r)
	case "/todo":
		todoResponse(w, r)
	default:
		notFound(w, r)
	}

	return serveHTTP(w)
}

func serveHTTP(w *Response) string {
	var headers string
	for key, val := range w.Header {
		 headers += fmt.Sprintf("%s: %s\r\n", key, strings.Join(val, ", "))
	}
	if strings.Compare(w.Body, NoBody) == 0 {
		return fmt.Sprintf("%s %d %s\r\n%s\r\n",
			w.Proto, w.StatusCode, w.Status, headers)
	}
	return fmt.Sprintf("%s %d %s\r\n%s\r\n%s",
		w.Proto, w.StatusCode, w.Status, headers, w.Body)
}
