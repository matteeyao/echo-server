package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
)

type MyMux struct {
}

func (p *MyMux) ServeHTTP(r *Request) string {
	switch r.URL.Path {
	case "/head_request":
		switch r.Method {
		case "GET":
			return fmt.Sprintf(
				"HTTP/1.1 %d %s\r\nAllow: %s, %s\r\nConnection: Close\r\n\r\n%s",
				StatusMethodNotAllowed,
				StatusText(StatusMethodNotAllowed),
				"HEAD",
				"OPTIONS",
				"")
		case "HEAD":
			return fmt.Sprintf(
				"HTTP/1.1 %d %s\r\nConnection: Close\r\n\r\n%s",
				StatusOK,
				StatusText(StatusOK),
				"")
		}
	case "/simple_get":
		return fmt.Sprintf(
			"HTTP/1.1 %d %s\r\nConnection: Close\r\n\r\n%s",
			StatusOK,
			StatusText(StatusOK),
			"")
	case "/simple_get_with_body":
		return fmt.Sprintf(
			"HTTP/1.1 %d %s\r\nConnection: Close\r\n\r\n%s",
			StatusOK,
			StatusText(StatusOK),
			"Hello world")
	case "/simple_head":
		return fmt.Sprintf(
			"HTTP/1.1 %d %s\r\nConnection: Close\r\n\r\n%s",
			StatusOK,
			StatusText(StatusOK),
			"")
	case "/method_options":
		return fmt.Sprintf(
			"HTTP/1.1 %d %s\r\nAllow: %s\nConnection: Close\r\n\r\n%s",
			StatusOK,
			StatusText(StatusOK),
			"GET, HEAD, OPTIONS",
			"")
	case "/method_options2":
		return fmt.Sprintf(
			"HTTP/1.1 %d %s\r\nAllow: %s\nConnection: Close\r\n\r\n%s",
			StatusOK,
			StatusText(StatusOK),
			"GET, HEAD, OPTIONS, PUT, POST",
			"")
	case "/redirect":
		return fmt.Sprintf(
			"HTTP/1.1 %d %s\r\nLocation: %s\nConnection: Close\r\n\r\n%s",
			StatusMovedPermanently,
			StatusText(StatusMovedPermanently),
			"http://127.0.0.1:5000/simple_get",
			"")
	case "/echo_body":
		switch r.Method {
		case "POST":
			rbody := r.Body
			r.Body = nil
			var bout bytes.Buffer
			if rbody != nil {
				_, err := io.Copy(&bout, rbody)
				if err != nil {
					log.Fatalf("Request: copying body: %v", err)
				}
				rbody.Close()
			}

			return fmt.Sprintf(
				"HTTP/1.1 %d %s\r\nConnection: Close\r\n\r\n%s",
				StatusOK,
				StatusText(StatusOK),
				bout.String())
		}
	case "/text_response":
		return fmt.Sprintf(
			"HTTP/1.1 %d %s\r\nContent-Type: %s\nConnection: Close\r\n\r\n%s",
			StatusOK,
			StatusText(StatusOK),
			"text/plain;charset=utf-8",
			"text response")
	case "/html_response":
		return fmt.Sprintf(
			"HTTP/1.1 %d %s\r\nContent-Type: %s\nConnection: Close\r\n\r\n%s",
			StatusOK,
			StatusText(StatusOK),
			"text/html;charset=utf-8",
			"<html><body><p>HTML Response</p></body></html>")
	case "/json_response":
		return fmt.Sprintf(
			"HTTP/1.1 %d %s\r\nContent-Type: %s\nConnection: Close\r\n\r\n%s",
			StatusOK,
			StatusText(StatusOK),
			"application/json;charset=utf-8",
			"{\"key1\":\"value1\",\"key2\":\"value2\"}")
	case "/xml_response":
		return fmt.Sprintf(
			"HTTP/1.1 %d %s\r\nContent-Type: %s\nConnection: Close\r\n\r\n%s",
			StatusOK,
			StatusText(StatusOK),
			"application/xml;charset=utf-8",
			"<note><body>XML Response</body></note>")
	case "/kitteh.jpg":
		return fmt.Sprintf(
			"HTTP/1.1 %d %s\r\nContent-Type: %s\nConnection: Close\r\n\r\n%s",
			StatusOK,
			StatusText(StatusOK),
			"image/jpeg",
			"test body")
	case "/doggo.png":
		return fmt.Sprintf(
			"HTTP/1.1 %d %s\r\nContent-Type: %s\nConnection: Close\r\n\r\n%s",
			StatusOK,
			StatusText(StatusOK),
			"image/png",
			"test body")
	case "/kisses.gif":
		return fmt.Sprintf(
			"HTTP/1.1 %d %s\r\nContent-Type: %s\r\nConnection: Close\r\n\r\n%s",
			StatusOK,
			StatusText(StatusOK),
			"image/gif",
			"test body")
	case "/health-check.html":
		return fmt.Sprintf(
			"HTTP/1.1 %d %s\r\nContent-Type: %s\r\nConnection: Close\r\n\r\n%s",
			StatusOK,
			StatusText(StatusOK),
			"text/html;charset=utf-8",
			"<html><body><<strong>Status:</strong> pass</body></html>")
	case "/todo":
		rbody := r.Body
		r.Body = nil
		var bout bytes.Buffer
		if rbody != nil {
			_, err := io.Copy(&bout, rbody)
			if err != nil {
				log.Fatalf("Request: copying body: %v", err)
			}
			rbody.Close()
		}
		switch r.Method {
		case "POST":
			switch r.Header.Get("Content-Type") {
			case "text/xml; charset=utf-8":
				return fmt.Sprintf(
					"HTTP/1.1 %d %s\r\nConnection: Close\r\n\r\n%s",
					StatusUnsupportedMediaType,
					StatusText(StatusUnsupportedMediaType),
					"")
			case "application/x-www-form-urlencoded":
				return fmt.Sprintf(
					"HTTP/1.1 %d %s\r\nConnection: Close\r\n\r\n%s",
					StatusBadRequest,
					StatusText(StatusBadRequest),
					"")
			default:
				return fmt.Sprintf(
					"HTTP/1.1 %d %s\r\nContent-Type: %s\r\nConnection: Close\r\n\r\n%s",
					StatusCreated,
					StatusText(StatusCreated),
					"application/json;charset=utf-8",
					bout.String())
			}
		}
	case "/todo/1":
		rbody := r.Body
		r.Body = nil
		var bout bytes.Buffer
		if rbody != nil {
			_, err := io.Copy(&bout, rbody)
			if err != nil {
				log.Fatalf("Request: copying body: %v", err)
			}
			rbody.Close()
		}
		switch r.Method {
		case "PUT":
			switch r.Header.Get("Content-Type") {
			case "text/xml":
				return fmt.Sprintf(
					"HTTP/1.1 %d %s\r\nConnection: Close\r\n\r\n%s",
					StatusUnsupportedMediaType,
					StatusText(StatusUnsupportedMediaType),
					"")
			case "application/x-www-form-urlencoded":
				return fmt.Sprintf(
					"HTTP/1.1 %d %s\r\nConnection: Close\r\n\r\n%s",
					StatusBadRequest,
					StatusText(StatusBadRequest),
					"")
			default:
				return fmt.Sprintf(
					"HTTP/1.1 %d %s\r\nContent-Type: %s\r\nConnection: Close\r\n\r\n%s",
					StatusOK,
					StatusText(StatusOK),
					"application/json;charset=utf-8",
					bout.String())
			}
		case "DELETE":
			return fmt.Sprintf(
				"HTTP/1.1 %d %s\r\nConnection: Close\r\n\r\n%s",
				StatusNoContent,
				StatusText(StatusNoContent),
				"")
		}
	case "/todo/1000":
		switch r.Method {
		case "DELETE":
			return fmt.Sprintf(
				"HTTP/1.1 %d %s\r\nConnection: Close\r\n\r\n%s",
				StatusNoContent,
				StatusText(StatusNoContent),
				"")
		}
	default:
		return fmt.Sprintf(
			"HTTP/1.1 %d %s\r\nConnection: Close\r\n\r\n%s",
			StatusNotFound,
			StatusText(StatusNotFound),
			"")
	}
	return ""
}
