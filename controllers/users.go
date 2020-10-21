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
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
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

func CreateUser(context *gin.Context) {
	var input CreateUsersInput

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.Users{
		Name:     input.Name,
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
	}

	models.DB.Create(&user)
	context.JSON(http.StatusOK, gin.H{"data": user})
}

func UpdateUser(context *gin.Context) {
	var user models.Users
	if err := models.DB.Where("id = ?", context.Param("id")).First(&user).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input UpdateUsersInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&user).Updates(input)
	context.JSON(http.StatusOK, gin.H{"data": user})
}

func DeleteUser(context *gin.Context) {
	var user models.Users
	if err := models.DB.Where("id = ?", context.Param("id")).First(&user).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&user)
	context.JSON(http.StatusOK, gin.H{"data": true})
}
