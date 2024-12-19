package domain

import (
	"github.com/DonCuervoC/banking_go_api_hex/dto"
	"github.com/DonCuervoC/banking_go_api_hex/errs"
)

//1. Domain object

// Customer es la entidad principal que representa a un cliente en el sistema.
// Esta estructura contiene los datos básicos de un cliente.
type Customer struct {
	Id          string `db:"customer_id"`
	Name        string `db:"name"`
	City        string `db:"city"`
	ZipCode     string `db:"zipcode"`
	DateOfBirth string `db:"date_of_birth"`
	Status      bool   `db:"status"`
}

func (c Customer) statusAsString() string {
	if c.Status {
		return "active"
	}
	return "inactive"
}

func (c Customer) ToDto() dto.CustomerResponseDto {

	return dto.CustomerResponseDto{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		ZipCode:     c.ZipCode,
		DateOfBirth: c.DateOfBirth,
		Status:      c.statusAsString(),
	}

}

func ToDtoList(customers []Customer) []dto.CustomerResponseDto {
	var dtos []dto.CustomerResponseDto
	for _, customer := range customers {
		dtos = append(dtos, customer.ToDto())
	}
	return dtos
}

// 1.1 introduce the contract
// ICustomerRepository es una **interfaz** que define el contrato (o puerto) para los repositorios de clientes.
// Un "puerto" en la arquitectura hexagonal es un punto de conexión entre las capas internas y externas.
type ICustomerRepository interface {
	// Este contrato garantiza que cualquier implementación (base de datos, APIs, etc.) tendrá esta función.

	// FindAll() ([]Customer, error) // Método para obtener todos los clientes
	//FindAll() ([]Customer, *errs.AppError)
	FindAll(string) ([]Customer, *errs.AppError)
	FindById(string) (*Customer, *errs.AppError)
}
