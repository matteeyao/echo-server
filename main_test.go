package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"testing"
)

func TestConn(t *testing.T) {
	message := "Hi there!\n"

	go func() {
		conn, err := net.Dial("tcp", "localhost:8080")
		if err != nil {
			t.Error(err)
			return
		}
		defer conn.Close()

		if _, err := fmt.Fprintf(conn, message); err != nil {
			t.Error(err)
			return
		}
	}()

	l, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		t.Fatal(err)
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			return
		}
		defer conn.Close()

		buf, err := ioutil.ReadAll(conn)
		if err != nil {
			t.Fatal(err)
		}

		fmt.Println(string(buf[:]))
		if msg := string(buf[:]); msg != message {
			t.Fatalf("Unexpected message:\nGot:\t\t%s\nExpected:\t%s\n", msg, message)
		}
		return // Done
	}
}
