package main

import (
	"go/token"
	"reflect"
	"testing"
)

func diff(t *testing.T, prefix string, have, want interface{}) {
	t.Helper()
	hv := reflect.ValueOf(have).Elem()
	wv := reflect.ValueOf(want).Elem()
	if hv.Type() != wv.Type() {
		t.Errorf("%s: type mismatch %v want %v", prefix, hv.Type(), wv.Type())
	}
	for i := 0; i < hv.NumField(); i++ {
		name := hv.Type().Field(i).Name
		if !token.IsExported(name) {
			continue
		}
		hf := hv.Field(i).Interface()
		wf := wv.Field(i).Interface()
		if !reflect.DeepEqual(hf, wf) {
			t.Errorf("%s:\n\n%s (Actual) = %v\r\n%s (Expected) = %v", prefix, name, hf, name, wf)
		}
	}
}
