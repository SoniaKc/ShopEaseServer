package handlers

import (
	"net/http"
	"shop-ease-server/internal/models"
	"shop-ease-server/internal/storage"

	"github.com/gin-gonic/gin"
)

func AddAdresse(c *gin.Context) {
	var req models.AddAdresseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Bad request": err.Error()})
		return
	}

	if err := storage.AddAdresse(req.Login, req.NomAdresse, req.Numero, req.NomRue, req.CodePostal, req.Ville, req.Pays); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"AddBoutique": "Succeeded to create a new user"})
}
