package app

import (
	//"fmt"
	"net/http"
	//"strconv"

	"github.com/DonCuervoC/banking_go_api_hex/service"
	"github.com/gin-gonic/gin"
)

// CustomerHandlers define los controladores HTTP para las operaciones de cliente.
type CustomerHandlers struct {
	service service.CustomerService // Inyecta la lógica de negocio
}

// getAllCustomer maneja solicitudes para obtener todos los clientes.
func (ch *CustomerHandlers) getAllCustomer(c *gin.Context) {
	// Llama al servicio para obtener los clientes
	customers, _ := ch.service.GetAllCustomer()

	// Responde con XML o JSON dependiendo del encabezado solicitado
	if c.GetHeader("Content-Type") == "application/xml" {
		// Responder con XML
		c.XML(http.StatusOK, customers)
	} else {
		// Responder con JSON
		c.JSON(http.StatusOK, customers)
	}
}
