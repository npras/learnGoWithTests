package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response string
	t        *testing.T
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)
	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				log.Println("spy store got cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

type SpyResponseWriter struct {
	written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}

func TestServer(t *testing.T) {
	t.Run("returns data from store", func(t *testing.T) {
		data := "hello, WORLD"
		store := &SpyStore{response: data, t: t}
		svr := Server(store)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		respRecorder := httptest.NewRecorder()

		svr.ServeHTTP(respRecorder, req)

		if got := respRecorder.Body.String(); got != data {
			t.Errorf(`got: %q, want: %q`, got, data)
		}
	})

	t.Run("tells store to cancel work if req is cancelled", func(t *testing.T) {
		data := "hello, WORLD"
		store := &SpyStore{response: data, t: t}
		svr := Server(store)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		cancellingCtx, cancelFn := context.WithCancel(req.Context())
		time.AfterFunc(5*time.Millisecond, cancelFn)
		req = req.WithContext(cancellingCtx)

		resp := &SpyResponseWriter{}

		svr.ServeHTTP(resp, req)

		if resp.written {
			t.Error("a response should NOT have been written")
		}
	})
}
