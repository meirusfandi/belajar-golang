package main

import (
	"../config"
	"../controllers"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := config.DBInit()
	inDB := &controllers.InDB{DB: db}

	router := gin.Default()

	router.GET("/user/:id", inDB.GetUserById)
	router.GET("/users", inDB.GetAllUser)
	router.POST("/user/add", inDB.CreateUser)
	router.PUT("/user/update", inDB.UpdateUser)
	router.DELETE("/user/delete/:id", inDB.DeleteUser)
	router.Run(":3306")
}
