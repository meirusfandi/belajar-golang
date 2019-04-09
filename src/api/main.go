package main

import (
	"api/config"
	"api/controllers"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := config.DBInit()
	inDB := &controllers.InDB{DB: db}

	router := gin.Default()

	router.GET("/user/:Id", inDB.GetUserById)
	router.GET("/users", inDB.GetAllUsers)
	router.POST("/user/add", inDB.CreateUser)
	router.PUT("/user/update", inDB.UpdateUser)
	router.DELETE("/user/delete/:Id", inDB.DeleteUser)
	fmt.Println("Starting on localhost:8080")
	router.Run(":3306")
}
