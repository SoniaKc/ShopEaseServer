package handlers

import (
	"net/http"
    "shop-ease-server/internal/models"
    "shop-ease-server/internal/storage"

    "github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context){
	var req models.AddUserRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"Bad request": err.Error()})
        return
    }

    if err := storage.AddUser(req.Firstname, req.Lastname); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, gin.H{"UserCreation": "Succeeded to create a new user"})
}