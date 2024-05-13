package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	mainpage := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.Execute(w, nil)
	}

	router.HandleFunc("/", mainpage)
	router.HandleFunc("/characters", GetAllCharacters)                 // get all the characters
	router.HandleFunc("/character/{name}", GetCharacter)               // get a specifc character
	router.HandleFunc("/characters/add", AddCharacter).Methods("POST") // Add a new character

	fmt.Println("Server running on port 10000")
	log.Fatal(http.ListenAndServe(":10000", router))
}
