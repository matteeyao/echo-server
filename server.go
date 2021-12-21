package main

import (
	"bufio"
	"bytes"
	"log"
	"net"
	"sync"
)

// A Server defines parameters for running an HTTP server.
// The zero value for Server is a valid configuration.
type Server struct {
	// Addr optionally specifies the TCP address for the server to listen on,
	// in the form "host:port".
	Addr string
	Handler Handler
}

// ListenAndServe listens on the TCP network address addr and then calls
// Serve with handler to handle requests on incoming connections.
// Accepted connections are configured to enable TCP keep-alives.
func ListenAndServe(addr string, handler Handler) {
	server := &Server{Addr: addr, Handler: handler}
	server.ListenAndServe()
}

func (srv *Server) ListenAndServe() {
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

func (srv *Server) handleConn(c net.Conn) {
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

func (srv *Server) ReadRequest(c net.Conn, request []byte, wg *sync.WaitGroup) {
	req, err := readRequest(bufio.NewReader(bytes.NewReader(request)))
	if err != nil {
		return
	}
	res := new(Response)
	s := srv.Handler.ServeHTTP(res, req)
	c.Write([]byte(s))
	wg.Done()
}
