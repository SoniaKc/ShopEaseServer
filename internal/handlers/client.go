package handlers

import (
	"net/http"
	"shop-ease-server/internal/models"
	"shop-ease-server/internal/storage"
	"strings"

	"github.com/gin-gonic/gin"
)

func AddClient(c *gin.Context) {
	var req models.AddClientRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Bad request": err.Error()})
		return
	}

	if err := storage.AddClient(req.Login, req.Password, req.Nom, req.Prenom, req.Email, req.Date_naissance, req.Telephone); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"AddClient": "Succeeded to create a new user"})
}

func GetClient(c *gin.Context) {
	login := c.Query("login")
	if login == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "paramètre 'login' requis"})
		return
	}

	client, err := storage.GetClient(login)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, client)
}

func DeleteClient(c *gin.Context) {
	login := c.Query("login")
	if login == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "paramètre 'login' requis"})
		return
	}

	err := storage.DeleteClient(login)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"DeleteClient": "Client deleted successfully"})
}

func UpdateClient(c *gin.Context) {
	var req models.UpdateClientRequest
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
	if req.Prenom != "" {
		updates["prenom"] = req.Prenom
	}
	if req.Email != "" {
		updates["email"] = req.Email
	}
	if req.DateNaissance != "" {
		updates["date_naissance"] = req.DateNaissance
	}
	if req.Telephone != "" {
		updates["telephone"] = req.Telephone
	}

	err := storage.UpdateClient(req.Login, updates)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update client", "details": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"UpdateClient": "Client updated successfully"})
}
