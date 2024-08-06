package main

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incr'ing thrice makes value 3", func(t *testing.T) {
		ctr := NewCounter()
		ctr.Incr()
		ctr.Incr()
		ctr.Incr()
		want := 3
		assertCounter(t, ctr, want)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		ctr := NewCounter()
		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func() {
				ctr.Incr()
				wg.Done()
			}()
		}
		wg.Wait()
		assertCounter(t, ctr, wantedCount)
	})
}

func assertCounter(t testing.TB, got *Counter, want int) {
	t.Helper()
	if got := got.Value(); got != want {
		t.Errorf("got: %d, want: %d", got, want)
	}
}
