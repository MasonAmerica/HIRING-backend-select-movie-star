package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const invalidRequestBody = "Invalid request body"
const missingAttribute = "One or more of the required attributes are missing"
const successfulCreation = "Successfully create: "

func createActorHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var actor Actor
	err := decoder.Decode(&actor)
	if err != nil {
		http.Error(w, invalidRequestBody, 400)
		return
	}
	if actor.MovieId <= 0 || actor.Name == "" || actor.ImdbId == "" {
		http.Error(w, missingAttribute, 400)
		return
	}
	createActor(&actor)
	fmt.Fprint(w, successfulCreation, actor)
}

func createMovieHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var movie Movie
	err := decoder.Decode(&movie)
	if err != nil {
		http.Error(w, invalidRequestBody, 400)
		return
	}
	if movie.ImdbId == "" || movie.Title == "" || movie.Director == "" || movie.Year <= 0 ||
		movie.Rating == "" || movie.Genres == "" || movie.Runtime <= 0 || movie.Country == "" ||
		movie.Language == "" || movie.ImdbScore <= 0 || movie.ImdbVotes <= 0 || movie.MetacriticScore <= 0 {
		http.Error(w, missingAttribute, 400)
		return
	}
	createMovie(&movie)
	fmt.Fprint(w, successfulCreation, movie)
}
