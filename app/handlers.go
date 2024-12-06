package app

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func greet(c *gin.Context) {
	// Responder con un mensaje en formato JSON
	c.JSON(200, gin.H{
		"message": "Hello, welcome!",
	})
}

type CustomerXML_JSON struct {
	ID      string `json:"id" xml:"id"` // Agregamos un campo ID
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	ZipCode string `json:"zip_code" xml:"zip_code"`
}

// Datos de ejemplo para los clientes
var customers = []CustomerXML_JSON{
	{ID: "1", Name: "Cristiano", City: "Madrid", ZipCode: "J5T78G"},
	{ID: "2", Name: "Kylian", City: "Paris", ZipCode: "4T86JT"},
	{ID: "3", Name: "Iker", City: "Madrid", ZipCode: "J6759S"},
}

func getAllCustomersXMLJSON(c *gin.Context) {
	// Verifica el tipo de contenido solicitado (XML o JSON)
	if c.GetHeader("Content-Type") == "application/xml" {
		// Responder con XML
		c.XML(http.StatusOK, customers)
	} else {
		// Responder con JSON
		c.JSON(http.StatusOK, customers)
	}
}

func getCustomerByID(c *gin.Context) {
	id := c.Param("customer_id") // Obtiene el parámetro 'id' de la URL

	for _, customer := range customers {
		if customer.ID == id {
			if c.GetHeader("Content-Type") == "application/xml" {
				// Responder con XML solo el cliente encontrado
				c.XML(http.StatusOK, customer)
				return
			} else {
				// Responder con JSON solo el cliente encontrado
				c.JSON(http.StatusOK, customer)
				return
			}
		}
	}

	// Si no se encuentra, devolver un error
	c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
}

func createCustomer(c *gin.Context) {
	var newCustomer CustomerXML_JSON

	// Bind JSON body to struct
	if err := c.ShouldBindJSON(&newCustomer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	fmt.Println(newCustomer.Name)

	// Validar que Name, City y ZipCode tengan al menos 3 caracteres
	if len(newCustomer.Name) < 3 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name must have at least 3 characters"})
		return
	}
	if len(newCustomer.City) < 3 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "City must have at least 3 characters"})
		return
	}
	if len(newCustomer.ZipCode) < 3 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ZipCode must have at least 3 characters"})
		return
	}

	// Generar un nuevo ID para el cliente
	maxID := 0
	for _, customer := range customers {
		// Asume que los IDs son enteros
		if customerID, err := strconv.Atoi(customer.ID); err == nil && customerID > maxID {
			maxID = customerID
		}
	}
	newCustomer.ID = strconv.Itoa(maxID + 1) // Nuevo ID es el siguiente entero

	// Agregar el nuevo cliente a la lista
	customers = append(customers, newCustomer)

	// Responder con el cliente creado
	c.JSON(http.StatusCreated, newCustomer)
}
