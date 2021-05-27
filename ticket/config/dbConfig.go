package config

import (
	"fmt"
	"os"
	"ticket/types"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DBConfig - Set database
func DBConfig() *gorm.DB {
	// mysql "gorm.io/driver/mysql"
	dbConfig := types.DBConfig{
		User:     os.Getenv("MYSQL_USER"),
		Password: os.Getenv("MYSQL_PASSWORD"),
		Host:     os.Getenv("MYSQL_HOST"),
		Port:     os.Getenv("MYSQL_PORT"),
		DBName:   os.Getenv("MYSQL_DATABASE"),
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DBName)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Println("DB connected")
	}
	return database
}
