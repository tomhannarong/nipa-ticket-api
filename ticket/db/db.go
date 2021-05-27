package db

import (
	"log"
	"ticket/config"
	"ticket/model"
	"ticket/seed"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var db *gorm.DB

// GetDB - call this method to get db
func GetDB() *gorm.DB {
	return db
}

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file ")
	}
}

// SetupDB - setup dabase for sharing to all api
func SetupDB() {

	database := config.DBConfig()

	database.AutoMigrate(&model.Ticket{})
	database.AutoMigrate(&model.User{})
	database.AutoMigrate(&model.Status{})
	database.AutoMigrate(&model.Contact{})

	seed.Load(database)
	db = database
}
