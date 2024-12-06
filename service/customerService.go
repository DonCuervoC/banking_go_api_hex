package service

import "github.com/DonCuervoC/banking_go_api_hex/domain"

// CustomerService define el contrato (puerto) de los servicios de cliente.
type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, error) // Método para obtener todos los clientes
}

// DefaultCustomerService implementa CustomerService utilizando un repositorio.
type DefaultCustomerService struct {
	repo domain.ICustomerRepository // Repositorio inyectado
}

// GetAllCustomer llama al repositorio para obtener los clientes.
func (s DefaultCustomerService) GetAllCustomer() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

// NewCustomerService crea una nueva instancia del servicio con un repositorio inyectado.
func NewCustomerService(repository domain.ICustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}
