package handlers

import (
	"net/http"
	"shop-ease-server/internal/models"
	"shop-ease-server/internal/storage"

	"github.com/gin-gonic/gin"
)

func AddBoutique(c *gin.Context) {
	var req models.AddBoutiqueRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Bad request": err.Error()})
		return
	}

	if err := storage.AddBoutique(req.Login, req.Password, req.Nom, req.Email, req.Telephone, req.Siret, req.Forme_juridique, req.Siege_social, req.Pays_enregistrement, req.Iban); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"AddBoutique": "Succeeded to create a new user"})
}
