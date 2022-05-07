package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)

}

func getMovieById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, m := range movies {
		if m.ID == params["id"] {
			json.NewEncoder(w).Encode(m)
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {

}
func updateMovie(w http.ResponseWriter, r *http.Request) {

}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
}
func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "48", Title: "movie one", Director: &Director{
		Firstname: "arsalan",
		Lastname:  "karimzad",
	}})

	movies = append(movies, Movie{ID: "2", Isbn: "49", Title: "Movie two", Director: &Director{
		Firstname: "sahand",
		Lastname:  "karimzad",
	}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovieById).Methods("GET")
	r.HandleFunc("/moveis", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Start a server at 8080")

	log.Fatal(http.ListenAndServe(":8000", r))

}
