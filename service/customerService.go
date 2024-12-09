package service

import "github.com/DonCuervoC/banking_go_api_hex/domain"

//2.1 Conectar los puertos o contratos
// CustomerService define el contrato (puerto) de los servicios de cliente.
type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, error) // Método para obtener todos los clientes
	GetCustomer(string) (*domain.Customer, error)
}

//2.2 implementar
// DefaultCustomerService implementa CustomerService utilizando un repositorio.
type DefaultCustomerService struct {
	repo domain.ICustomerRepository // Repositorio inyectado
}

// GetAllCustomer llama al repositorio para obtener los clientes.
func (s DefaultCustomerService) GetAllCustomer() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

// GetCustomer llama al repositorio para obtener a un solo cliente.
func (s DefaultCustomerService) GetCustomer(id string) (*domain.Customer, error) {
	return s.repo.FindById(id)
}

// NewCustomerService crea una nueva instancia del servicio con un repositorio inyectado.
func NewCustomerService(repository domain.ICustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}
