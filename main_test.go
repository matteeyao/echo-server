package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"testing"
)

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
