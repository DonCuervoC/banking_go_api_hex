package app

import (
	"fmt"
	//"log"
	"net/http"
	"strconv"

	//"strconv"

	//"github.com/DonCuervoC/banking_go_api_hex/errs"
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

	writeResponse(c, http.StatusOK, customers)
}

func (ch *CustomerHandlers) getCustomer(c *gin.Context) {
	// Obtenemos directamente el customer_id como string
	Id := c.Param("customer_id")

	// Validamos si el ID es numérico
	if _, err := strconv.Atoi(Id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID, must be a number"}) // Retornamos un error 400 si el ID no es un número.
		return
	}

	// Llamamos al servicio con el Id (string)
	customer, err := ch.service.GetCustomer(Id)

	// Mejoramos la comprobación de error con la structura de error personalizada
	if err != nil {
		fmt.Println(err.Message)
		//c.JSON(err.Code, gin.H{"error": err.Message}) // Error al obtener los datos.
		writeResponse(c, err.Code, err.AsMessage())
		return
	}
	writeResponse(c, http.StatusOK, customer)
}

func writeResponse(c *gin.Context, statusCode int, data interface{}) {
	// Determinamos el formato de la respuesta (XML o JSON) según el encabezado recibido.
	if c.GetHeader("Content-Type") == "application/xml" {
		// Responder con XML
		c.XML(statusCode, data)
	} else {
		// Responder con JSON
		c.JSON(statusCode, data)
	}
}