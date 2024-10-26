package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var AppVars Config // Variable global para almacenar las variables de configuración

type Config struct {
	DBUser string
	DBPass string
	DBName string
	DBPort string
	DBHost string
	JwtKey string
}

// LoadConfig carga las variables de entorno del fichero .env
// Esto incluye las credenciales a la base de datos y JWT secret
func LoadConfig() {

	// Intentamos cargar el archivo
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading env file: %v", err)
	}

	AppVars = Config{
		DBUser: os.Getenv("DB_USER"),
		DBPass: os.Getenv("DB_PASS"),
		DBName: os.Getenv("DB_NAME"),
		DBPort: os.Getenv("DB_PORT"),
		DBHost: os.Getenv("DB_HOST"),
		JwtKey: os.Getenv("JWT_KEY")}
}
