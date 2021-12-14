package main

import (
	"fmt"
)

type MyMux struct {
}

// Error replies to the request with the specified error message and HTTP code.
// It does not otherwise end the request; the caller should ensure no further
// writes are done to w.
// The error message should be plain text.
func Error(w ResponseWriter, error string, code int) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	fmt.Fprintln(w, error)
}

// NotFound replies to the request with an HTTP 404 not found error.
func NotFound(w ResponseWriter, r *Request) { Error(w, "404 page not found", StatusNotFound) }

func (p *MyMux) ServeHTTP(w ResponseWriter, r *Request) {
	switch r.URL.Path {
	case "/head_request":
		w.WriteHeader(StatusMethodNotAllowed) // 405
	//case "/price":
	//	item := req.URL.Query().Get("item")
	//	price, ok := db[item]
	//	if !ok {
	//		w.WriteHeader(StatusNotFound) // 404
	//		fmt.Fprintf(w, "no such item: %q\n", item)
	//		return
	//	}
	//	fmt.Fprintf(w, "%s\n", price)
	//default:
	//	w.WriteHeader(http.StatusNotFound) // 404
	//	fmt.Fprintf(w, "no such page: %s\n", req.URL)
	}
}
