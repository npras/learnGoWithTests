package main


import (
  "testing"
  "bytes"
)


type SpySleeper struct {
  Calls int
}


func (s *SpySleeper) Sleep() {
  s.Calls++
}


func TestCountdown(t *testing.T) {
  buffer := &bytes.Buffer{}
  spySleeper := &SpySleeper{}
  Countdown(buffer, spySleeper)

  got := buffer.String()
  want := `3
2
1
Go!`

  if got != want {
    t.Errorf("got: %q, want: %q", got, want)
  }
  if spySleeper.Calls != 3 {
    t.Errorf("spySleeper calls: got: %d, want: 3", spySleeper.Calls)
  }
}
