package main

import (
	"fmt"
	"net/http"
	"strings"
)

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	out := ""
	if player == "Floyd" {
		out = "10"
	}
	if player == "Pepper" {
		out = "20"
	}
	fmt.Fprint(w, out)
}
