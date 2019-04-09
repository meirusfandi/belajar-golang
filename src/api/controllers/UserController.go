package controllers

import (
	"net/http"

	"../models"

	"github.com/gin-gonic/gin"
)

// to get one data with {id}
func (idb *InDB) GetUserById(c *gin.Context) {
	var (
		user   models.User
		result gin.H
	)
	id := c.Param("id")
	err := idb.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		result = gin.H{
			"result": err.Error(),
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": user,
			"count":  1,
		}
	}

	c.JSON(http.StatusOK, result)
}

// get all data in user
func (idb *InDB) GetAllUsers(c *gin.Context) {
	var (
		users  []models.User
		result gin.H
	)

	idb.DB.Find(&users)
	if len(users) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": users,
			"count":  len(users),
		}
	}

	c.JSON(http.StatusOK, result)
}

// create new data to the database
func (idb *InDB) CreateUser(c *gin.Context) {
	var (
		user   models.User
		result gin.H
	)
	username := c.PostForm("username")
	password := c.PostForm("password")
	name := c.PostForm("name")
	user.userame = username
	user.password = password
	user.name = name
	idb.DB.Create(&user)
	result = gin.H{
		"result": user,
	}
	c.JSON(http.StatusOK, result)
}

// update data with {id} as query
func (idb *InDB) UpdateUser(c *gin.Context) {
	id := c.Query("id")
	username := c.PostForm("username")
	password := c.PostForm("password")
	name := c.PostForm("name")
	var (
		user    models.User
		newUser models.User
		result  gin.H
	)

	err := idb.DB.First(&user, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}
	newUser.userame = username
	newUser.password = password
	newUser.name = name
	err = idb.DB.Model(&user).Updates(newUser).Error
	if err != nil {
		result = gin.H{
			"result": "update failed",
		}
	} else {
		result = gin.H{
			"result": "successfully updated data",
		}
	}
	c.JSON(http.StatusOK, result)
}

// delete data with {id}
func (idb *InDB) DeleteUser(c *gin.Context) {
	var (
		user   models.User
		result gin.H
	)
	id := c.Param("id")
	err := idb.DB.First(&user, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}
	err = idb.DB.Delete(&user).Error
	if err != nil {
		result = gin.H{
			"result": "delete failed",
		}
	} else {
		result = gin.H{
			"result": "Data deleted successfully",
		}
	}

	c.JSON(http.StatusOK, result)
}
