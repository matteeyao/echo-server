package main

func readTransfer(w *Response, r Request) (err error) {
	w.Header = r.Header
	w.Proto = r.Proto
	w.Close = true
	w.Header.Del("Connection")
	w.Header.Add("Connection", "Close")
	w.Body = r.Body
	return nil
}
