package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my website! Version 0")
	})
	log.Fatal(http.ListenAndServe(":8500", nil))
}
