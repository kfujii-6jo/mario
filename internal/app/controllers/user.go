package controllers

import (
    "mario/internal/app/models"
    "net/http"
    "github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
    users := models.GetAllUsers()
    c.JSON(http.StatusOK, gin.H{"users": users})
}

func CreateUser(c *gin.Context) {
    var newUser models.User

    if err := c.ShouldBindJSON(&newUser); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    models.AddUser(newUser)
    c.JSON(http.StatusCreated, newUser)
}
