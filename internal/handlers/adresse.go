package handlers

import (
	"net/http"
	"shop-ease-server/internal/models"
	"shop-ease-server/internal/storage"
	"strings"

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
	c.JSON(http.StatusCreated, gin.H{"AddAdresse": "Succeeded to create a new adresse"})
}

func GetAdresse(c *gin.Context) {
	login := c.Query("login")
	nomAdresse := c.Query("nom_adresse")

	if login == "" || nomAdresse == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Les paramètres 'login' et 'nom_adresse' sont requis dans l'URL",
		})
		return
	}

	adresse, err := storage.GetAdresse(login, nomAdresse)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, adresse)
}

func DeleteAdresse(c *gin.Context) {
	login := c.Query("login")
	nomAdresse := c.Query("nom_adresse")

	if login == "" || nomAdresse == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Les paramètres 'login' et 'nom_adresse' sont requis dans l'URL",
		})
		return
	}

	err := storage.DeleteAdresse(login, nomAdresse)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"DeleteAdresse": "Adresse deleted successfully"})
}

func UpdateAdresse(c *gin.Context) {
	var req models.UpdateAdresseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Bad request": err.Error()})
		return
	}

	updates := make(map[string]interface{})
	if req.Numero != "" {
		updates["numero"] = req.Numero
	}
	if req.NomRue != "" {
		updates["nom_rue"] = req.NomRue
	}
	if req.CodePostal != "" {
		updates["code_postal"] = req.CodePostal
	}
	if req.Ville != "" {
		updates["ville"] = req.Ville
	}
	if req.Pays != "" {
		updates["pays"] = req.Pays
	}

	err := storage.UpdateAdresse(req.Login, req.NomAdresse, updates)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update adresse", "details": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"UpdateAdresse": "Adresse updated successfully"})
}
