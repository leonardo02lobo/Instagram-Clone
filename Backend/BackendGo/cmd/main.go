package main

import (
	"fmt"
	"instagram-clone/config"
	"instagram-clone/routes"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	// Inicializar la conexión a la base de datos
	config.Connect()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4321"}, // <- dominio del frontend
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true, // <- PERMITE ENVIAR COOKIES
		MaxAge:           12 * time.Hour,
	}))

	routes.Rutas(r)

	fmt.Println("Servidor iniciado en http://localhost:8080")
	r.Run(":8080")

}
