package main

import "fmt"

func badStringError(what, val string) error { return fmt.Errorf("%s %q", what, val) }
