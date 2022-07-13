package main

import (
	"fmt"
	"net/http"
)

type Person struct {
	FirstName string
	LastName  string
}

func (person Person) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	person.FirstName = "Ümit Kenan"
	person.LastName = "UZ"

	fmt.Fprintf(w, "<p>"+person.FirstName+" "+person.LastName+" "+"</p>")

}

func main() {
	var person Person

	err := http.ListenAndServe("localhost:8000", person)

	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
