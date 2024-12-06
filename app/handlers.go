package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func greet(c *gin.Context) {
	// Responder con un mensaje en formato JSON
	c.JSON(200, gin.H{
		"message": "Hello, welcome!",
	})
}

type CustomerXML_JSON struct {
	Name    string `json:"full_name" xml:"first_name"`
	City    string `json:"city" xml:"city"`
	ZipCode string `json:"zip_code" xml:"zip_code"`
}

func getAllCustomersXMLJSON(c *gin.Context) {
	// Datos de ejemplo para los clientes
	customers := []CustomerXML_JSON{
		{Name: "Cristiano", City: "Madrid", ZipCode: "J5T78G"},
		{Name: "Kylian", City: "Madrid", ZipCode: "4T86JT"},
		{Name: "Iker", City: "Madrid", ZipCode: "J6759S"},
	}

	// Verifica el tipo de contenido solicitado (XML o JSON)
	if c.GetHeader("Content-Type") == "application/xml" {
		// Responder con XML
		c.XML(http.StatusOK, customers)
	} else {
		// Responder con JSON
		c.JSON(http.StatusOK, customers)
	}
}
