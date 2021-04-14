package main

import (
	"log"
	"net/http"

	poker "github.com/mrityunjaygr8/tdd"
)

const dbFile = "game.db.json"

func main() {
	store, closeFunc, err := poker.FileSystemPlayerStoreFromFile(dbFile)
	if err != nil {
		log.Panic(err)
	}

	defer closeFunc()
	server := poker.NewPlayerServer(store)
	log.Fatal(http.ListenAndServe(":5000", server))
}
