package main

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(urlA, urlB string) string {
	durationA := measureResponseTime(urlA)
	durationB := measureResponseTime(urlB)
	if durationA > durationB {
		return urlB
	}
	return urlA
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}

func main() {
	fmt.Println("vim-go")

}
