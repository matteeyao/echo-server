package main

// A Handler responds to an HTTP request.
//
// ServeHTTP should write reply headers and data to the ResponseWriter
// and then return. Returning signals that the request is finished; it
// is not valid to use the ResponseWriter or read from the
// Request.Body after or concurrently with the completion of the
// ServeHTTP call.
//
// Depending on the HTTP client software, HTTP protocol version, and
// any intermediaries between the client and the Go server, it may not
// be possible to read from the Request.Body after writing to the
// ResponseWriter. Cautious handlers should read the Request.Body
// first, and then reply.
//
// Except for reading the body, handlers should not modify the
// provided Request.
//
// If ServeHTTP panics, the server (the caller of ServeHTTP) assumes
// that the effect of the panic was isolated to the active request.
// It recovers the panic, logs a stack trace to the server error log,
// and either closes the network connection or sends an HTTP/2
// RST_STREAM, depending on the HTTP protocol. To abort a handler so
// the client sees an interrupted response but the server doesn't log
// an error, panic with the value ErrAbortHandler.
type Handler interface {
	ServeHTTP(*Response, *Request) string
}

func ApplyStatusToResponse(w *Response, statusCode int) {
	w.StatusCode = statusCode
	w.Status = StatusText(w.StatusCode)
}

func HeadRequest (w *Response, r *Request) {
	switch r.Method {
	case "GET":
		ApplyStatusToResponse(w, StatusMethodNotAllowed)
		w.Header.Add("Allow", "HEAD")
		w.Header.Add("Allow", "OPTIONS")
	default:
		ApplyStatusToResponse(w, StatusOK)
		return
	}
}

func NotFound (w *Response, r *Request) {
	ApplyStatusToResponse(w, StatusNotFound)
}

func SimpleGet (w *Response, r *Request) {
	ApplyStatusToResponse(w, StatusOK)
}

func SimpleGetWithBody (w *Response, r *Request) {
	ApplyStatusToResponse(w, StatusOK)
	w.Header.Add("Content-Type", "text/html")
	expectedBody := "Hello world"
	w.Body = expectedBody
}

func SimpleHead (w *Response, r *Request) {
	ApplyStatusToResponse(w, StatusOK)
}

func MethodOptions (w *Response, r *Request) {
	ApplyStatusToResponse(w, StatusOK)
	w.Header.Add("Allow", "GET")
	w.Header.Add("Allow", "HEAD")
	w.Header.Add("Allow", "OPTIONS")
}

func MethodOptions2 (w *Response, r *Request) {
	ApplyStatusToResponse(w, StatusOK)
	w.Header.Add("Allow", "GET")
	w.Header.Add("Allow", "HEAD")
	w.Header.Add("Allow", "OPTIONS")
	w.Header.Add("Allow", "PUT")
	w.Header.Add("Allow", "POST")
}

func Redirect (w *Response, r *Request) {
	ApplyStatusToResponse(w, StatusMovedPermanently)
	w.Header.Add("Location", "http://127.0.0.1:5000/simple_get")
}

func EchoBody (w *Response, r *Request) {
	switch r.Method {
	case "POST":
		ApplyStatusToResponse(w, StatusOK)
	default:
		ApplyStatusToResponse(w, StatusMethodNotAllowed)
	}
}

