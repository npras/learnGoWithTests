package main

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
	"time"
)

const (
	sleepStr = "zzzzz\n"
)

type SpecialSpy struct {
	Outputs []string
}

func (ss *SpecialSpy) Sleep() {
	ss.Outputs = append(ss.Outputs, sleepStr)
}

func (ss *SpecialSpy) Write(written []byte) (n int, err error) {
	ss.Outputs = append(ss.Outputs, string(written[:]))
	return
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second
	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()
	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}

func TestCountdown(t *testing.T) {
	t.Run("prints 3, 2, 1, Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := &SpecialSpy{}
		Countdown(buffer, spySleeper)
		got := buffer.String()
		want := `3
2
1
Go!`
		if got != want {
			t.Errorf("got: %q, want: %q", got, want)
		}
	})

	t.Run("prints 3, 2, 1, Go!", func(t *testing.T) {
		spy := SpecialSpy{}
		spySleeper := &spy
		spyWriter := &spy
		Countdown(spyWriter, spySleeper)
		got := spy.Outputs
		fmt.Println(got)
		want := []string{"3\n", sleepStr, "2\n", sleepStr, "1\n", sleepStr, "Go!"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %q, want: %q", got, want)
		}
	})
}
