package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"testing"
)

// arg1 means argument 1 and arg2 means argument 2, and the expected stands for the 'result we expect'
type addTest struct {
	arg1, arg2, expected int
}

var addTests = []addTest{
	addTest{2, 3, 5},
	addTest{4, 8, 12},
	addTest{6, 9, 15},
	addTest{3, 10, 13},
}

func TestAdd(t *testing.T) {

	for _, test := range addTests {
		if output := Add(test.arg1, test.arg2); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(4, 6)
	}
}


func ExampleAdd() {
	fmt.Println(Add(4, 6))
	// Output: 10
}

func TestConn(t *testing.T) {
	message := "Hi there!\n"

	go func() {
		conn, err := net.Dial("tcp", "localhost:8080")
		if err != nil {
			t.Fatal(err)
		}
		defer conn.Close()

		if _, err := fmt.Fprintf(conn, message); err != nil {
			t.Fatal(err)
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
