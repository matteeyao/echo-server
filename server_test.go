package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"sync"
	"testing"
	"time"
)

type MockedConn struct {}
var testBuffer = make([]byte, 1028)

// Read implements the Conn Read method.
func (c *MockedConn) Read(b []byte) (int, error) {
	testBuffer = append(testBuffer, b...)
	return 0, nil
}

// Write implements the Conn Write method.
func (c *MockedConn) Write(b []byte) (int, error) {
	return 0, nil
}

// Close closes the connection.
func (c *MockedConn) Close() error {
	return nil
}

type MockedAddr struct {}

func (c *MockedAddr) Network() string {
	return ""
}

func (c *MockedAddr) String() string {
	return ""
}

func (c *MockedConn) LocalAddr() net.Addr {
	return &MockedAddr{}
}


func (c *MockedConn) RemoteAddr() net.Addr {
	return &MockedAddr{}
}

// SetDeadline implements the Conn SetDeadline method.
func (c *MockedConn) SetDeadline(t time.Time) error {
	return nil
}

// SetReadDeadline implements the Conn SetReadDeadline method.
func (c *MockedConn) SetReadDeadline(t time.Time) error {
	return nil
}

// SetWriteDeadline implements the Conn SetWriteDeadline method.
func (c *MockedConn) SetWriteDeadline(t time.Time) error {
	return nil
}

func ClearTestBuffer() {
	testBuffer = make([]byte, 1028)
}

func TestEchoRequestBody(t *testing.T) {
	testRequest := "GET / HTTP/1.1 Content-Length: 4 test"
	testPhraseInBytes := []byte(testRequest)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	mux := &MyMux{}
	server := &Server{Addr: "localhost:5000", Handler: mux}
	server.ProcessRequest(&MockedConn{}, testRequest, wg)
	if actual := testBuffer; bytes.Compare(testPhraseInBytes, actual) == 0 {
		t.Fatalf("Unexpected message:\nGot:\t\t%s\nExpected:\t%s\n", actual, testPhraseInBytes)
		ClearTestBuffer()
	}
	ClearTestBuffer()
	return
}

func TestConn(t *testing.T) {
	message := "Hi there!\n"

	server, client := net.Pipe()
	go func() {
		client.Write([]byte(message))
		client.Close()
	}()
	buf, err := ioutil.ReadAll(server)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(buf[:]))
	if msg := string(buf[:]); msg != message {
		t.Fatalf("Unexpected message:\nGot:\t\t%s\nExpected:\t%s\n", msg, message)
	}
	return
}
