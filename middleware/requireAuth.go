package middleware

import (
	"github.com/gin-gonic/gin"
)

func RequireAuth(c *gin.Context) {
	//Get the cookie off req

	//Decode / validate it

	//Check the expiration

	//Find the user with the token sub

	//Attach to request

	//Continue
	c.Next()
}
