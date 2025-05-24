package handlers

import (
	"net/http"
	"shop-ease-server/internal/models"
	"shop-ease-server/internal/storage"
	"strings"

	"github.com/gin-gonic/gin"
)

func AddParametre(c *gin.Context) {
	var req models.AddParametreRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Bad request": err.Error()})
		return
	}

	if err := storage.AddParametre(req.Login, req.Type, req.Langue, req.Cookies, req.Notifications); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"AddParametre": "Succeeded to create a new parametre"})
}

func GetParametre(c *gin.Context) {
	var req models.GetParametreRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Bad request": err.Error()})
		return
	}

	parametre, err := storage.GetParametre(req.Login, req.Type)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, parametre)
}

func DeleteParametre(c *gin.Context) {
	var req models.GetParametreRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Bad request": err.Error()})
		return
	}

	err := storage.DeleteParametre(req.Login, req.Type)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"DeleteParametre": "Parametre deleted successfully"})
}

func UpdateParametre(c *gin.Context) {
	var req models.UpdateParametreRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Bad request": err.Error()})
		return
	}

	updates := make(map[string]interface{})
	if req.Langue != "" {
		updates["langue"] = req.Langue
	}
	if req.Cookies != "" {
		updates["cookies"] = req.Cookies
	}
	if req.Notifications != "" {
		updates["notifications"] = req.Notifications
	}

	err := storage.UpdateParametre(req.Login, req.Type, updates)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update parametres", "details": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"UpdateParametre": "Parametres updated successfully"})
}
