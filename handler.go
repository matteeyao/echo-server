package main

type handler interface {
	ServeHTTP(*Response, *Request) string
}

func applyStatusToResponse(w *Response, statusCode int) {
	w.StatusCode = statusCode
	w.Status = StatusText(w.StatusCode)
}

func headRequest(w *Response, r *Request) {
	switch r.Method {
	case "GET":
		applyStatusToResponse(w, StatusMethodNotAllowed)
		w.Header.Add("Allow", "HEAD")
		w.Header.Add("Allow", "OPTIONS")
	case "HEAD":
		applyStatusToResponse(w, StatusOK)
	case "OPTIONS":
		applyStatusToResponse(w, StatusOK)
	}
}

func notFound(w *Response, r *Request) {
	applyStatusToResponse(w, StatusNotFound)
}

func simpleGet(w *Response, r *Request) {
	applyStatusToResponse(w, StatusOK)
}

func simpleGetWithBody(w *Response, r *Request) {
	applyStatusToResponse(w, StatusOK)
	w.Header.Add("Content-Type", "text/html")
	expectedBody := "Hello world"
	w.Body = expectedBody
}

func simpleHead(w *Response, r *Request) {
	applyStatusToResponse(w, StatusOK)
}

func methodOptions(w *Response, r *Request) {
	applyStatusToResponse(w, StatusOK)
	w.Header.Add("Allow", "GET")
	w.Header.Add("Allow", "HEAD")
	w.Header.Add("Allow", "OPTIONS")
}

func methodOptions2(w *Response, r *Request) {
	applyStatusToResponse(w, StatusOK)
	w.Header.Add("Allow", "GET")
	w.Header.Add("Allow", "HEAD")
	w.Header.Add("Allow", "OPTIONS")
	w.Header.Add("Allow", "PUT")
	w.Header.Add("Allow", "POST")
}

func redirect(w *Response, r *Request) {
	applyStatusToResponse(w, StatusMovedPermanently)
	host := viperEnvVariable("REDIRECT_HOST")
	port := viperEnvVariable("server.port")
	endpoint := viperEnvVariable("REDIRECT_ENDPOINT")
	address := host + ":" + port + endpoint
	w.Header.Add("Location", address)
}

func echoBody(w *Response, r *Request) {
	switch r.Method {
	case "POST":
		applyStatusToResponse(w, StatusOK)
	default:
		applyStatusToResponse(w, StatusMethodNotAllowed)
	}
}

func textResponse(w *Response, r *Request) {
	applyStatusToResponse(w, StatusOK)
	w.Header.Add("Content-Type", "text/plain;charset=utf-8")
	expectedBody := "text response"
	w.Body = expectedBody
}

func htmlResponse(w *Response, r *Request) {
	applyStatusToResponse(w, StatusOK)
	w.Header.Add("Content-Type", "text/html;charset=utf-8")
	expectedBody := "<html><body><p>HTML Response</p></body></html>"
	w.Body = expectedBody
}

func jsonResponse(w *Response, r *Request) {
	applyStatusToResponse(w, StatusOK)
	w.Header.Add("Content-Type", "application/json;charset=utf-8")
	expectedBody := "{\"key1\":\"value1\",\"key2\":\"value2\"}"
	w.Body = expectedBody
}

func xmlResponse(w *Response, r *Request) {
	applyStatusToResponse(w, StatusOK)
	w.Header.Add("Content-Type", "application/xml;charset=utf-8")
	expectedBody := "<note><body>XML Response</body></note>"
	w.Body = expectedBody
}

func kittehResponse(w *Response, r *Request) {
	applyStatusToResponse(w, StatusOK)
	w.Header.Add("Content-Type", "image/jpeg")
	expectedBody := "image/jpeg"
	w.Body = expectedBody
}

func doggoResponse(w *Response, r *Request) {
	applyStatusToResponse(w, StatusOK)
	w.Header.Add("Content-Type", "image/png")
	expectedBody := "image/png"
	w.Body = expectedBody
}

func kissesResponse(w *Response, r *Request) {
	applyStatusToResponse(w, StatusOK)
	w.Header.Add("Content-Type", "image/gif")
	expectedBody := "image/gif"
	w.Body = expectedBody
}

func healthCheckResponse(w *Response, r *Request) {
	applyStatusToResponse(w, StatusOK)
	w.Header.Add("Content-Type", "text/html;charset=utf-8")
	expectedBody := "<html><body><<strong>Status:</strong> pass</body></html>"
	w.Body = expectedBody
}

func todoResponse(w *Response, r *Request) {
	switch r.Method {
	case "POST":
		switch r.Header.Get("Content-Type") {
		case "application/x-www-form-urlencoded":
			clearResponse(w)
			applyStatusToResponse(w, StatusBadRequest)
		case "text/xml; charset=utf-8":
			clearResponse(w)
			applyStatusToResponse(w, StatusUnsupportedMediaType)
		default:
			applyStatusToResponse(w, StatusCreated)
			w.Header.Add("Content-Type", "application/json;charset=utf-8")
		}
	}
}

func clearResponse(w *Response) {
	w.Header = *new(header)
	w.Body = ""
}

func handleTodoResponse(w *Response, r *Request) {
	switch r.Method {
	case "PUT":
		switch r.Header.Get("Content-Type") {
		case "text/xml":
			clearResponse(w)
			applyStatusToResponse(w, StatusUnsupportedMediaType)
		case "application/x-www-form-urlencoded":
			clearResponse(w)
			applyStatusToResponse(w, StatusBadRequest)
		default:
			applyStatusToResponse(w, StatusOK)
			w.Header.Add("Content-Type", "application/json;charset=utf-8")
		}
	case "DELETE":
		clearResponse(w)
		applyStatusToResponse(w, StatusNoContent)
	}
}
