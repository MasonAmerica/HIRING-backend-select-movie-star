package main

import (
	"github.com/jinzhu/gorm"
)

type Actor struct {
	gorm.Model
	Id      int    `gorm:"column:id;PRIMARY_KEY;NOT NULL"`
	MovieId int    `gorm:"column:movie_id;NOT NULL"`
	ImdbId  string `gorm:"column:imdb_id;type:varchar(20);NOT NULL"`
	Name    string `gorm:"column:name;type:varchar(255);NOT NULL"`
}

type Actors []Actor
