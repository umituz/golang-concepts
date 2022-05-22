package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World With GoLang :)")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
