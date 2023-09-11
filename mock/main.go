package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("pong"))
	})

	http.ListenAndServe(":8080", nil)
}
