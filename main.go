package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"sync"
)

func echo(c net.Conn, body string, wg *sync.WaitGroup) {
	fmt.Fprintln(c, "\t", body)
	wg.Done()
}

func handleConn(c net.Conn) {
	wg := &sync.WaitGroup{}
	input := bufio.NewScanner(c)
	for input.Scan() {
		wg.Add(1)
		go echo(c, input.Text(), wg)
	}
	wg.Wait()
	// NOTE: ignoring potential errors from input.Err()
	c.Close()
}

func main() {
	ln, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
