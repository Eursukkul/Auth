package main

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chalermphanFCC/jwt-auth/controller"
	"gitlab.com/chalermphanFCC/jwt-auth/database"
	"gitlab.com/chalermphanFCC/jwt-auth/middlewares"
)

func main() {
	database.Connect("")
	database.Migration()

	router := initRouter()
	router.Run(":8080")

}

func initRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/register", controller.RegisterUser)
	router.POST("/login", controller.GenerateToken)

	auth := router.Group("/auth")
	auth.Use(middlewares.Auth())
	{
		auth.GET("/users", controller.Ping)
	}

	return router
}
