package main

import (
	"fmt"
	"strings"
)

var NoBody = ""

type MyMux struct {}

func (p *MyMux) ServeHTTP(w *Response, r *Request) string {
	readTransfer(w, *r)

	switch r.URL.Path {
	case "/head_request":
		HeadRequest(w, r)
	case "/simple_get":
		SimpleGet(w, r)
	case "/simple_get_with_body":
		SimpleGetWithBody(w, r)
	case "/simple_head":
		SimpleHead(w, r)
	case "/method_options":
		MethodOptions(w, r)
	case "/method_options2":
		MethodOptions2(w, r)
	case "/redirect":
		Redirect(w, r)
	case "/echo_body":
		EchoBody(w, r)
	case "/text_response":
		TextResponse(w, r)
	case "/html_response":
		HTMLResponse(w, r)
	case "/json_response":
		JSONResponse(w, r)
	case "/xml_response":
		XMLResponse(w, r)
	default:
		NotFound(w, r)
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
