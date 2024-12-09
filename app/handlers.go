package app

import (
	"net/http"

	//"strconv"

	"github.com/DonCuervoC/banking_go_api_hex/service"
	"github.com/gin-gonic/gin"
)

// 3
// CustomerHandlers es el controlador que maneja las solicitudes HTTP relacionadas con clientes.
// Aquí no se realiza lógica de negocio, sino que se delega a la capa de servicios.
type CustomerHandlers struct {
	service service.CustomerService // Se inyecta la lógica de negocio desde el servicio.
}

// getAllCustomer maneja la solicitud GET para obtener todos los clientes.
// Actúa como un "puerto de entrada" que conecta el mundo exterior (HTTP) con la lógica de negocio.
func (ch *CustomerHandlers) getAllCustomer(c *gin.Context) {
	// Llamamos al servicio para obtener la lista de clientes.
	// El servicio se encarga de la lógica de negocio y de comunicarse con el repositorio.
	customers, err := ch.service.GetAllCustomer()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Determinamos el formato de la respuesta (XML o JSON) según el encabezado recibido.
	if c.GetHeader("Content-Type") == "application/xml" {
		// Responder con XML
		c.XML(http.StatusOK, customers)
	} else {
		// Responder con JSON
		c.JSON(http.StatusOK, customers)
	}
}

func (ch *CustomerHandlers) getCustomer(c *gin.Context) {
	// Obtenemos directamente el customer_id como string
	Id := c.Param("customer_id")
	// fmt.Println("getCustomer")

	if Id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"}) // Retornamos un error 400 si el ID es inválido o vacío.
		return
	}

	// Llamamos al servicio con el Id (string)
	customer, err := ch.service.GetCustomer(Id)

	if err != nil || customer.Id == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"}) // Retornamos un error 404 si no encontramos al usuario.
		return
	}

	// Determinamos el formato de la respuesta (XML o JSON) según el encabezado recibido.
	if c.GetHeader("Content-Type") == "application/xml" {
		// Responder con XML
		c.XML(http.StatusOK, customer)
	} else {
		// Responder con JSON
		c.JSON(http.StatusOK, customer)
	}

}
