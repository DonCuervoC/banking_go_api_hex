package domain

//2. Adapter

// CustomerRepositoryStub es un repositorio simulado (adaptador secundario) que implementa `ICustomerRepository`.
// En lugar de conectarse a una base de datos real, devuelve datos "hardcoded".
type CustomerRepositoryStub struct {
	customers []Customer // Lista de clientes predefinidos
}

// FindAll es la implementación del método definido en la interfaz `ICustomerRepository`.
// Devuelve la lista de clientes simulados sin realizar operaciones reales en la base de datos.
func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

// NewCustomerRepositoryStub es un constructor que crea una nueva instancia de `CustomerRepositoryStub`.
// Inicializa una lista de clientes simulados para usarlos en desarrollo o pruebas.
func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customer := []Customer{
		//	{Id: "1001", Name: "Nelson", City: "Montreal", ZipCode: "ZT6LF5", DateOfBirth: "2000-01-01", Status: "1", Password: "Abc#123"},
		//	{Id: "1002", Name: "Alfonso", City: "Montreal", ZipCode: "JK8LR7", DateOfBirth: "2001-02-02", Status: "1", Password: "Abc#123"},
		//	{Id: "1003", Name: "Cristiano", City: "Montreal", ZipCode: "PU1XR0", DateOfBirth: "2002-03-03", Status: "1", Password: "Abc#123"},
	}
	return CustomerRepositoryStub{customers: customer}
}
