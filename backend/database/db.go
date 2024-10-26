package database

import (
	"clientapp/config"
	"clientapp/models"

	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB // Definimos una variable global para la base de datos

// ConnectDB establece la conexión con la base de datos PostgreSQL
func ConnectDB() {

	dsn := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s",
		config.AppVars.DBUser,
		config.AppVars.DBPass,
		config.AppVars.DBName,
		config.AppVars.DBPort,
		config.AppVars.DBHost,
	)

	// Intentamos abrir una conexión con la base de datos usando GORM
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// Migramos los modelos a la base de datos: creamos las tablas en la db
	err = DB.AutoMigrate(
		&models.User{},
		&models.Promotion{},
		&models.PointsTransaction{},
	)
}
