package controllers

import (
	"net/http"
	"recipehub-api/models"

	"github.com/gin-gonic/gin"
)

type CreateCategoriesInput struct {
	Name string `json:"name" binding:"required"`
}

type UpdateCategoriesInput struct {
	Name string `json:"name"`
}

// GET /categories
// Find all categories
func FindCategories(c *gin.Context) {
	var categories []models.Categories
	models.DB.Find(&categories)

	c.JSON(http.StatusOK, gin.H{"data": categories})
}

// GET /categories/:id
// Find a categories
func FindCategoriesById(c *gin.Context) {
	// Get model if exist
	var categories models.Categories
	if err := models.DB.Where("id = ?", c.Param("id")).First(&categories).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": categories})
}

// POST /categories
// Create new categories
func CreateCategories(c *gin.Context) {
	// Validate input
	var input CreateCategoriesInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create categories
	categories := models.Categories{Name: input.Name}
	models.DB.Create(&categories)

	c.JSON(http.StatusOK, gin.H{"data": categories})
}

// PATCH /categories/:id
// Update a categories
func UpdateCategories(c *gin.Context) {
	// Get model if exist
	var categories models.Categories
	if err := models.DB.Where("id = ?", c.Param("id")).First(&categories).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateCategoriesInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&categories).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": categories})
}

// DELETE /categories/:id
// Delete a categories
func DeleteCategories(c *gin.Context) {
	// Get model if exist
	var categories models.Categories
	if err := models.DB.Where("id = ?", c.Param("id")).First(&categories).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&categories)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
