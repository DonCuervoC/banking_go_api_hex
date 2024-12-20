package app

import (
	"log"
	"net/http"
	"strings"

	"github.com/DonCuervoC/banking_go_api_hex/domain"
	//"github.com/DonCuervoC/banking_go_api_hex/errs"
	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	repo domain.AuthRepository
}

func (a AuthMiddleware) NamedRoute(routeName string, handler gin.HandlerFunc) gin.HandlerFunc {

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			c.Abort() // Detener la ejecución
			return
		}

		token := getTokenFromHeader(authHeader)
		currentRoute := c.FullPath() // Aquí obtenemos la ruta completa

		log.Printf("routeName : %s", routeName)
		log.Printf("currentRoute : %s", currentRoute)

		currentRouteVars := paramsToMap(c.Params) // Convertimos los parámetros de la ruta a un mapa
		log.Printf("currentRouteVars paramsToMap() : %s", currentRouteVars)

		// fmt.Println("01 **********************************************************:")
		// fmt.Println("token:", token)
		// fmt.Println("02 **********************************************************:")
		// fmt.Println("currentRoute:", currentRoute)
		// fmt.Println("03 **********************************************************:")
		// fmt.Println("currentRoute:", currentRouteVars)
		// fmt.Println("04 **********************************************************:")

		// Verificación de autorización utilizando la ruta completa y las variables
		isAuthorized := a.repo.IsAuthorized(token, routeName, currentRouteVars)

		// fmt.Println("usAuthorized:", isAuthorized)
		// fmt.Println("**********************************************************:")
		if !isAuthorized {
			c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
			c.Abort() // Detener la ejecución
			return
		}

		c.Set("routeName", routeName)
		handler(c) // Llamar al controlador si está autorizado
	}
}

// // Middleware function for authorization
// func (a AuthMiddleware) AuthorizationHandler(routeName string) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		authHeader := c.GetHeader("Authorization")

// 		if authHeader != "" {
// 			token := getTokenFromHeader(authHeader)

// 			// Gin provides access to route name through `FullPath`
// 			currentRoute := c.FullPath()

// 			// fmt.Println("currentRoute: ", currentRoute)

// 			currentRouteVars := paramsToMap(c.Params) // Convert gin.Params to map[string]string // params["customer_id"] o params["account_id"]

// 			// Check authorization
// 			isAuthorized := a.repo.IsAuthorized(token, currentRoute, currentRouteVars)

// 			if isAuthorized {
// 				c.Next() // Continue to the next handler
// 			} else {
// 				appError := errs.AppError{Code: http.StatusForbidden, Message: "Unauthorized"}
// 				c.JSON(appError.Code, gin.H{"error": appError.Message})
// 				c.Abort() // Stop further execution
// 			}
// 		} else {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
// 			c.Abort() // Stop further execution
// 		}
// 	}
// }

// Convert gin.Params to map[string]string
func paramsToMap(params gin.Params) map[string]string {
	m := make(map[string]string)
	for _, p := range params {
		m[p.Key] = p.Value
	}
	return m
}

// Extract the token from the Authorization header
func getTokenFromHeader(header string) string {
	/*
	   Token is expected in the format:
	   "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50cyI6W.yI5NTQ3MCIsIjk1NDcyIiw"
	*/
	splitToken := strings.Split(header, "Bearer")
	if len(splitToken) == 2 {
		return strings.TrimSpace(splitToken[1])
	}
	return ""
}
