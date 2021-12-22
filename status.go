package main

const (
	StatusOK                   	= 200 // RFC 7231, 6.3.1
	StatusCreated             	= 201 // RFC 7231, 6.3.2
	StatusNoContent            	= 204 // RFC 7231, 6.3.5

	StatusMovedPermanently  	= 301 // RFC 7231, 6.4.2

	StatusNotFound              = 404 // RFC 7231, 6.5.4
	StatusMethodNotAllowed      = 405 // RFC 7231, 6.5.5
	StatusUnsupportedMediaType  = 415 // RFC 7231, 6.5.13
)

var statusText = map[int]string{
	StatusOK:                   "OK",
	StatusCreated:              "Created",
	StatusNoContent:            "No Content",

	StatusMovedPermanently:  	"Moved Permanently",

	StatusNotFound:             "Not Found",
	StatusMethodNotAllowed:     "Method Not Allowed",
	StatusUnsupportedMediaType: "Unsupported Media Type",
}

// StatusText returns a text for the HTTP status code. It returns the empty
// string if the code is unknown.
func StatusText(code int) string {
	return statusText[code]
}
