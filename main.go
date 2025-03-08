package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

/*
	`json:"id"`

these are json name tag that help us when we are encoding are decoding them it serve as alias to their field name
*/
type Movie struct {
	ID       string    `json:"id"`
	ISBN     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	// Set JSON response header
	w.Header().Set("Content-Type", "application/json")

	// Convert movies slice to JSON and write to response
	json.NewEncoder(w).Encode(movies)
}
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	// Set JSON response header
	w.Header().Set("Content-Type", "application/json")
	// mux.vars give us url parameter
	params := mux.Vars(r) // this will return map so to access our value we do params["id"]
	for index, items := range movies {
		if items.ID == params["id"] {
			//  in here we used ... to unpack the slice as we cannot pass slice to append but we have to pass invidual values
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	// Set JSON response header
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	// Set JSON response header
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	// strconv.Itoa(...)
	// Converts the random integer to a string.
	// Itoa stands for "Integer to ASCII".
	movie.ID = strconv.Itoa(rand.IntN(100000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	// Set JSON response header
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, items := range movies {
		if items.ID == params["id"] {
			//  in here we used ... to unpack the slice as we cannot pass slice to append but we have to pass invidual values
			movies = append(movies[:index], movies[index+1:]...)
			//  movie that i am adding
			var movie_to_add Movie
			_ = json.NewDecoder(r.Body).Decode(&movie_to_add)
			movie_to_add.ID = params["id"]
			movies = append(movies, movie_to_add)
			json.NewEncoder(w).Encode(movie_to_add)
			return
		}
	}

}
func main() {
	movies = append(movies, Movie{ID: "1", ISBN: "10", Title: "MOVIE ONE", Director: &Director{Firstname: "best", Lastname: "director"}}, Movie{ID: "2", ISBN: "11", Title: "MOVIE TWO", Director: &Director{Firstname: "best", Lastname: "director"}})
	r := mux.NewRouter()
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movie", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movie/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movie/{id}", deleteMovie).Methods("DELETE")
	fmt.Println("STARTING SERVER AT 8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
