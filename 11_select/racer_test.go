package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func makeDelayedServer(delay time.Duration) *httptest.Server {
	handlerFn := func(res http.ResponseWriter, req *http.Request) {
		time.Sleep(delay)
		res.WriteHeader(http.StatusOK)
	}
	return httptest.NewServer(http.HandlerFunc(handlerFn))
}

func TestRacer(t *testing.T) {
	t.Run("compares speed of 2 urls, returns fastest", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)
		defer slowServer.Close()
		defer fastServer.Close()

		slowUrl := slowServer.URL
		fastUrl := fastServer.URL
		fmt.Println(slowUrl)
		fmt.Println(fastUrl)

		want := fastUrl
		got, _ := Racer(slowUrl, fastUrl)

		if got != want {
			t.Errorf("got: %q, want: %q", got, want)
		}
	})

	t.Run("returns error if no respond within 10s", func(t *testing.T) {
		serverA := makeDelayedServer(1 * time.Second)
		defer serverA.Close()

		_, err := ConfigurableRacer(serverA.URL, serverA.URL, 0*time.Second)

		if err == nil {
			t.Errorf("Expected an error, but didn't get one")
		}
	})
}
