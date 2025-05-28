package handlers

import (
	"net/http"
	"shop-ease-server/internal/models"
	"shop-ease-server/internal/storage"
	"strings"

	"github.com/gin-gonic/gin"
)

func AddCommentaire(c *gin.Context) {
	var req models.AddCommentaireRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Bad request": err.Error()})
		return
	}

	if err := storage.AddCommentaire(req.LoginBoutique, req.NomProduit, req.IdClient, req.Note, req.Commentaire); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"AddCommentaire": "Succeeded to create a new commentaire"})
}

func GetAllComsProduit(c *gin.Context) {
	loginBoutique := c.Query("login_boutique")
	nomProduit := c.Query("nom_produit")

	if loginBoutique == "" || nomProduit == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Les paramètres 'loginBoutique' et 'nomProduit' sont requis dans l'URL",
		})
		return
	}

	commentaires, err := storage.GetAllComsProduit(loginBoutique, nomProduit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, commentaires)
}

func GetAllComsClient(c *gin.Context) {
	idClient := c.Query("idClient")

	if idClient == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Le paramètre 'idClient' est requis dans l'URL",
		})
		return
	}

	commentaires, err := storage.GetAllComsClient(idClient)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, commentaires)
}

func DeleteCommentaire(c *gin.Context) {
	loginBoutique := c.Query("loginBoutique")
	nomProduit := c.Query("nom_produit")
	idClient := c.Query("idClient")

	if loginBoutique == "" || nomProduit == "" || idClient == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Les paramètres 'loginBoutique', 'nom_produit' et 'idClient' sont requis dans l'URL",
		})
		return
	}

	err := storage.DeleteCommentaire(loginBoutique, nomProduit, idClient)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"DeleteCommentaire": "Commentaire deleted successfully"})
}

func UpdateCommentaire(c *gin.Context) {
	var req models.UpdateCommentaireRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Bad request": err.Error()})
		return
	}

	updates := make(map[string]interface{})
	if req.Note != "" {
		updates["note"] = req.Note
	}
	if req.Commentaire != "" {
		updates["commentaire"] = req.Commentaire
	}

	err := storage.UpdateCommentaire(req.LoginBoutique, req.NomProduit, req.IdClient, updates)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update commentairee", "details": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"UpdateCommentaire": "Commentaire updated successfully"})
}
