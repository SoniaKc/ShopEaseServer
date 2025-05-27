package handlers

import (
	"net/http"
	"shop-ease-server/internal/models"
	"shop-ease-server/internal/storage"
	"strings"

	"github.com/gin-gonic/gin"
)

func AddVente(c *gin.Context) {
	var req models.AddVenteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Bad request": err.Error()})
		return
	}

	if err := storage.AddVente(req.IdTransaction, req.LoginBoutique, req.NomProduit, req.IdClient, req.Quantite, req.Total, req.Date_vente, req.Statut); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"AddVente": "Succeeded to create a new vente"})
}

func GetAllTransaction(c *gin.Context) {
	idClient := c.Query("idClient")

	if idClient == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Le paramètre 'idClient' est requis",
		})
		return
	}

	transactions, err := storage.GetAllVentesClient(idClient)
	if err != nil {
		if strings.Contains(err.Error(), "aucune transaction") {
			c.JSON(http.StatusOK, gin.H{
				"error": "Le paramètre 'idClient' est requis",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Erreur lors de la récupération des transactions",
				"details": err.Error(),
			})
		}
		return
	}

	c.JSON(http.StatusOK, transactions)
}

func GetAllVentesClient(c *gin.Context) {
	idClient := c.Query("idClient")

	if idClient == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Le paramètre 'idClient' est requis",
		})
		return
	}

	transactions, err := storage.GetAllVentesClient(idClient)
	if err != nil {
		if strings.Contains(err.Error(), "aucune transaction") {
			c.JSON(http.StatusOK, gin.H{
				"data": "aucune transaction",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Erreur lors de la récupération des transactions",
				"details": err.Error(),
			})
		}
		return
	}

	c.JSON(http.StatusOK, transactions)
}

func GetAllVentesBoutique(c *gin.Context) {
	loginBoutique := c.Query("login_boutique")

	if loginBoutique == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Le paramètre 'login_boutique' est requis",
		})
		return
	}

	result, err := storage.GetAllVentesBoutique(loginBoutique)
	if err != nil {
		if strings.Contains(err.Error(), "aucune transaction") {
			c.JSON(http.StatusOK, gin.H{
				"data": "aucune transaction",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Erreur lors de la récupération des ventes",
				"details": err.Error(),
			})
		}
		return
	}

	c.JSON(http.StatusOK, result)
}

func DeleteAllTransaction(c *gin.Context) {
	idTransaction := c.Query("idTransaction")

	if idTransaction == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Le paramètre 'idTransaction' est requis dans l'URL",
		})
		return
	}

	err := storage.DeleteAllTransaction(idTransaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"DeleteAllTransaction": "Transactions deleted successfully"})
}

func UpdateTransactionStatut(c *gin.Context) {
	idTransaction := c.Query("idTransaction")
	statut := c.Query("statut")

	if idTransaction == "" || statut == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Paramètres 'idTransaction' et 'statut' requis",
		})
		return
	}

	err := storage.UpdateTransactionStatut(idTransaction, statut)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Erreur lors de la mise à jour du statut",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"UpdateTransactionStatut": "Succeeded to update statut"})
}
