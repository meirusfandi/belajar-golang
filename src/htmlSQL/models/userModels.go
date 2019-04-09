package models

import (
	"database/sql"
	"entities"
)

type UserModel struct {
	DB *sql.DB 
}

func (userModel UserModel) findAll() ([]entities.User, error){
	rows, err := userModel.DB.Query("select * from user")

	if err != nil {
		return nil, err
	} else {
		users := []entities.User{}
		for rows.Next() {
			var id string
			var username string
			var password string
			var name string
		}
	}
}