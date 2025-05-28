package handlers

import (
	"net/http"
	"shop-ease-server/internal/models"
	"shop-ease-server/internal/storage"
	"strings"

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
	c.JSON(http.StatusCreated, gin.H{"AddBoutique": "Succeeded to create a new boutique"})
}

func GetBoutique(c *gin.Context) {
	login := c.Query("login")
	if login == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "paramètre 'login' requis"})
		return
	}

	boutique, err := storage.GetBoutique(login)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, boutique)
}

func DeleteBoutique(c *gin.Context) {
	login := c.Query("login")
	if login == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "paramètre 'login' requis"})
		return
	}

	err := storage.DeleteBoutique(login)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"DeleteBoutique": "Boutique deleted successfully"})
}

func UpdateBoutique(c *gin.Context) {
	var req models.UpdateBoutiqueRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Bad request": err.Error()})
		return
	}

	updates := make(map[string]interface{})
	if req.Password != "" {
		updates["password"] = req.Password
	}
	if req.Nom != "" {
		updates["nom"] = req.Nom
	}
	if req.Email != "" {
		updates["email"] = req.Email
	}
	if req.Telephone != "" {
		updates["telephone"] = req.Telephone
	}
	if req.Siret != "" {
		updates["siret"] = req.Siret
	}
	if req.Forme_juridique != "" {
		updates["forme_juridique"] = req.Forme_juridique
	}
	if req.Siege_social != "" {
		updates["siege_social"] = req.Siege_social
	}
	if req.Pays_enregistrement != "" {
		updates["pays_enregistrement"] = req.Pays_enregistrement
	}
	if req.Iban != "" {
		updates["iban"] = req.Iban
	}

	err := storage.UpdateBoutique(req.Login, updates)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update boutique", "details": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"UpdateBoutique": "Boutique updated successfully"})
}
