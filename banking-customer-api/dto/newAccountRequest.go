package dto

import (
	"strings"

	"github.com/DonCuervoC/banking_go_api_hex/errs"
)

type NewAccountRequestDto struct {
	CustomerId  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

func (r NewAccountRequestDto) Validate() *errs.AppError {
	if r.Amount < 5000 {
		return errs.NewValidationError("To open a new account, you need to deposit at least 5000.00")
	}

	// Lista de tipos de cuenta válidos
	validAccountTypes := []string{
		"saving",
		"cheking",
	}

	// Validar si el tipo de cuenta está en la lista de tipos válidos
	accountType := strings.ToLower(r.AccountType)
	for _, validType := range validAccountTypes {
		if accountType == validType {
			return nil
		}
	}

	// Si no encontramos un tipo válido
	return errs.NewValidationError("Account type must be saving, cheking, visa, mastercard, etc.")
}
