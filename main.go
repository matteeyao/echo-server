package main

import (
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
)

func parseRequestBody(request string) string {
	fields := strings.Fields(request)
	lastHeaderField := "Content-Length:"
	var requestHeadersEndIdx int
	for idx, field := range fields {
		if field == lastHeaderField {
			requestHeadersEndIdx = idx
		}
	}
	body := fields[requestHeadersEndIdx + 2:]
	return strings.Join(body, " ")
}

func echoRequestBody(c net.Conn, request string, wg *sync.WaitGroup) {
	requestBody := parseRequestBody(request)
	respContent := fmt.Sprintf("%s", requestBody)
	c.Write([]byte(respContent))
	wg.Done()
}

func handleConn(c net.Conn) {
	respHeaders := "HTTP/1.1 200 OK\r\nContent-Type: text/html\r\nConnection: Close\r\n\r\n"
	c.Write([]byte(respHeaders))
	wg := &sync.WaitGroup{}
	request := make([]byte, 1028)
	readLength, err := c.Read(request)
	if err != nil {
		return
	}
	wg.Add(1)
	go echoRequestBody(c, string(request[:readLength]), wg)
	wg.Wait()
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
			log.Print("Connection aborted:", err)
			continue
		}
		go handleConn(conn)
	}
}
