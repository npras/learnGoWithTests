package main

import (
	"log"
	"net/http"
)

func main() {
	srv := NewPlayerServer(NewInMemoryPlayerStore())
	log.Fatal(http.ListenAndServe(":5000", srv))
}
