package handlers

import (
	"net/http"
	"shop-ease-server/internal/models"
	"shop-ease-server/internal/storage"

	"github.com/gin-gonic/gin"
)

func AddFavori(c *gin.Context) {
	var req models.AddFavorisRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Bad request": err.Error()})
		return
	}

	if err := storage.AddFavori(req.LoginBoutique, req.NomProduit, req.IdClient); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"AddFavori": "Succeeded to create a new favori"})
}

func GetAllFavoris(c *gin.Context) {
	idClient := c.Query("idClient")

	if idClient == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Paramètre 'idClient' requis",
		})
		return
	}

	favoris, err := storage.GetAllFavoris(idClient)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(favoris) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "Favoris vide",
			"data":    []interface{}{},
		})
		return
	}

	c.JSON(http.StatusOK, favoris)
}

func DeleteFavoris(c *gin.Context) {
	idClient := c.Query("idClient")
	loginBoutique := c.Query("login_boutique")
	nomProduit := c.Query("nom_produit")

	if idClient == "" || loginBoutique == "" || nomProduit == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Paramètres 'idClient', 'login_boutique' et 'nom_produit' requis",
		})
		return
	}

	err := storage.DeleteFavoris(loginBoutique, nomProduit, idClient)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Erreur lors de la suppression du favori",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"DeleteFavori": "Succeeded to delete favori"})
}

func DeleteFavorisByProduit(c *gin.Context) {
	loginBoutique := c.Query("login_boutique")
	nomProduit := c.Query("nom_produit")

	if loginBoutique == "" || nomProduit == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Paramètres 'login_boutique' et 'u' requis",
		})
		return
	}

	err := storage.DeleteFavorisByProduit(loginBoutique, nomProduit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Erreur lors de la suppression du favori",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"DeleteFavoriByProduit": "Succeeded to delete favori"})
}

func DeleteFavorisByClient(c *gin.Context) {
	idClient := c.Query("idClient")

	if idClient == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Paramètre 'idClient' requis",
		})
		return
	}

	err := storage.DeleteFavorisByClient(idClient)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Erreur lors de la suppression du favori",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"DeleteFavoriByClient": "Succeeded to delete favori"})
}
