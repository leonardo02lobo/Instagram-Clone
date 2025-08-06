package config

import (
	"log"
	"os"
)

type AppConfig struct {
	JWTSecret string
}

var Config AppConfig

func LoadConfig() {
	// Cargar el secreto JWT desde variables de entorno
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		// Valor por defecto (SOLO para desarrollo)
		jwtSecret = "LEONARDO-LOBO-SECRET"
		log.Println("ADVERTENCIA: Usando JWT secret por defecto. En producción usa JWT_SECRET en variables de entorno")
	}

	Config = AppConfig{
		JWTSecret: jwtSecret,
	}
}
