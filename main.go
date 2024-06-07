package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}
func getMovie(w http.ResponseWriter, r *http.Request)    {}
func createMovie(w http.ResponseWriter, r *http.Request) {}
func updateMovie(w http.ResponseWriter, r *http.Request) {}
func deleteMovie(j http.ResponseWriter, r *http.Request) {}

var movies []Movie

func main() {
	router := mux.NewRouter()

	movies = append(movies, Movie{
		ID:    "1",
		Isbn:  "4232422",
		Title: "Movie One",
		Director: &Director{
			Firstname: "João Paulo",
			Lastname:  "Rodrigues",
		},
	})
	movies = append(movies, Movie{
		ID:    "2",
		Isbn:  "3539235",
		Title: "Movie Two",
		Director: &Director{
			Firstname: "João Paulo",
			Lastname:  "Rodrigues",
		},
	})
	router.HandleFunc("/movies", getMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	router.HandleFunc("/movies", createMovie).Methods("POST")
	router.HandleFunc("/movies", updateMovie).Methods("PUT")
	router.HandleFunc("/movies", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 3333...")
	log.Fatal(http.ListenAndServe(":3333", router))
}
