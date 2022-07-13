package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var homepage = "index.html"

func loadFile(fileName string) (string, error) {
	bytes, err := ioutil.ReadFile(fileName)

	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	var body, _ = loadFile(homepage)
	fmt.Fprintf(w, body)
	//http.ServeFile(w,r,body)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9000", nil)
}
