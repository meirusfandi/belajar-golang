package models

import (
	"database/sql"
	"entities"
)

type UserModels struct {
	DB *sql.DB
}

func (userModels UserModels) findAll() ([]entities.User, error) {
	rows, err := userModels.DB.Query("select * from user")

	if err != nil {
		return nil, err
	} else {
		users := []entities.User{}
		for rows.Next() {
			var id string
			var username string
			var password string
			var name string

			errs := rows.Scan(&id, &username, &password, &name)
			if errs != nil {
				return nil, errs
			} else {
				user := entities.User{id, username, password, name}
				users = append(users, user)
			}
		}
		return users, nil
	}
}
