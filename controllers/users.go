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

func FindUsers(ginContext *gin.Context) {
	var users []models.Users
	models.DB.Find(&users)

	ginContext.JSON(http.StatusOK, gin.H{"data": users})
}
