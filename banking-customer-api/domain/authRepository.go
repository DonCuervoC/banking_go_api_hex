package domain

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/DonCuervoC/banking_go_api_hex/logger"
)

type AuthRepository interface {
	IsAuthorized(token string, routeName string, vars map[string]string) bool
}

type RemoteAuthRepository struct {
}

func (r RemoteAuthRepository) IsAuthorized01(token string, routeName string, vars map[string]string) bool {

	u := buildVerifyURL(token, routeName, vars)
	fmt.Println("*****************************")
	fmt.Println("Auth Repository")
	fmt.Println("verify url bulded : ", u)
	// fmt.Println("URL for token verification:" + u)

	if response, err := http.Get(u); err != nil {
		fmt.Println("Error while sending..." + err.Error())
		return false
	} else {

		fmt.Println("XXXX *****************************")
		fmt.Println("Response status from auth server:" + response.Status)
		fmt.Println("XXXX *****************************")
		fmt.Println("Response body from auth server:", response.Body)
		fmt.Println("XXXX *****************************")

		m := map[string]bool{}

		if err = json.NewDecoder(response.Body).Decode(&m); err != nil {
			logger.Error("Error while decoding response from auth server:" + err.Error())
			return false
		}
		return m["isAuthorized"]
	}
}

func (r RemoteAuthRepository) IsAuthorized(token string, routeName string, vars map[string]string) bool {

	// Construir la URL de verificación
	u := buildVerifyURL(token, routeName, vars)
	fmt.Println("*****************************")
	fmt.Println("Auth Repository")
	fmt.Println("verify URL built: ", u)

	// Realizar la solicitud HTTP GET
	response, err := http.Get(u)
	if err != nil {
		fmt.Println("Error while sending request to auth server:", err.Error())
		return false
	}
	defer response.Body.Close()

	// Leer el cuerpo de la respuesta como texto
	bodyBytes, err := io.ReadAll(response.Body)

	if err != nil {
		fmt.Println("Error reading response body:", err.Error())
		return false
	}

	// Registrar el cuerpo de la respuesta para depuración
	bodyString := string(bodyBytes)
	fmt.Println("Response body from auth server:", bodyString)

	// Decodificar la respuesta en un mapa de booleanos
	m := map[string]bool{}
	if err = json.Unmarshal(bodyBytes, &m); err != nil {
		fmt.Println("Error while decoding response from auth server:", err.Error())
		return false
	}

	// Verificar si la clave "isAuthorized" existe en el mapa
	if isAuthorized, exists := m["isAuthorized"]; exists {
		return isAuthorized
	}

	// Si la clave no existe, registrar el error y devolver falso
	fmt.Println("'isAuthorized' key not found in response")
	return false
}

/*
This will generate a url for token verification in the below format

/auth/verify?token={token string}

	&routeName={current route name}
	&customer_id={customer id from the current route}
	&account_id={account id from current route if available}

Sample: /auth/verify?token=aaaa.bbbb.cccc&routeName=MakeTransaction&customer_id=2000&account_id=95470
*/

//BANKING API SERVER
// /auth/verify?routeName=GetAllCustomers&token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjdXN0b21lcl9pZCI6IiIsImFjY291bnRzIjpudWxsLCJ1c2VyZV9lbWFpbCI6ImFkbWluQHRlc3QuY29tIiwicm9sZSI6ImFkbWluIiwiZXhwIjoxNzM0NzI2MzA0fQ.cjSnZfHIhEz1Okv8ZsBXetm2ekwYRDtU55dGawuM5P4

// AUTH API SERVER
// /auth/verify?routeName=GetAllCustomers&token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjdXN0b21lcl9pZCI6IiIsImFjY291bnRzIjpudWxsLCJ1c2VyZV9lbWFpbCI6ImFkbWluQHRlc3QuY29tIiwicm9sZSI6ImFkbWluIiwiZXhwIjoxNzM0NzI2MzA0fQ.cjSnZfHIhEz1Okv8ZsBXetm2ekwYRDtU55dGawuM5P4"

func buildVerifyURL(token string, routeName string, vars map[string]string) string {
	u := url.URL{Host: "localhost:8080", Path: "/auth/verify", Scheme: "http"}
	q := u.Query()
	q.Add("token", token)
	q.Add("routeName", routeName)
	for k, v := range vars {
		q.Add(k, v)
	}
	u.RawQuery = q.Encode()
	return u.String()
}

func NewAuthRepository() RemoteAuthRepository {
	return RemoteAuthRepository{}
}
