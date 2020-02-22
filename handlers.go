package main

import (
	"github.com/gin-gonic/gin"
)

func getUsers(c *gin.Context) {

	if usersCache := RedisGet("all_users"); usersCache != nil {
		c.JSON(200, usersCache)
		return
	}

	users := SelectAllUsers()

	c.JSON(200, users)
}

func getUser(c *gin.Context) {

	username := c.Param("username")
	user := SelectUser(username)

	c.JSON(200, user)
}

func postUsers(c *gin.Context) {

	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user = Create(user)

	c.JSON(201, user)
}

func putUsers(c *gin.Context) {
	username := c.Param("username")

	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user.Username = username

	rowsAffected := Update(user)

	c.JSON(200, gin.H{
		"rows_affected": rowsAffected,
	})
}

func deleteUsers(c *gin.Context) {
	username := c.Param("username")

	rowsAffected := Delete(username)

	c.JSON(200, gin.H{
		"rows_affected": rowsAffected,
	})
}
