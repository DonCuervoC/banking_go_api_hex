package app

import (
	"log"

	"github.com/DonCuervoC/banking_go_api_hex/domain"
	"github.com/DonCuervoC/banking_go_api_hex/service"
	"github.com/gin-gonic/gin"
)

func Start() {

	// Inicializa un enrutador de Gin
	router := gin.Default()

	// Inyección de dependencias:
	// Creamos un repositorio simulado (adaptador derecho) y lo pasamos al servicio.
	customerHandlers := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	// Definimos una ruta GET para obtener todos los clientes.
	router.GET("/customers", customerHandlers.getAllCustomer)

	if err := router.Run("localhost:8000"); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
