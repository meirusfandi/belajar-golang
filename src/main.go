package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world")
}

func CallFindAll() {
	db, err := config.getMySQLDB()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		if err != nil {
			fmt.Println(err.Error())
		} else {
			userModels := models.userModels{
				Db: db,
			}
			if err != nil {
				fmt.Println(err)
			} else {

			}
		}
	}
}
