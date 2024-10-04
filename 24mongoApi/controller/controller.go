package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pratikdev/learning-golang/mongoApi/controller/mongo_connection"
	"github.com/pratikdev/learning-golang/mongoApi/model"
)

func GetMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	movie := mongo_connection.GetMovie(params["id"])

	json.NewEncoder(w).Encode(movie)
}

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	allMovies := mongo_connection.GetAllMovies()

	json.NewEncoder(w).Encode(allMovies)
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	mongo_connection.DeleteMovie(params["id"])

	json.NewEncoder(w).Encode(map[string]string{"message": "Deleted movie successfully"})
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	mongo_connection.DeleteAllMovies()

	json.NewEncoder(w).Encode(map[string]string{"message": "Deleted all movies succesfully"})
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var movie model.Netflix
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		log.Fatal(err)
	}

	mongo_connection.InsertMovie(movie)

	json.NewEncoder(w).Encode(map[string]string{"message": "Create movie successfully"})
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	mongo_connection.UpdateMovie(params["id"])

	json.NewEncoder(w).Encode(map[string]string{"message": "Movie marked as watched successfully"})
}
