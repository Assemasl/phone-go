package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"phone-go/database"
	"phone-go/models"
)

func GetPhones(c *gin.Context) {
	var phones []models.Phone
	database.DB.Find(&phones)
	c.JSON(http.StatusOK, phones)
}

func GetPhone(c *gin.Context) {
	id := c.Param("id")
	var phone models.Phone
	if err := database.DB.First(&phone, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Phone not found"})
		return
	}
	c.JSON(http.StatusOK, phone)
}

func CreatePhone(c *gin.Context) {
	var phone models.Phone
	if err := c.ShouldBindJSON(&phone); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&phone)
	c.JSON(http.StatusOK, phone)
}

func UpdatePhone(c *gin.Context) {
	id := c.Param("id")
	var phone models.Phone
	if err := database.DB.First(&phone, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Phone not found"})
		return
	}
	if err := c.ShouldBindJSON(&phone); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Save(&phone)
	c.JSON(http.StatusOK, phone)
}

func DeletePhone(c *gin.Context) {
	id := c.Param("id")
	database.DB.Delete(&models.Phone{}, id)
	c.Status(http.StatusNoContent)
}
