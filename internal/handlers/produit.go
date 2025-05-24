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

	if err := storage.AddProduit(req.LoginBoutique, req.Nom, req.Categories, req.Reduction, req.Prix, req.Description); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"AddProduit": "Produit créé avec succès"})
}

func GetProduit(c *gin.Context) {
	var req models.GetProduitRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Bad request": err.Error()})
		return
	}

	produit, err := storage.GetProduit(req.LoginBoutique, req.Nom)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, produit)

	/*
	   id, err := strconv.Atoi(req)
	   if err != nil {
	       c.JSON(http.StatusBadRequest, gin.H{"Bad request": "ID invalide"})
	       return
	   }

	   produit, err := storage.GetProduit(id)
	   if err != nil {
	       if strings.Contains(err.Error(), "non trouvé") {
	           c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	       } else {
	           c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	       }
	       return
	   }
	   c.JSON(http.StatusOK, produit)*/
}

func DeleteProduit(c *gin.Context) {
	var req models.GetProduitRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Bad request": err.Error()})
		return
	}

	err := storage.DeleteProduit(req.LoginBoutique, req.Nom)
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
	if req.Nom != "" {
		updates["nom"] = req.Nom
	}
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
