package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/goJwt/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}
func Hello(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Hello"})
}
func main() {
	fmt.Println("Hello World")

	r := gin.Default()
	r.GET("/", Hello)
	r.Run()

}
