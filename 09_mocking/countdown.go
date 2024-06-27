package main


import (
  "io"
  "fmt"
  "os"
  "time"
)


const (
  finalWord = "Go!"
  countdownStart = 3
)


type Sleeper interface {
  Sleep()
}


type DefaultSleeper struct {}


func (s *DefaultSleeper) Sleep() {
  time.Sleep(1 * time.Second)
}


func Countdown(out io.Writer, sleeper Sleeper) {
  for i := countdownStart; i > 0; i-- {
    fmt.Fprintln(out, i)
    sleeper.Sleep()
  }
  fmt.Fprint(out, finalWord)
}


func main() {
  sleeper := &DefaultSleeper{}
  Countdown(os.Stdout, sleeper)
}
