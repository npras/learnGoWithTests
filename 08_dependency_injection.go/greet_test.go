package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Joe")
	got := buffer.String()
	want := "Hello Joe"
	if got != want {
		t.Errorf("got: %q, want: %q", got, want)
	}
}
