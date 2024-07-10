package main

import (
	"fmt"
	"net/http"
	"time"
)

/*
func Racer(urlA, urlB string) (string, error) {
	durationA := measureResponseTime(urlA)
	durationB := measureResponseTime(urlB)
	if durationA > durationB {
		return urlB, nil
	}
	return urlA, nil
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
*/

var timeOutLimit = 10 * time.Second

func Racer(urlA, urlB string) (winner string, error error) {
	return ConfigurableRacer(urlA, urlB, timeOutLimit)
}

func ConfigurableRacer(urlA, urlB string, timeout time.Duration) (winner string, error error) {
	select {
	case <-ping(urlA):
		return urlA, nil
	case <-ping(urlB):
		return urlB, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", urlA, urlB)
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

func main() {
	fmt.Println("vim-go")

}
