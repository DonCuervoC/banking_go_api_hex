package service

import (
	"github.com/DonCuervoC/banking_go_api_hex/domain"
	"github.com/DonCuervoC/banking_go_api_hex/dto"
	"github.com/DonCuervoC/banking_go_api_hex/errs"
)

// 2.1 Conectar los puertos o contratos
// CustomerService define el contrato (puerto) de los servicios de cliente.
type CustomerService interface {
	// GetAllCustomer() ([]domain.Customer, error) // Método para obtener todos los clientes
	GetAllCustomer(string) ([]dto.CustomerResponseDto, *errs.AppError)
	GetCustomer(string) (*dto.CustomerResponseDto, *errs.AppError)
}

// 2.2 implementar
// DefaultCustomerService implementa CustomerService utilizando un repositorio.
type DefaultCustomerService struct {
	repo domain.ICustomerRepository // Repositorio inyectado
}

// GetAllCustomer llama al repositorio para obtener los clientes.
// func (s DefaultCustomerService) GetAllCustomer() ([]domain.Customer, error) {
func (s DefaultCustomerService) GetAllCustomer(status string) ([]dto.CustomerResponseDto, *errs.AppError) {

	customers, err := s.repo.FindAll(status)
	if err != nil {
		return nil, err
	}

	// customerDtos := domain.ToDtoList(customers)
	customerDtos := domain.ToDtoList(customers)

	return customerDtos, nil
}

// GetCustomer llama al repositorio para obtener a un solo cliente.
func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponseDto, *errs.AppError) {
	c, err := s.repo.FindById(id)

	if err != nil {
		return nil, err
	}

	response := c.ToDto()
	return &response, nil
}

// NewCustomerService crea una nueva instancia del servicio con un repositorio inyectado.
func NewCustomerService(repository domain.ICustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}

type CustomerService01 interface {
	// GetAllCustomer() ([]domain.Customer, error) // Método para obtener todos los clientes
	GetAllCustomer(string) ([]domain.Customer, *errs.AppError)
	GetCustomer(string) (*domain.Customer, *errs.AppError)
}

func (s DefaultCustomerService) GetCustomer01(id string) (*domain.Customer, *errs.AppError) {
	return s.repo.FindById(id)
}

func (s DefaultCustomerService) GetAllCustomer01(status string) ([]domain.Customer, *errs.AppError) {
	return s.repo.FindAll(status)
}
