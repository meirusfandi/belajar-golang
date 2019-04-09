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
	// router := gin.New()
	// router.Use(gin.Logger())

	router.GET("/user/:id", inDB.GetUserById)
	router.GET("/users", inDB.GetAllUsers)
	router.POST("/user", inDB.CreateUser)
	router.PUT("/user", inDB.UpdateUser)
	router.DELETE("/user/:id", inDB.DeleteUser)
	fmt.Println("Starting on localhost:8080")
	router.Run(":3306")
}
