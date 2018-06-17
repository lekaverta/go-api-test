package main

import (
	"encoding/json"
	"database/sql"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
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

func getMusics(w http.ResponseWriter, r *http.Request) {
	var musics []Music

	connectionString := "postgresql://wwg@localhost:26257/wwg?sslmode=disable"
	db, _ :=  sql.Open("postgres", connectionString)
	rows, _ := db.Query("SELECT artist, title FROM musics")
    defer rows.Close()
    for rows.Next() {
        var artist, title string
        rows.Scan(&artist, &title);
        musics = append(musics, Music{Title: title, Artist: artist})
    }

	json.NewEncoder(w).Encode(musics)
}

func createMusic(w http.ResponseWriter, r *http.Request) {
	var music Music
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &music);

	connectionString := "postgresql://wwg@localhost:26257/wwg?sslmode=disable"
	db, _ :=  sql.Open("postgres", connectionString)
	db.Exec("INSERT INTO musics (title, artist) VALUES ('" + music.Title + "', '" + music.Artist + "')");

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(music)
}