package handlers

import (
	"net/http"
	"shop-ease-server/internal/models"
	"shop-ease-server/internal/storage"

	"github.com/gin-gonic/gin"
)

func AddParametre(c *gin.Context) {
	var req models.AddParametreRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Bad request": err.Error()})
		return
	}

	if err := storage.AddParametre(req.Login, req.Type, req.Langue, req.Cookies, req.Notifications); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"AddBoutique": "Succeeded to create a new user"})
}
