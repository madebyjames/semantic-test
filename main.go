package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleRoot)
}

func handleRoot(w http.ResponseWriter, _ *http.Request) {
	_, err := io.WriteString(w, "Hello world")
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
}
