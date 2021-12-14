package main

import (
	"testing"
)

func TestParseRequestLine(t *testing.T) {
	testRequestLine := "GET /foo HTTP/1.1"
	actualMethod, actualRequestURI, actualProto, actualOk := parseRequestLine(testRequestLine)
	expectedMethod, expectedRequestURI, expectedProto, expectedOk := "GET", "/foo", "HTTP/1.1", true
	if actualMethod != expectedMethod || actualRequestURI != expectedRequestURI ||
		actualProto != expectedProto || actualOk != expectedOk {
		t.Fatalf("Unexpected result:\nGot:\t\t%s\nExpected:\t%s\n", actualMethod, expectedMethod)
		ClearTestBuffer()
	}
	ClearTestBuffer()
	return
}
