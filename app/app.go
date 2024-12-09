package app

import (
	"log"

	"github.com/DonCuervoC/banking_go_api_hex/domain"
	"github.com/DonCuervoC/banking_go_api_hex/service"
	"github.com/gin-gonic/gin"
)

// Start es la función principal que inicia el servidor y configura las rutas.
// Aquí conectamos las diferentes partes de la aplicación (adaptadores primarios y secundarios).
func Start() {

	// Gin es un framework que nos permite manejar solicitudes HTTP (Adaptador Primario).
	router := gin.Default()

	//4.
	// Inyección de dependencias:
	// Creamos un "repositorio" (CustomerRepositoryStub), que es un adaptador secundario, y lo inyectamos en el servicio.
	// Luego pasamos ese servicio al controlador (CustomerHandlers), que maneja las solicitudes HTTP.
	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	// Definimos una ruta HTTP GET para obtener todos los clientes.
	// Esta ruta usa la función `getAllCustomer` del controlador `CustomerHandlers`.
	router.GET("/customers", ch.getAllCustomer)
	router.GET("/customer/:customer_id", ch.getCustomer)

	if err := router.Run("localhost:8000"); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
