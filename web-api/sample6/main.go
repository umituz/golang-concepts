package main

import (
	"bytes"
	"html/template"
	"net/http"
)

type Page struct {
	Title           string
	Author          string
	Header          string
	PageDescription string
	Content         string
	URI             string
}

func handler(w http.ResponseWriter, r *http.Request) {
	// string combining
	var builder bytes.Buffer
	builder.WriteString("Hello Golang!\n")

	uri := "go.dev"

	page := Page{
		Title:           "Golang Is Great ",
		Author:          "Ümit Kenan UZ",
		Header:          "You cannot be great like Golang xD",
		PageDescription: "But you can try xD",
		Content:         builder.String(),
		URI:             "http://" + uri}
	t, _ := template.ParseFiles("page.html")

	t.Execute(w, page)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9000", nil)
}
