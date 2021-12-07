package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"sync"
)

// HandleFunc registers the handler function for the given pattern
// in the DefaultServeMux.
// The documentation for ServeMux explains how patterns are matched.
func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
	DefaultServeMux.HandleFunc(pattern, handler)
}

// READING IN THE BODY, AND NOT JUST THE REQUEST HEADER, IF YOU CAN CONSUME THE REQUEST HEADER AND REQUEST BODY

func echo(c net.Conn, uri string, wg *sync.WaitGroup) {
	//u := strings.Fields(uri)[1]
	//body := strings.Split(u, "=")[1]
	//content := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/html\r\nConnection: Close\r\n\r\n%s\r\n", body)
	content := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/html\r\nConnection: Close\r\n\r\n%+v\r\n", uri)
	c.Write([]byte(content))
	wg.Done()
}

func handleConn(c net.Conn) {
	wg := &sync.WaitGroup{}

	// Returns request header, need to parse request body TODO: Use Postman checkbox for body
	input := bufio.NewScanner(c)
	input.Scan()
	wg.Add(1)
	go echo(c, input.Text(), wg)

	wg.Wait()
	// NOTE: ignoring potential errors from input.Err()
	c.Close()
}

func main() {
	ln, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
