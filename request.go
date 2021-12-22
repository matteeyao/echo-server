package main

import (
	"bufio"
	"io"
	"net/textproto"
	"net/url"
	"sync"
)

type Request struct {
	Method 		string
	URL 		*url.URL
	Proto		string
	Header 		Header
	Body 		string
	RequestURI 	string
}

var textprotoReaderPool sync.Pool

func newTextprotoReader(br *bufio.Reader) *textproto.Reader {
	if v := textprotoReaderPool.Get(); v != nil {
		tr := v.(*textproto.Reader)
		tr.R = br
		return tr
	}
	return textproto.NewReader(br)
}

func putTextprotoReader(r *textproto.Reader) {
	r.R = nil
	textprotoReaderPool.Put(r)
}

// parseRequestLine parses "GET /foo HTTP/1.1" into its three parts.
func parseRequestLine(line string) (method, requestURI, proto string, ok bool) {
	method, rest, ok1 := Cut(line, " ")
	requestURI, proto, ok2 := Cut(rest, " ")
	if !ok1 || !ok2 {
		return "", "", "", false
	}
	return method, requestURI, proto, true
}

func readRequest(b *bufio.Reader) (req *Request, err error) {
	tp := newTextprotoReader(b)
	req = new(Request)

	// First line: GET /index.html HTTP/1.0
	var s string
	if s, err = tp.ReadLine(); err != nil {
		return nil, err
	}
	defer func() {
		putTextprotoReader(tp)
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	var ok bool
	req.Method, req.RequestURI, req.Proto, ok = parseRequestLine(s)
	if !ok {
		return nil, badStringError("malformed HTTP request", s)
	}
	if !isValidMethod(req.Method) {
		return nil, badStringError("invalid method", req.Method)
	}

	if req.URL, err = url.ParseRequestURI(req.RequestURI); err != nil {
		return nil, err
	}

	// Subsequent lines: Key: value.
	MIMEHeader, err := tp.ReadMIMEHeader()
	if err != nil {
		return nil, err
	}
	req.Header = Header(MIMEHeader)

	// Body
	for {
		str, err := tp.ReadLine()
		if len(str) == 0 {
			return req, nil
		}
		if err != nil {
			return req, nil
		}
		req.Body += str + "\n"
	}
}
