package app

import (
	"log"
	"net/http"
)

func Start() {
	//define routes
	// Registrar la función handler para la ruta "/greet"
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomersXMLJSON)

	//starting server on localhost:8000
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
