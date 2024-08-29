package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/goJwt/initializers"
	"github.com/goJwt/models"
	"github.com/golang-jwt/jwt"
)

func RequireAuth(c *gin.Context) {
	//Get the cookie off req
	authHeader, err := c.GetHeader("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	//Decode / validate it
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		//Check the expiration
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		//Find the user with the token sub
		var user models.User
		initializers.DB.First(&user, claims["sub"])

		//Attach to request

		//Continue
		c.Next()

		fmt.Println(claims["foo"], claims["nbf"])

	} else {
		fmt.Println(err)
	}

}
