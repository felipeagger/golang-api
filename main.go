package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World Go!",
		})
	})

	router.GET("/users", getUsers)
	router.GET("/users/:username", getUser)
	router.POST("/users", postUsers)
	router.PUT("/users/:username", putUsers)
	router.DELETE("/users/:username", deleteUsers)

	router.Run(":" + os.Getenv("PORT"))

}
