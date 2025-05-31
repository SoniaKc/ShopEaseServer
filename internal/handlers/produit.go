package handlers

import (
	"net/http"
	"shop-ease-server/internal/models"
	"shop-ease-server/internal/storage"
	"strings"

	"github.com/gin-gonic/gin"
)

func AddProduit(c *gin.Context) {
	var req models.AddProduitRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Bad request": err.Error()})
		return
	}

	if err := storage.AddProduit(req.LoginBoutique, req.Nom, req.Categories, req.Reduction, req.Prix, req.Description, req.Image); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"AddProduit": "Produit créé avec succès"})
}

func GetProduit(c *gin.Context) {
	loginBoutique := c.Query("login_boutique")
	nom := c.Query("nom")

	if loginBoutique == "" || nom == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Les paramètres 'loginBoutique' et 'nom' sont requis dans l'URL",
		})
		return
	}

	produit, err := storage.GetProduit(loginBoutique, nom)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, produit)
}

func GetPopulaires(c *gin.Context) {
	produit, err := storage.GetPopulaires()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, produit)
}

func GetAllProduit(c *gin.Context) {
	loginBoutique := c.Query("login_boutique")

	if loginBoutique == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Paramètre 'login_boutique' requis",
		})
		return
	}

	produits, err := storage.GetAllProduit(loginBoutique)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(produits) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "Aucun produit",
			"data":    []interface{}{},
		})
		return
	}

	c.JSON(http.StatusOK, produits)
}

func DeleteProduit(c *gin.Context) {
	loginBoutique := c.Query("login_boutique")
	nom := c.Query("nom")

	if loginBoutique == "" || nom == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Les paramètres 'loginBoutique' et 'nom' sont requis dans l'URL",
		})
		return
	}

	err := storage.DeleteProduit(loginBoutique, nom)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"DeleteProduit": "Produit deleted successfully"})
}

func UpdateProduit(c *gin.Context) {
	var req models.UpdateProduitRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Bad request": err.Error()})
		return
	}

	updates := make(map[string]interface{})
	/*if req.Nom != "" {
		updates["nom"] = req.Nom
	}*/
	if req.Categories != "" {
		updates["categories"] = req.Categories
	}
	if req.Reduction != "" {
		updates["reduction"] = req.Reduction
	}
	if req.Prix != "" {
		updates["prix"] = req.Prix
	}
	if req.Description != "" {
		updates["description"] = req.Description
	}
	if req.Image != nil {
		updates["image"] = req.Image
	}

	err := storage.UpdateProduit(req.LoginBoutique, req.Nom, updates)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update parametres", "details": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"UpdateProduit": "Produit updated successfully"})
}
