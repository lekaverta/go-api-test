package main

import (
	"encoding/json"
	"net/http"
	"io/ioutil"

	"github.com/gorilla/mux"
)

func main() {
	routes := mux.NewRouter().StrictSlash(true)

	routes.HandleFunc("/musics", getMusics).Methods("GET")
	routes.HandleFunc("/musics", createMusic).Methods("POST")

	http.ListenAndServe(":3001", routes)
}

type Music struct {
	Title string
	Artist string
}

var musics = []Music{
	Music{Title: "Brain", Artist: "Banks"},
	Music{Title: "Come Down", Artist: "Anderson .Paak"},
}

func getMusics(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(musics)
}

func createMusic(w http.ResponseWriter, r *http.Request) {
	var music Music
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	json.Unmarshal(body, &music);
	musics = append(musics, music)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(music)
}