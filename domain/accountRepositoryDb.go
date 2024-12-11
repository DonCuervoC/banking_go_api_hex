package domain

import (
	"strconv"

	"github.com/DonCuervoC/banking_go_api_hex/errs"
	"github.com/DonCuervoC/banking_go_api_hex/logger"
	"github.com/jmoiron/sqlx"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (d AccountRepositoryDb) SaveAccount01(a Account) (*Account, *errs.AppError) {
	query := `
		INSERT INTO accounts (customer_id, opening_date, account_type, amount, status)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING account_id;
	`
	result, err := d.client.Exec(query, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	// Comprobar si hubo un error al ejecutar la consulta SQL
	if err != nil {
		logger.Error("Error while creating new account: " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last insert id for account: " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	a.AccountId = strconv.FormatInt(id, 10)
	return &a, nil
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

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{dbClient}
}
