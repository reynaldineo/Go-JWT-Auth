package main

import (
	"github.com/gin-gonic/gin"
	"github.com/reynaldineo/Go-JWT-Auth/initializers"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}