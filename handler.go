package main

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
	case "HEAD":
		ApplyStatusToResponse(w, StatusOK)
	case "OPTIONS":
		ApplyStatusToResponse(w, StatusOK)
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
	host := viperEnvVariable("REDIRECT_HOST")
	port := viperEnvVariable("server.port")
	endpoint := viperEnvVariable("REDIRECT_ENDPOINT")
	address := host + ":" + port + endpoint
	w.Header.Add("Location", address)
}

func EchoBody (w *Response, r *Request) {
	switch r.Method {
	case "POST":
		ApplyStatusToResponse(w, StatusOK)
	default:
		ApplyStatusToResponse(w, StatusMethodNotAllowed)
	}
}

