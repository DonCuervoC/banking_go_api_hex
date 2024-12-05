package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

type CustomerXML_JSON struct {
	Name    string `json:"full_name" xml:"first_name"`
	City    string `json:"city" xml:"city"`
	ZipCode string `json:"zip_code" xml:"zip_code"`
}

func getAllCustomersXMLJSON(w http.ResponseWriter, r *http.Request) {
	customers := []CustomerXML_JSON{
		{Name: "Cristiano", City: "Madrid", ZipCode: "J5T78G"},
		{Name: "Kylian", City: "Madrid", ZipCode: "4T86JT"},
		{Name: "Iker", City: "Madrid", ZipCode: "J6759S"},
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Set("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}
