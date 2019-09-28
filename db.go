package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

const EqualOperator = " = ?"
const GreaterOrEqualOperator = " >= ?"
const LikeOperator = " LIKE ?"

var DataBase *gorm.DB

func init() {
	db, err := gorm.Open("sqlite3", "database.db")
	if err != nil {
		fmt.Print(err)
		panic("failed to connect to database")
	}
	db.AutoMigrate(&Actors{}, &Movies{})

	DataBase = db
}

func createActor(actor *Actor) {
	DataBase.Create(&actor)
}

func createMovie(movie *Movie) {
	DataBase.Create(&movie)
}

func searchActors(attr string, operator string, val string) *Actors {
	var actors Actors
	DataBase.Where(attr+operator, val).Find(&actors)
	return &actors
}

func searchMovies(attr string, operator string, val string) *Movies {
	var movies Movies
	DataBase.Where(attr+operator, val).Find(&movies)
	return &movies
}

func magicSearch() *Movies {
	var movies Movies
	DataBase.Table("movies").Select("*").
		Joins("inner join actors on movies.id = actors.movie_id AND actors.name = ? AND movies.director = ?", "Uma Thurman", "Quentin Tarantino").
		Scan(&movies)
	return &movies
}
