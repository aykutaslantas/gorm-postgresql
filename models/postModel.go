package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Name string
	Body string
}