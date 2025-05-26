package handlers

import (
	"net/http"
	"shop-ease-server/internal/storage"

	"github.com/gin-gonic/gin"
)

func AddFavori(c *gin.Context) {
	idClient := c.Query("idClient")
	idProduit := c.Query("idProduit")

	if idClient == "" || idProduit == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Paramètres 'idClient' et 'idProduit' requis",
		})
		return
	}

	if err := storage.AddFavori(idProduit, idClient); err != nil {
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
	idProduit := c.Query("idProduit")

	if idClient == "" || idProduit == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Paramètres 'idClient' et 'idProduit' requis",
		})
		return
	}

	err := storage.DeleteFavoris(idProduit, idClient)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Erreur lors de la suppression du favori",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"DeleteFavori": "Succeeded to delete favori"})
}
