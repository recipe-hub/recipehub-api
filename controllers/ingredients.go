package controllers

import (
	"net/http"
	"recipehub-api/models"

	"github.com/gin-gonic/gin"
)

type CreateIngredientsInput struct {
	Name string `json:"name" binding:"required"`
}

type UpdateIngredientsInput struct {
	Name string `json:"name"`
}

// GET /ingredients
// Find all ingredients
func FindIngredients(c *gin.Context) {
	var ingredients []models.Ingredients
	models.DB.Find(&ingredients)

	c.JSON(http.StatusOK, gin.H{"data": ingredients})
}

// GET /ingredients/:id
// Find a ingredient
func FindIngredientsById(c *gin.Context) {
	// Get model if exist
	var ingredients models.Ingredients
	if err := models.DB.Where("id = ?", c.Param("id")).First(&ingredients).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": ingredients})
}

// POST /ingredients
// Create new ingredient
func CreateIngredients(c *gin.Context) {
	// Validate input
	var input CreateIngredientsInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create ingredient
	ingredients := models.Ingredients{Name: input.Name}
	models.DB.Create(&ingredients)

	c.JSON(http.StatusOK, gin.H{"data": ingredients})
}

// PATCH /ingredients/:id
// Update a ingredient
func UpdateIngredients(c *gin.Context) {
	// Get model if exist
	var ingredients models.Ingredients
	if err := models.DB.Where("id = ?", c.Param("id")).First(&ingredients).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateIngredientsInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&ingredients).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": ingredients})
}

// DELETE /ingredients/:id
// Delete a ingredient
func DeleteIngredients(c *gin.Context) {
	// Get model if exist
	var ingredients models.Ingredients
	if err := models.DB.Where("id = ?", c.Param("id")).First(&ingredients).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&ingredients)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
