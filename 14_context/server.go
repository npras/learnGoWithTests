package main

import (
	"context"
	"fmt"
	"net/http"
)

func Server(store Store) http.HandlerFunc {
	return func(respWriter http.ResponseWriter, req *http.Request) {
		data, err := store.Fetch(req.Context())

		if err != nil {
			return // log error however
		}
		fmt.Fprint(respWriter, data)
	}
}

type Store interface {
	Fetch(ctx context.Context) (string, error)
}
