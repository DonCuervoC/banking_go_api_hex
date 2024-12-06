package domain

// CustomerRepositoryStub es un repositorio simulado para clientes.
type CustomerRepositoryStub struct {
	customers []Customer // Lista de clientes predefinidos
}

// FindAll devuelve todos los clientes simulados.
func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

// NewCustomerRepositoryStub crea una nueva instancia del repositorio simulado.
func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customer := []Customer{
		{Id: "1001", Name: "Nelson", City: "Montreal", ZipCode: "ZT6LF5", DateOfBirth: "2000-01-01", Status: "1", Password: "Abc#123"},
		{Id: "1002", Name: "Alfonso", City: "Montreal", ZipCode: "JK8LR7", DateOfBirth: "2001-02-02", Status: "1", Password: "Abc#123"},
		{Id: "1003", Name: "Cristiano", City: "Montreal", ZipCode: "PU1XR0", DateOfBirth: "2002-03-03", Status: "1", Password: "Abc#123"},
	}
	return CustomerRepositoryStub{customers: customer}
}
