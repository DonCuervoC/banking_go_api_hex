package domain

import (
	"github.com/DonCuervoC/banking_go_api_hex/dto"
	"github.com/DonCuervoC/banking_go_api_hex/errs"
)

type Account struct {
	AccountId   string
	CustomerId  string
	OpeningDate string
	AccountType string
	Amount      float64
	Status      bool
}

func (a Account) ToNewAccountResponseDto() dto.NewAccountResponseDto {
	return dto.NewAccountResponseDto{
		AccountId: a.AccountId,
	}
}

type AccountRepository interface {
	SaveAccount(Account) (*Account, *errs.AppError)
}
