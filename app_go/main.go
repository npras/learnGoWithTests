package main

import (
	"log"
	"net/http"
	"os"
)

const dbFileName = "game.db.json"

func main() {
	dataFile, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)
	store, err := NewFileSystemStore(dataFile)
	if err != nil {
		log.Fatalf("Problem creating filesystem PlayStore: %v", err)
	}
	srv := NewPlayerServer(store)

	if err := http.ListenAndServe(":5000", srv); err != nil {
		log.Fatalf("couldn't listen on port 5000, %v", err)
	}
}
