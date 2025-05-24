package handlers

import (
	"net/http"
	"shop-ease-server/internal/models"
	"shop-ease-server/internal/storage"
	"strings"

	"github.com/gin-gonic/gin"
)

func AddPaiement(c *gin.Context) {
	var req models.AddPaiementRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Bad request": err.Error()})
		return
	}

	if err := storage.AddPaiement(req.Login, req.NomCarte, req.NomPersonneCarte, req.Numero, req.CVC, req.DateExpiration); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"AddPaiement": "Succeeded to create a new user"})
}

func GetPaiement(c *gin.Context) {
	var req models.GetPaiementRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Bad request": err.Error()})
		return
	}

	paiement, err := storage.GetPaiement(req.Login, req.NomCarte)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, paiement)
}

func DeletePaiement(c *gin.Context) {
	var req models.GetPaiementRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Bad request": err.Error()})
		return
	}

	err := storage.DeletePaiement(req.Login, req.NomCarte)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"DeletePaiement": "Paiement deleted successfully"})
}

func UpdatePaiement(c *gin.Context) {
	var req models.UpdatePaiementRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Bad request": err.Error()})
		return
	}

	updates := make(map[string]interface{})
	if req.NomPersonneCarte != "" {
		updates["nom_personne_carte"] = req.NomPersonneCarte
	}
	if req.Numero != "" {
		updates["numero"] = req.Numero
	}
	if req.CVC != "" {
		updates["cvc"] = req.CVC
	}
	if req.DateExpiration != "" {
		updates["date_expiration"] = req.DateExpiration
	}

	err := storage.UpdatePaiement(req.Login, req.NomCarte, updates)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update paiement", "details": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"UpdatePaiement": "Paiement updated successfully"})
}
