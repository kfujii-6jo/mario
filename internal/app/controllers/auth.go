package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"
    "mario/internal/app/models"
)

func Login(c *gin.Context) {
    var user models.User
    var input struct {
        Username string `json:"username" binding:"required"`
        Password string `json:"password" binding:"required"`
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user.Username = "test"
    hashedPassword, _ := hashPassword("password123")
    user.Password = hashedPassword

    if input.Username != user.Username || compareHashAndPassword(user.Password, input.Password ) != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func compareHashAndPassword(hashedPassword string, plainPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err
}
