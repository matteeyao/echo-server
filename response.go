package main

// Response represents the response from an HTTP request.
type Response struct {
	Status     	string
	StatusCode 	int
	Proto      	string
	Header 		header
	Body 		string
	Close 		bool
}
