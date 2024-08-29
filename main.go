package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/goJwt/controllers"
	"github.com/goJwt/initializers"
	"github.com/goJwt/middleware"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}
func Hello(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Hello"})
}
func main() {
	fmt.Println("Hello World")

	r := gin.Default()
	r.GET("/", Hello)
	r.POST("/create", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.Run()
}
