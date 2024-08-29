package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/goJwt/initializers"
	"github.com/goJwt/models"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}
	c.BindJSON(&body)

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to hash password."})
		return
	}

	user := models.User{Email: body.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to create user."})
		return
	}

	c.JSON(200, gin.H{"message": "User created"})
}

func Login(c *gin.Context) {
	//Get the email and password off the request body
	var body struct {
		Email    string
		Password string
	}
	err := c.Bind(&body)
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid email or password."})
	}

	//Look up requested user
	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid email or password."})
		return
	}

	//compare passed in password with saved password
	bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	//Generate a JWT token

	//return with the user
}
