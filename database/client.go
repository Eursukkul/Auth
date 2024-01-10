package database

import (
	"log"

	"gitlab.com/chalermphanFCC/jwt-auth/models"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var Instance *gorm.DB
var dbError error

func Connect(connectionString string) {
	Instance, dbError = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if dbError != nil {
		log.Fatal(dbError)
		panic("failed to connect database")
	}

	log.Println("Database connection established")
}

func Migration() {
	Instance.AutoMigrate(&models.User{})
	log.Println("Database migrated")
}