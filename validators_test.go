package main

import (
	"testing"
)

func TestIsValidMethod(t *testing.T) {
	method := "GET"
	if actual, expected := isValidMethod(method), true; actual != expected  {
		t.Fatalf("Unexpected output:\nGot:\t%t\nExpected:\t%t\n", actual, expected)
	}
	return
}

func TestIsNotValidMethod(t *testing.T) {
	method := " "
	if actual, expected := isValidMethod(method), false; actual != expected  {
		t.Fatalf("Unexpected output:\nGot:\t%t\nExpected:\t%t\n", actual, expected)
	}
	return
}
