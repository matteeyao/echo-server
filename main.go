package main

func main() {
	mux := &MyMux{}
	ListenAndServe("localhost:5000", mux)
}
