package handlers

import (
	"net/http"
	"shop-ease-server/internal/models"
	"shop-ease-server/internal/storage"

	"github.com/gin-gonic/gin"
)

func AddPaiement(c *gin.Context) {
	var req models.AddPaiementRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Bad request": err.Error()})
		return
	}

	if err := storage.AddPaiement(req.Login, req.NomCarte, req.NomPersonneCarte, req.CVC, req.DateExpiration); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"AddBoutique": "Succeeded to create a new user"})
}
