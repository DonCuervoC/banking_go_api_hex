package domain

import (
	"github.com/DonCuervoC/banking_go_api_hex/dto"
	"github.com/DonCuervoC/banking_go_api_hex/errs"
)

type Account struct {
	AccountId   string  `db:"account_id"`
	CustomerId  string  `db:"customer_id"`
	OpeningDate string  `db:"opening_date"`
	AccountType string  `db:"account_type"`
	Amount      float64 `db:"amount"`
	Status      bool    `db:"status"`
}

func (a Account) ToNewAccountResponseDto() dto.NewAccountResponseDto {
	return dto.NewAccountResponseDto{
		AccountId: a.AccountId,
	}
}

type AccountRepository interface {
	SaveAccount(Account) (*Account, *errs.AppError)
	SaveTransaction(transaction Transaction) (*Transaction, *errs.AppError)
	FindBy(accountId string) (*Account, *errs.AppError)
}

func (a Account) CanWithdraw(amount float64) bool {
	return a.Amount >= amount
}
