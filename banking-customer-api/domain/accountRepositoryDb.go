package domain

import (
	"database/sql"
	"strconv"

	"github.com/DonCuervoC/banking_go_api_hex/errs"
	"github.com/DonCuervoC/banking_go_api_hex/logger"
	"github.com/jmoiron/sqlx"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (d AccountRepositoryDb) SaveAccount(a Account) (*Account, *errs.AppError) {
	// Cambiar la consulta SQL para usar RETURNING para obtener el account_id generado
	query := `
		INSERT INTO accounts (customer_id, opening_date, account_type, amount, status)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING account_id;
	`

	// Usar QueryRow en lugar de Exec para capturar el valor de retorno
	var accountId int64
	err := d.client.QueryRow(query, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status).Scan(&accountId)
	if err != nil {
		logger.Error("Error while creating new account: " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	// Asignar el account_id obtenido al objeto Account
	a.AccountId = strconv.FormatInt(accountId, 10)
	return &a, nil
}

func (d AccountRepositoryDb) FindBy(accountId string) (*Account, *errs.AppError) {
	queryGetAccount := `
		SELECT account_id, customer_id, opening_date, account_type, amount
		FROM accounts
		WHERE account_id = $1;
	`
	var account Account
	err := d.client.Get(&account, queryGetAccount, accountId) // sqlx se encarga de mapear los campos
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("Account not found for ID: " + accountId)
			return nil, errs.NewNotFoundError("Account not found")
		}
		logger.Error("Error while fetching account information: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	return &account, nil
}

func (d AccountRepositoryDb) SaveTransaction(t Transaction) (*Transaction, *errs.AppError) {
	// Iniciar un bloque de transacción
	tx, err := d.client.Begin()
	if err != nil {
		logger.Error("Error while starting a new transaction for bank account transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	// Insertar la transacción bancaria
	queryInsertTransaction := `
		INSERT INTO transactions (account_id, amount, transaction_type, transaction_date)
		VALUES ($1, $2, $3, $4)
		RETURNING transaction_id;
	`
	var transactionId int64
	err = tx.QueryRow(queryInsertTransaction, t.AccountId, t.Amount, t.TransactionType, t.TransactionDate).Scan(&transactionId)
	if err != nil {
		tx.Rollback()
		logger.Error("Error while saving transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	// Actualizar el balance de la cuenta
	var updateAccountQuery string
	if t.IsWithdrawal() {
		updateAccountQuery = `
			UPDATE accounts
			SET amount = amount - $1
			WHERE account_id = $2;
		`
	} else {
		updateAccountQuery = `
			UPDATE accounts
			SET amount = amount + $1
			WHERE account_id = $2;
		`
	}

	_, err = tx.Exec(updateAccountQuery, t.Amount, t.AccountId)
	if err != nil {
		tx.Rollback()
		logger.Error("Error while updating account balance: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	// Confirmar la transacción si todo está bien
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error("Error while committing transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	// Obtener la información actualizada de la cuenta
	account, appErr := d.FindBy(t.AccountId)
	if appErr != nil {
		return nil, appErr
	}

	// Actualizar el objeto de transacción con los nuevos datos
	t.TransactionId = strconv.FormatInt(transactionId, 10)
	t.Amount = account.Amount

	return &t, nil
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{dbClient}
}
