package router

import (
	"github.com/gorilla/mux"
	"github.com/pratikdev/learning-golang/mongoApi/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movie/{id}", controller.GetMovie).Methods("GET")       // returns one movie with given id
	router.HandleFunc("/api/movies", controller.GetAllMovies).Methods("GET")       // returns all movies
	router.HandleFunc("/api/movie/{id}", controller.DeleteMovie).Methods("DELETE") // deletes one movie with given id
	router.HandleFunc("/api/movies", controller.DeleteAllMovies).Methods("DELETE") // deletes all movies
	router.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("PUT")  // marks one movie as watched with the given id
	router.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")        // creates one new movie

	return router
}
