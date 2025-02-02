package controllers

import (
    "github.com/gin-gonic/gin"
    "event_api/database"
    "event_api/repository"
    "event_api/structs"
	"event_api/middleware"
    "net/http"
    "fmt"
)

// Register 
func Register(c *gin.Context) {
	var user structs.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	hashedPassword, err := middleware.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = hashedPassword
	
	err = repository.AddUser(database.DbConnection, user)
	if err != nil {
        fmt.Println("Error registering user:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

    c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}


// Login endpoint
func Login(c *gin.Context) {
	var user structs.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	foundUser, err := repository.LogUser(database.DbConnection, user.Email)
	if err != nil {
		if err.Error() == "user not found" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
			fmt.Println(err)
		}
		return
	}

	if !middleware.CheckPasswordHash(user.Password, foundUser.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := middleware.GenerateToken(foundUser.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}