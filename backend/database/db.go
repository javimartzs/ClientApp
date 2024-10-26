package database

import (
	"clientapp/config"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	dsn := fmt.Sprintf(
		"user=%s password=%s dbname=%s port=%s host=%s",
		config.Vars.DBUser,
		config.Vars.DBPass,
		config.Vars.DBName,
		config.Vars.DBPort,
		config.Vars.DBHost,
	)
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	DB.AutoMigrate(
		&models.User{},
	)
	fmt.Println("Connection to database OK")

}
