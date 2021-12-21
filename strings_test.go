package main

import "testing"

func TestCut(t *testing.T) {
	requestLine := "GET /todos HTTP/1.1"
	expectedMethod := "GET"
	expectedRemaining := "/todos HTTP/1.1"
	if actualMethod, actualRemaining, _ := Cut(requestLine, " ");
		!(actualMethod == expectedMethod && actualRemaining == expectedRemaining) {
			t.Fatalf("Unexpected result:\nGot:\t%s, %s\nExpected:\t%s, %s\n",
				actualMethod, actualRemaining, expectedMethod, expectedRemaining)
		}
	return
}
