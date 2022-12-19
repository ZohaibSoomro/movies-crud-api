package db

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/zohaibsoomro/crud-api/model"
)

var Movies []model.Movie

func GetMoviesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "GET" {
		http.Error(w, "Request method not supported", http.StatusMethodNotAllowed)
		return
	}
	if err := json.NewEncoder(w).Encode(Movies); err != nil {
		return
	}
}

func GetMovieByIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Path[strings.LastIndex(r.URL.Path, "/")+1:]
	println(id)
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
		}
	}()
	found := false
	for _, v := range Movies {
		if strings.EqualFold(v.Id, id) {
			found = true
			if err := json.NewEncoder(w).Encode(v); err != nil {
				http.Error(w, "Some Error occurred!", http.StatusNotFound)
			}
		}
	}
	if !found {
		http.Error(w, "Movie not found!", http.StatusNotFound)
	}

}

func CreateMovieHanlder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie model.Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		http.Error(w, "Some Error occurred!", http.StatusNotFound)
		return
	}
	Movies = append(Movies, movie)
	if err := json.NewEncoder(w).Encode(movie); err != nil {
		http.Error(w, "Some Error occurred!", http.StatusNotFound)
		return
	}
}
func UpdateMovieHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Path[strings.LastIndex(r.URL.Path, "/")+1:]
	println(id)
	var movie model.Movie

	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		http.Error(w, "Invalid Request Body!", http.StatusNotAcceptable)
		return
	}
	movie.Id = id
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
		}
	}()

	for index, v := range Movies {
		if strings.EqualFold(v.Id, id) {
			Movies = append(Movies[:index], Movies[index+1:]...)
			Movies = append(Movies, movie)
			movie.Id = id
			json.NewEncoder(w).Encode(Movies)
			return
		}
	}
}
func DeleteMovieHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Path[strings.LastIndex(r.URL.Path, "/")+1:]
	println(id)
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
		}
	}()

	for index, v := range Movies {
		if strings.EqualFold(v.Id, id) {
			Movies = append(Movies[:index], Movies[index+1:]...)
			if err := json.NewEncoder(w).Encode(Movies); err != nil {
				http.Error(w, "Some Error occurred!", http.StatusNotFound)
			}
			return
		}
	}
	http.Error(w, "Movie not found!", http.StatusNotFound)
}
