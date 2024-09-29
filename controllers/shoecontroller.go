package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shoe_inventory/config"
	"shoe_inventory/models"
)

func GetShoes(c *gin.Context) {
	var shoes []models.Shoe
	config.DB.Find(&shoes)
	c.JSON(http.StatusOK, gin.H{"data": shoes})
}

func CreateShoe(c *gin.Context) {
	var input models.Shoe
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shoe := models.Shoe{Name: input.Name, Size: input.Size, Price: input.Price, Quantity: input.Quantity}
	config.DB.Create(&shoe)

	c.JSON(http.StatusOK, gin.H{"data": shoe})
}

func GetShoeByID(c *gin.Context) {
	var shoe models.Shoe
	if err := config.DB.Where("id = ?", c.Param("id")).First(&shoe).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Shoe not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": shoe})
}

func UpdateShoe(c *gin.Context) {
	var shoe models.Shoe
	if err := config.DB.Where("id = ?", c.Param("id")).First(&shoe).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Shoe not found"})
		return
	}

	var input models.Shoe
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&shoe).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": shoe})
}

func DeleteShoe(c *gin.Context) {
	var shoe models.Shoe
	if err := config.DB.Where("id = ?", c.Param("id")).First(&shoe).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Shoe not found"})
		return
	}

	config.DB.Delete(&shoe)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
