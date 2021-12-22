package main

type Response struct {
	Status     	string
	StatusCode 	int
	Proto      	string
	Header 		Header
	Body 		string
	Close 		bool
}
