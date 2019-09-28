package main

import "github.com/jinzhu/gorm"

type Movie struct {
	gorm.Model
	Id              int     `gorm:"column:id;PRIMARY_KEY;NOT NULL"`
	ImdbId          string  `gorm:"column:imdb_id;type:varchar(20);NOT NULL"`
	Title           string  `gorm:"column:title;type:varchar(255);NOT NULL"`
	Director        string  `gorm:"column:director;type:varchar(255);NOT NULL"`
	Year            int     `gorm:"column:year;type:int;NOT NULL"`
	Rating          string  `gorm:"column:rating;type:varchar(8);NOT NULL"`
	Genres          string  `gorm:"column:genres;type:varchar(255);NOT NULL"`
	Runtime         int     `gorm:"column:runtime;type:int;NOT NULL"`
	Country         string  `gorm:"column:country;type:varchar(255);NOT NULL"`
	Language        string  `gorm:"column:language;type:varchar(255);NOT NULL"`
	ImdbScore       float64 `gorm:"column:imdb_score;type:float;NOT NULL"`
	ImdbVotes       int     `gorm:"column:imdb_votes;type:int;NOT NULL"`
	MetacriticScore float64 `gorm:"column:metacritic_score;type:float;NOT NULL"`
}

type Movies []Movie
