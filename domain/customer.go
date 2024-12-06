package domain

// Customer representa a un cliente en el sistema.
type Customer struct {
	Id          string
	Name        string
	City        string
	ZipCode     string
	DateOfBirth string
	Status      string
	Password    string
}

// ICustomerRepository define el contrato (puerto) para interactuar con los datos del cliente.
type ICustomerRepository interface {
	FindAll() ([]Customer, error) // Método para obtener todos los clientes
}
