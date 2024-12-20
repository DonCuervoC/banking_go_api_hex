package app

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/DonCuervoC/banking_go_api_hex/dto"
	"github.com/DonCuervoC/banking_go_api_hex/logger"
	"github.com/DonCuervoC/banking_go_api_hex/service"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service service.AuthService
}

func (h AuthHandler) NotImplementedHandler(c *gin.Context) {
	writeResponse(c, http.StatusOK, "Handler not implemented...")
}

func (h AuthHandler) Login(c *gin.Context) {

	var loginRequest dto.NewLoginRequestDto

	// Leer el body crudo
	body, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to read request body"})
		return
	}
	// Decodificar el JSON del body al DTO
	err = json.NewDecoder(bytes.NewReader(body)).Decode(&loginRequest)
	if err != nil {
		writeResponse(c, http.StatusBadRequest, "Invalid JSON format: "+err.Error())
		return
	}

	token, appError := h.service.Login(loginRequest)

	if appError != nil {
		writeResponse(c, appError.Code, appError.Message)
	}
	writeResponse(c, http.StatusOK, token)

}

func (h AuthHandler) Verify(c *gin.Context) {
	// Obtener el parámetro 'token' de la Query
	token := c.Query("token")
	routeName := c.Query("routeName")

	log.Printf("01 Auth Handler Verify ****************************************")
	log.Printf("token: %s", token)
	log.Printf("routeName: %s", routeName)

	if token == "" {
		// Si el token no existe, devolver un 403 Forbidden
		writeResponse(c, http.StatusForbidden, notAuthorizedResponse("missing token"))
		return
	}

	if routeName == "" {
		// Si el token no existe, devolver un 403 Forbidden
		writeResponse(c, http.StatusForbidden, notAuthorizedResponse("missing route name"))
		return
	}

	log.Printf("02 Auth Handler Verify ****************************************")
	log.Printf("checking for existing token...")

	// Crear un mapa con los parámetros necesarios
	urlParams := map[string]string{
		"token":     token,
		"routeName": routeName,
	}

	log.Printf("03 Auth Handler Verify ****************************************")
	log.Printf("url params : %s", urlParams)

	// Llamar al servicio para verificar el token
	log.Printf("04 Auth Handler Verify ****************************************")
	log.Printf("Calling to service...")
	appErr := h.service.Verify(urlParams)

	if appErr != nil {
		// Si hay un error, devolver una respuesta de no autorizado
		writeResponse(c, appErr.Code, notAuthorizedResponse(appErr.Message))
		return
	}

	// Si todo va bien, devolver una respuesta de éxito
	writeResponse(c, http.StatusOK, authorizedResponse())
}

func (h AuthHandler) Refresh(c *gin.Context) {
	var refreshRequest dto.RefreshTokenRequest

	// Decodificar el cuerpo JSON
	if err := c.ShouldBindJSON(&refreshRequest); err != nil {
		logger.Error("Error while decoding refresh token request: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Llamar al servicio para refrescar el token
	token, appErr := h.service.Refresh(refreshRequest)
	if appErr != nil {
		c.JSON(appErr.Code, gin.H{"error": appErr.AsMessage()})
		return
	}

	// Responder con el nuevo token
	c.JSON(http.StatusOK, token)
}

func notAuthorizedResponse(msg string) map[string]interface{} {
	return map[string]interface{}{
		"isAuthorized": false,
		"message":      msg,
	}
}

func authorizedResponse() map[string]bool {
	return map[string]bool{"isAuthorized": true}
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
