package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

var characters = []RaiderioCharacter{}

func getRaiderioCharacter(name string, realm string) RaiderioCharacter {
	resp, err := http.Get(fmt.Sprintf("https://raider.io/api/v1/characters/profile?region=eu&realm=%s&name=%s&fields=mythic_plus_recent_runs", realm, name))
	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var character RaiderioCharacter
	charByte := []byte(body)

	if err := json.Unmarshal(charByte, &character); err != nil {
		panic(err)
	}

	return character
}

func GetAllCharacters(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint GetAllCharacters")
	tmpl := template.Must(template.ParseFiles("charlist.html"))
	tmpl.Execute(w, characters)

}

func GetCharacter(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get a specific character endpoint")
	vars := mux.Vars(r)

	name := vars["name"]

	for _, character := range characters {
		if character.Name == name {
			json.NewEncoder(w).Encode(character)
		}
	}

}

func AddCharacter(w http.ResponseWriter, r *http.Request) {
	fmt.Println("add a new character endpoint")

	name := r.PostFormValue("name")

	realm := r.PostFormValue("realm")

	character := getRaiderioCharacter(name, realm)
	characters = append(characters, character)
	tmpl := template.Must(template.ParseFiles("charlist.html"))
	tmpl.Execute(w, characters)

}
