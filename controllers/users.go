package controllers

import (
	"net/http"
	"recipehub-api/models"

	"github.com/gin-gonic/gin"
)

type CreateUsersInput struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateUsersInput struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func FindUsers(context *gin.Context) {
	var users []models.Users
	models.DB.Find(&users)

	context.JSON(http.StatusOK, gin.H{"data": users})
}

func FindUserById(context *gin.Context) {
	var user models.Users

	if err := models.DB.Where("id = ?", context.Param("id")).First(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": user})
}
