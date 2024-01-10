package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/chalermphanFCC/jwt-auth/auth"
	"gitlab.com/chalermphanFCC/jwt-auth/database"
	"gitlab.com/chalermphanFCC/jwt-auth/models"
)

type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GenerateToken(context *gin.Context) {
	var request TokenRequest
	var user models.User

	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	record := database.Instance.Where("email = ?", request.Email).First(&user)
	if record.Error != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}

	tokenString, err := auth.GenerateToken(user.UserName, user.Email)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, gin.H{"token": tokenString})
}
