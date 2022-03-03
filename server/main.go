package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Persona struct {
	Id       int
	Nombre   string
	Apellido string
}

var Personas []Persona

func getPersonas(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Personas)
}

func main() {
	handleRequest()
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home page HIt")
}

func returnSinglePersona(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	// Loop over all of our Articles
	// if the article.Id equals the key we pass in
	// return the article encoded as JSON
	for _, persona := range Personas {
		id := string(persona.Id)
		if id == key {
			json.NewEncoder(w).Encode(persona)
		}
	}
}

func createNewPersona(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var persona Persona
	json.Unmarshal(reqBody, &persona)
	// update our global personas array to include
	// our new persona
	Personas = append(Personas, persona)

	json.NewEncoder(w).Encode(persona)
}

func handleRequest() {
	Personas = []Persona{
		Persona{Id: 1, Nombre: "pepe", Apellido: "garcia"},
	}
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/personas", getPersonas)
	// NOTE: Ordering is important here! This has to be defined before
	// the other `/persona` endpoint.
	myRouter.HandleFunc("/postPersona", createNewPersona).Methods("POST")
	myRouter.HandleFunc("/persona/{id}", returnSinglePersona)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}
