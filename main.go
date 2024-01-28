package main

import (
	"github.com/gin-gonic/gin"
	"github.com/reynaldineo/Go-JWT-Auth/controllers"
	"github.com/reynaldineo/Go-JWT-Auth/initializers"
	"github.com/reynaldineo/Go-JWT-Auth/middleware"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.GET("/me", middleware.RequireAuth, controllers.AccountDetails)

	r.Run()
}