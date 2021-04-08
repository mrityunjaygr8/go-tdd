package main

import (
	"fmt"
	"net/http"
)

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	out := "20"
	fmt.Fprintf(w, out)
}
