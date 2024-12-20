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

// /customers/2000/accounts/90720
func (h AccountHandlers) MakeTransaction(c *gin.Context) {

	// Obtener el customer_id de la URL
	customerId := c.Param("customer_id")
	account_id := c.Param("account_id")
	if customerId == "" || account_id == "" {
		writeResponse(c, http.StatusBadRequest, "Customer ID or account ID required")
		return
	}

	var request dto.TransactionRequestDto

	request.AccountId = account_id
	request.CustomerId = customerId

	// Leer el body crudo
	body, err := c.GetRawData()

	if err != nil {
		writeResponse(c, http.StatusBadRequest, "Unable to read request body")
		return
	}

	// Decodificar el JSON del body al DTO
	err = json.NewDecoder(bytes.NewReader(body)).Decode(&request)
	if err != nil {
		writeResponse(c, http.StatusBadRequest, "Invalid JSON format: "+err.Error())
		return
	}

	// make transaction
	account, appError := h.service.MakeTransaction(request)

	if appError != nil {
		writeResponse(c, appError.Code, appError.AsMessage())
	} else {
		writeResponse(c, http.StatusOK, account)
	}

}
