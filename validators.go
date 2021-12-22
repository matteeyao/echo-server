package main

import (
	"golang.org/x/net/http/httpguts"
	"strings"
)

func isValidMethod(method string) bool {
	return len(method) > 0 && strings.IndexFunc(method, isNotToken) == -1
}

func isNotToken(r rune) bool {
	return !httpguts.IsTokenRune(r)
}
