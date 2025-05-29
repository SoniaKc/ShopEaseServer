package handlers

import (
	"fmt"
	"net/http"
	"shop-ease-server/internal/models"
	"shop-ease-server/internal/storage"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func AddPanier(c *gin.Context) {
	var req models.AddPanierRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Bad request": err.Error()})
		return
	}

	quantite, err := strconv.Atoi(req.Quantite)
	if err != nil || quantite <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "La quantité doit être un nombre positif",
		})
		return
	}

	err = storage.AddPanier(req.LoginBoutique, req.NomProduit, req.IdClient, quantite)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"AddPanier": "Succeeded to create a new panier"})
}

func GetQteInPanier(c *gin.Context) {
	loginBoutique := c.Query("login_boutique")
	nomProduit := c.Query("nom_produit")
	idClient := c.Query("idClient")

	if loginBoutique == "" || nomProduit == "" || idClient == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Les paramètres 'idProduit' et 'idClient' sont requis",
		})
		return
	}

	quantite, err := storage.GetQteInPanier(loginBoutique, nomProduit, idClient)
	if err != nil {
		if strings.Contains(err.Error(), "n'est pas dans ce panier") {
			c.JSON(http.StatusNotFound, gin.H{
				"error":   err.Error(),
				"details": fmt.Sprintf("Produit %s non trouvé dans le panier du client %s", loginBoutique, nomProduit, idClient),
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Erreur lors de la récupération de la quantité",
				"details": err.Error(),
			})
		}
		return
	}

	c.JSON(http.StatusOK, quantite)
}

func GetFullPanier(c *gin.Context) {
	idClient := c.Query("idClient")

	if idClient == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Paramètre 'idClient' requis",
		})
		return
	}

	panier, err := storage.GetFullPanier(idClient)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(panier) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "Panier vide",
			"data":    []interface{}{},
		})
		return
	}

	c.JSON(http.StatusOK, panier)
}

func DeletePanier(c *gin.Context) {
	idClient := c.Query("idClient")
	loginBoutique := c.Query("login_boutique")
	nomProduit := c.Query("nom_produit")

	if idClient == "" || loginBoutique == "" || nomProduit == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Paramètres 'idClient' et 'idProduit' requis",
		})
		return
	}

	err := storage.DeletePanier(loginBoutique, nomProduit, idClient)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Erreur lors de la suppression du produit",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"DeletePanier": "Succeeded to delete panier"})
}

func UpdateQteInPanier(c *gin.Context) {
	var req models.UpdatePanierRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Bad request": err.Error()})
		return
	}

	quantite, err := strconv.Atoi(req.Quantite)
	if err != nil || quantite < 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "La quantité doit être un nombre positif ou nul",
		})
		return
	}

	err = storage.UpdateQteInPanier(req.LoginBoutique, req.NomProduit, req.IdClient, quantite)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Erreur lors de la mise à jour du panier",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"UpdateQteInPanier": "Succeeded to update qte in panier"})
}
