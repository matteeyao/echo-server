package main

import (
	"bufio"
	"bytes"
	"log"
	"net"
	"sync"
)

type server struct {
	Addr 	string
	Handler handler
}

func listenAndServe(addr string, handler handler) {
	server := &server{Addr: addr, Handler: handler}
	server.listenAndServe()
}

func (srv *server) listenAndServe() {
	ln, err := net.Listen("tcp", srv.Addr)
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
		go srv.handleConn(conn)
	}
}

func (srv *server) handleConn(c net.Conn) {
	wg := &sync.WaitGroup{}
	request := make([]byte, 1028)
	readLength, err := c.Read(request)
	if err != nil {
		return
	}
	wg.Add(1)
	go srv.ReadRequest(c, request[:readLength], wg)
	wg.Wait()
	c.Close()
}

// ReadRequest reads and parses an incoming request from b.
func (srv *server) ReadRequest(c net.Conn, request []byte, wg *sync.WaitGroup) {
	req, err := readRequest(bufio.NewReader(bytes.NewReader(request)))
	if err != nil {
		return
	}
	res := new(Response)
	s := srv.Handler.ServeHTTP(res, req)
	c.Write([]byte(s))
	wg.Done()
}
