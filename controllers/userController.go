package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/reynaldineo/Go-JWT-Auth/initializers"
	"github.com/reynaldineo/Go-JWT-Auth/models"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	// Get email and password from request body
	var body struct {
		Email    string 
		Password string
	}

	if c.Bind(&body) != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fields are empty",
		})
		return
	}

	// Hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error hashing password",
		})
		return
	}

	// Create user
	user := models.User{Email: body.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error creating user",
		})
		return
	}

	// Return
	c.JSON(http.StatusOK, gin.H{
		"message": "User created successfully",
	})
}

func Login(c *gin.Context){
	// Get email and password from request body
	var body struct {
		Email    string 
		Password string
	}

	if c.Bind(&body) != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fields are empty",
		})
		return
	}

	// Check if user exists
	var user models.User // store user data from database
	initializers.DB.Where("email = ?", body.Email).First(&user)

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Invalide email or password",
		})
		return
	}

	// Check if password is correct
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Invalid password",
		})
		return
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"subject": user.ID,
		"exp": time.Now().Add(time.Hour * 1).Unix(), // 1 day
	})
	
	// get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error create token",
		})
		return
	}
	
	// set token to cookie
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600, "", "", false, true)


	c.JSON(http.StatusOK, gin.H{
		"message": "Login success",
	})
}