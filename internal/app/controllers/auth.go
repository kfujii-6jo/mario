package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "mario/internal/app/models"
    "mario/internal/app/utils"
)

func Login(c *gin.Context) {
    var input struct {
        Username string `json:"username" binding:"required"`
        Password string `json:"password" binding:"required"`
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // var user models.User
    user := models.GetUserByName(input.Username)
    user.Password, _ = utils.HashPassword(user.Password)
    if user == nil {
	    c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
    }

    if utils.CompareHashAndPassword(user.Password, input.Password) != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
        return
    }

    tokenString, err := utils.GenerateJWT(user.Username)
    if err != nil {
    	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
     	return
    }
	c.SetCookie("jwt_token", tokenString, 3600, "/", "", true, true)
    c.JSON(http.StatusOK, gin.H{
    	"message": "Login successful",
    })
}
