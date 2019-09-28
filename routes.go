package main

import (
	"net/http"
)

const movieSearchRouteName = "MovieSearch"
const movieSearchRoutePath = "/movsrch"

const actorSearchRouteName = "ActorSearch"
const actorSearchRoutePath = "/actsrch"

const magicSearchRouteName = "MagicSearch"
const magicSearchRoutePath = "/magicsrch"

const createActorRouteName = "CreateActor"
const createActorRoutePath = "/NEWACT"

const createMovieRouteName = "CreateMovie"
const createMovieRoutePath = "/NEWMOV"

const getMethod = "GET"
const postMethod = "POST"

const customerIdentifier = "C"
const adminIdentifier = "ADMIN"

type Route struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc http.HandlerFunc
	identifier  string
}

type Routes []Route

var AllRoutes = Routes{
	Route{
		movieSearchRouteName,
		getMethod,
		movieSearchRoutePath,
		MovieSearchHandler,
		customerIdentifier,
	},
	Route{
		actorSearchRouteName,
		getMethod,
		actorSearchRoutePath,
		ActorSearchHandler,
		customerIdentifier,
	},
	Route{
		magicSearchRouteName,
		getMethod,
		magicSearchRoutePath,
		MagicSearchHandler,
		customerIdentifier,
	},
	Route{
		createActorRouteName,
		postMethod,
		createActorRoutePath,
		createActorHandler,
		adminIdentifier,
	},
	Route{
		createMovieRouteName,
		postMethod,
		createMovieRoutePath,
		createMovieHandler,
		adminIdentifier,
	},
}
