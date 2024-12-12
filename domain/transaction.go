package domain

import "github.com/DonCuervoC/banking_go_api_hex/dto"

const WITHDRAWAL = "withdrawal"

type Transaction struct {
	TransactionId   string  `db:"transaction_id"`
	AccountId       string  `db:"account_id"`
	Amount          float64 `db:"amount"`
	TransactionType string  `db:"transaction_type"`
	TransactionDate string  `db:"transaction_date"`
}

func (t Transaction) IsWithdrawal() bool {
	// if t.TransactionType == WITHDRAWAL {
	// 	return true
	// }
	// return false
	return t.TransactionType == WITHDRAWAL
}

func (t Transaction) ToDto() dto.TransactionResponseDto {
	return dto.TransactionResponseDto{
		TransactionId:   t.TransactionId,
		AccountId:       t.AccountId,
		Amount:          t.Amount,
		TransactionType: t.TransactionType,
		TransactionDate: t.TransactionDate,
	}
}
