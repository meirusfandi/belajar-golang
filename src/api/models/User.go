package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	id       int
	username string
	password string
	name     string
}
