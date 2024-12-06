package app

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Start() {
	// Crear una nueva instancia del router Gin
	router := gin.Default()

	// Definir las rutas
	router.GET("/greet", greet)                      // Ruta para el saludo
	router.GET("/customers", getAllCustomersXMLJSON) // Ruta para los clientes

	// Iniciar el servidor en el puerto 8000
	if err := router.Run("localhost:8000"); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
