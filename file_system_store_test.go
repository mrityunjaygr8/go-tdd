package main

import (
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("league from a reader", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 17}]`)
		defer cleanDatabase()

		store := FileSystemPlayerStore{database}

		got := store.GetLeague()

		want := []Player{
			{"Cleo", 10},
			{"Chris", 17},
		}

		assertLeague(t, got, want)

		got = store.GetLeague()
		assertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 17}
		]`)
		defer cleanDatabase()

		store := FileSystemPlayerStore{database: database}

		got := store.GetPlayerScore("Chris")
		want := 17

		assertScoreEquals(t, got, want)
	})

	t.Run("it records the win for existing players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 17}
		]`)
		defer cleanDatabase()

		store := FileSystemPlayerStore{database: database}

		store.RecordWin("Cleo")

		got := store.GetPlayerScore("Cleo")
		want := 11

		assertScoreEquals(t, got, want)
	})

	t.Run("it records the win for new players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 17}
		]`)
		defer cleanDatabase()

		store := FileSystemPlayerStore{database: database}

		store.RecordWin("Pepper")

		got := store.GetPlayerScore("Pepper")
		want := 1

		assertScoreEquals(t, got, want)
	})
}

func assertScoreEquals(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func createTempFile(t testing.TB, initialData string) (io.ReadWriteSeeker, func()) {
	t.Helper()

	tmpFile, err := ioutil.TempFile("", "db")

	if err != nil {
		t.Fatalf("could not create temp file: %v", err)
	}

	tmpFile.Write([]byte(initialData))

	removeFile := func() {
		tmpFile.Close()
		os.Remove(tmpFile.Name())
	}

	return tmpFile, removeFile
}
