// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 224.

// Reverb2 is a TCP server that simulates an echo.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"sync"
	"time"
)

type options struct {
	port string
}

var opt options

func init() {
	flag.StringVar(&opt.port, "p", os.Getenv("PORT"), "The default port to listen on")
	flag.Parse()

	if opt.port == "" {
		opt.port = "8000"
	}
}

func echo(c net.Conn, shout string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Fprintln(c, "\t", shout)
}

func scan(r io.Reader, lines chan<- string) {
	s := bufio.NewScanner(r)
	for s.Scan() {
		lines <- s.Text()
	}
	// scan will most likely try to read from the connection after it's closed
	// by handleConn. I don't know how to avoid this. Go seems to shun async io
	// in favour of goroutines, so it probably isn't worth avoiding.
	if s.Err() != nil {
		log.Print("scan: ", s.Err())
	}
}

func handleConn(c net.Conn) {
	wg := &sync.WaitGroup{}
	defer func() {
		wg.Wait()
		c.Close()
	}()
	lines := make(chan string)
	go scan(c, lines)
	timeout := 2 * time.Second
	timer := time.NewTimer(2 * time.Second)
	for {
		select {
		case line := <-lines:
			timer.Reset(timeout)
			wg.Add(1)
			go echo(c, line, wg)
		case <-timer.C:
			return
		}
	}
}

func main() {
	ln, err := net.Listen("tcp", ":" + opt.port)
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

// -----------------

// Add is our function that sums two integers
func Add(x, y int) (res int) {
	return x + y
}

// Subtract subtracts two integers
func Subtract(x, y int) (res int) {
	return x - y
}
