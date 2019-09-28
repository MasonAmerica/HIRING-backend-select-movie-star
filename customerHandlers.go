package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Actor related attr
const idAttr = "id"
const imdbAttr = "imdb_id"
const nameAttr = "name"

// Movie related attr
const titleAttr = "title"
const genresAttr = "genres"
const imdbScoreAttr = "imdb_score"

func MovieSearchHandler(w http.ResponseWriter, r *http.Request) {
	var movies *Movies
	title := r.FormValue(titleAttr)
	genres := r.FormValue(genresAttr)
	imdbScore := r.FormValue(imdbScoreAttr)
	if title != "" {
		movies = searchMovies(titleAttr, EqualOperator, title)
	} else if genres != "" {
		movies = searchMovies(genresAttr, LikeOperator, "%"+genres+"%")
	} else if imdbScore != "" {
		movies = searchMovies(imdbScoreAttr, GreaterOrEqualOperator, imdbScore)
	} else {
		http.Error(w, "Invalid search paramter", 400)
	}
	outputJson(w, movies)
}

func ActorSearchHandler(w http.ResponseWriter, r *http.Request) {
	var actors *Actors
	id := r.FormValue(idAttr)
	imdb := r.FormValue(imdbAttr)
	name := r.FormValue(nameAttr)
	if id != "" {
		actors = searchActors(idAttr, EqualOperator, id)
	} else if imdb != "" {
		actors = searchActors(imdbAttr, EqualOperator, imdb)
	} else if name != "" {
		actors = searchActors(nameAttr, EqualOperator, name)
	} else {
		http.Error(w, "Invalid search paramter", 400)
	}
	outputJson(w, actors)
}

// Magic search all movies directed by Quentin Tarantino that has Uma Thurman
func MagicSearchHandler(w http.ResponseWriter, r *http.Request) {
	outputJson(w, magicSearch())
}

func outputJson(w http.ResponseWriter, result interface{}) {
	if result == nil {
		fmt.Fprint(w, "[]")
		return
	}
	parsedJson, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(w, string(parsedJson))
}
