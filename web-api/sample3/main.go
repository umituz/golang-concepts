package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type API struct {
	Message string "json:message"
}

type User struct {
	ID        int    "json:id"
	FirstName string "json:firstname"
	LastName  string "json:lastname"
	Age       int    "json:age"
}

func main() {
	apiRoot := "/api"

	// /api
	http.HandleFunc(apiRoot, func(w http.ResponseWriter, r *http.Request) {
		message := API{"API Home"}
		output, err := json.Marshal(message)
		checkError(err)
		// w.Header().Set("Content-Type", "application/json")   // Not required
		// w.Write([]byte(output))              // Not required
		fmt.Fprintf(w, string(output))
	})

	// /api/users
	http.HandleFunc(apiRoot+"/users", func(w http.ResponseWriter, r *http.Request) {
		users := []User{
			User{ID: 1, FirstName: "Ümit Kenan", LastName: "UZ", Age: 26},
		}
		output, err := json.Marshal(users)
		checkError(err)
		// w.Header().Set("Content-Type", "application/json")   // Not required
		// w.Write([]byte(output))              // Not required
		fmt.Fprintf(w, string(output))
	})

	// /api/me
	http.HandleFunc(apiRoot+"/me", func(w http.ResponseWriter, r *http.Request) {
		user := User{3, "Ümit Kenan", "UZ", 26}
		output, err := json.Marshal(user)
		checkError(err)
		// w.Header().Set("Content-Type", "application/json")   // Not required
		// w.Write([]byte(output))              // Not required
		fmt.Fprintf(w, string(output))
	})

	// Serve on port 8080
	http.ListenAndServe(":8080", nil)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal Error : ", err.Error())
		os.Exit(1)
	}
}
