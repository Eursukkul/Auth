package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/chalermphanFCC/jwt-auth/database"
	"gitlab.com/chalermphanFCC/jwt-auth/models"
)

func RegisterUser(context *gin.Context) {
	var user models.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	if err := user.HashPassword(user.Password); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	record := database.Instance.Create(&user)
	if record.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": user})
}
