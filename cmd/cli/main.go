package main

import (
	"fmt"
	"log"
	"os"

	poker "github.com/mrityunjaygr8/tdd"
)

const dbFile = "game.db.json"

func main() {
	fmt.Println("Let's play poker")
	fmt.Println("Type (Name) wins to record a win")

	store, closeFunc, err := poker.FileSystemPlayerStoreFromFile(dbFile)
	if err != nil {
		log.Panic(err)
	}

	defer closeFunc()

	game := poker.NewCLI(store, os.Stdin)
	game.PlayPoker()
}
