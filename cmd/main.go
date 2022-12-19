package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/zohaibsoomro/crud-api/db"
	"github.com/zohaibsoomro/crud-api/model"
)

func main() {
	fileServer := http.FileServer(http.Dir("./pages"))
	db.Movies = append(db.Movies, model.Movie{Id: "56675", Title: "Phir Hera Pheri", Director: "Oshita Kumari", ReleaseDate: 2005})
	db.Movies = append(db.Movies, model.Movie{Id: "86770", Title: "Deewangi", Director: "Akshay Kumar", ReleaseDate: 2002})
	port := os.Getenv("HTTP_PLATFORM_PORT")
	routeHanlder := mux.NewRouter()
	routeHanlder.Handle("/", fileServer)
	routeHanlder.HandleFunc("/movies", db.GetMoviesHandler).Methods("GET")
	routeHanlder.HandleFunc("/movies/{id}", db.GetMovieByIdHandler).Methods("GET")
	routeHanlder.HandleFunc("/movies/create", db.CreateMovieHanlder).Methods("POST")
	routeHanlder.HandleFunc("/movies/update/{id}", db.UpdateMovieHandler).Methods("PUT")
	routeHanlder.HandleFunc("/movies/delete/{id}", db.DeleteMovieHandler).Methods("DELETE")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Server Started...")
	// handleError(http.ListenAndServe("127.0.0.1:"+port, nil))
	if err := http.ListenAndServe("127.0.0.1:"+port, routeHanlder); err != nil {
		log.Fatalf("Error %s", err.Error())
	}
	fmt.Println("Server Stopped.")

}
