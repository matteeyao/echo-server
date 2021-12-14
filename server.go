package main

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"net"
	"strings"
	"sync"
)

// A Server defines parameters for running an HTTP server.
// The zero value for Server is a valid configuration.
type Server struct {
	// Addr optionally specifies the TCP address for the server to listen on,
	// in the form "host:port". If empty, ":http" (port 80) is used.
	// The service names are defined in RFC 6335 and assigned by IANA.
	// See net.Dial for details of the address format.
	Addr string

	Handler Handler // handler to invoke, http.DefaultServeMux if nil
}

// ListenAndServe listens on the TCP network address addr and then calls
// Serve with handler to handle requests on incoming connections.
// Accepted connections are configured to enable TCP keep-alives.
//
// The handler is typically nil, in which case the DefaultServeMux is used.
//
// ListenAndServe always returns a non-nil error.
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
	//respHeaders := "HTTP/1.1 200 OK\r\nContent-Type: text/html\r\nConnection: Close\r\n\r\n"
	//c.Write([]byte(respHeaders))
	wg := &sync.WaitGroup{}
	request := make([]byte, 1028)
	readLength, err := c.Read(request)
	if err != nil {
		return
	}
	wg.Add(1)
	go srv.ProcessRequest(c, string(request[:readLength]), wg)
	wg.Wait()
	c.Close()
}

func (srv *Server) ProcessRequest(c net.Conn, request string, wg *sync.WaitGroup) {
	req, err := ReadRequest(bufio.NewReader(strings.NewReader(request)))
	if err != nil {
		return
	}
	//srv.Handler.ServeHTTP(, request)


	rbody := req.Body
	req.Body = nil
	var bout bytes.Buffer
	if rbody != nil {
		_, err := io.Copy(&bout, rbody)
		if err != nil {
			log.Fatalf("Request: copying body: %v", err)
		}
		rbody.Close()
	}

	c.Write(bout.Bytes())
	wg.Done()
}

// maxPostHandlerReadBytes is the max number of Request.Body bytes not
// consumed by a handler that the server will read from the client
// in order to keep a connection alive. If there are more bytes than
// this then the server to be paranoid instead sends a "Connection:
// close" response.
//
// This number is approximately what a typical machine's TCP buffer
// size is anyway.  (if we have the bytes on the machine, we might as
// well read them)
const maxPostHandlerReadBytes = 256 << 10
