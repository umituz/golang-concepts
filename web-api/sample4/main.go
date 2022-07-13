package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type API struct {
	Message string "json:message"
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	id := urlParams["id"]
	messageConcat := "User ID: " + id

	message := API{messageConcat}
	output, err := json.Marshal(message)

	if err != nil {
		fmt.Printf("Errrrroooooorrrrrr %s", err)
	}
	fmt.Fprintf(w, string(output))
}

func main() {
	gorillaRoute := mux.NewRouter()
	gorillaRoute.HandleFunc("/api/users/{id:[0-9]+}", UserHandler)

	http.Handle("/", gorillaRoute)

	http.ListenAndServe(":8080", nil)
}
