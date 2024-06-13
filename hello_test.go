package main

import "testing"

func TestHello(t *testing.T) {
  got := Hello("Prs")
  want := "Hello Prs!"

  if got != want {
    t.Errorf("got: %q, want: %q", got, want)
  }
}
