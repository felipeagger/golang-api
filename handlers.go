package main

import (
	"github.com/gin-gonic/gin"
)

func getUsers(c *gin.Context) {
	var users = RedisGet("users")

	c.JSON(200, gin.H{
		"users": users,
	})
}

func postUsers(c *gin.Context) {

	var user User

	c.BindJSON(&user)

	RedisSet("users", user)

	c.JSON(201, gin.H{
		"user": user,
	})
}

func putUsers(c *gin.Context) {
	username := c.Param("username")

	var user User
	c.BindJSON(&user)

	user.Username = username

	RedisSet("users", user)

	c.JSON(200, gin.H{
		"user": user,
	})
}

func deleteUsers(c *gin.Context) {
	username := c.Param("username")

	var user User
	RedisSet("users", user)

	c.JSON(200, gin.H{
		"username": username,
	})
}
