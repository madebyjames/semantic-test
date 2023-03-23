package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handleRoot)

	err := http.ListenAndServe(":3000", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("Web server closed")
	}
	if err != nil {
		fmt.Printf("Error starting web server: %s\n", err)
		os.Exit(1)
	}
}

func handleRoot(w http.ResponseWriter, _ *http.Request) {
	_, err := io.WriteString(w, "Hello world!")
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
}
