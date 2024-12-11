package app

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/DonCuervoC/banking_go_api_hex/dto"
	"github.com/DonCuervoC/banking_go_api_hex/service"
	"github.com/gin-gonic/gin"
)

type AccountHandlers struct {
	service service.AccountService
}

func (h AccountHandlers) NewAccount(c *gin.Context) {

	var request dto.NewAccountRequestDto

	// Obtener el customer_id de la URL
	customerId := c.Param("customer_id")
	if customerId == "" {
		writeResponse(c, http.StatusBadRequest, "Customer ID is required")
		return
	}

	// Asignar el customer_id al DTO
	request.CustomerId = customerId

	// Leer el body crudo
	body, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to read request body"})
		return
	}
	// Decodificar el JSON del body al DTO
	err = json.NewDecoder(bytes.NewReader(body)).Decode(&request)
	if err != nil {
		writeResponse(c, http.StatusBadRequest, "Invalid JSON format: "+err.Error())
		return
	}
	// Procesar el servicio
	account, appError := h.service.NewAccount(request)
	if appError != nil {
		writeResponse(c, appError.Code, appError.Message)
		return
	}

	// Respuesta exitosa
	writeResponse(c, http.StatusCreated, account)
}
