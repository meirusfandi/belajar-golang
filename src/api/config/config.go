package config

import (
	"api/models"

	"github.com/jinzhu/gorm"
)

// DBInit create connection to database
func DBInit() *gorm.DB {
	db, err := gorm.Open("mysql", "root@tcp(127.0.0.1:3306)/golang")
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(models.User{})
	return db
}
