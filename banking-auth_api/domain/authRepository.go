package domain

import (
	"database/sql"

	"github.com/DonCuervoC/banking_go_api_hex/errs"
	"github.com/DonCuervoC/banking_go_api_hex/logger"
	"github.com/jmoiron/sqlx"
)

type AuthRepository interface {
	FindCustomerByEmail(useremail string, password string) (*Login, *errs.AppError)
	GenerateAndSaveRefreshTokenToStore(authToken AuthToken) (string, *errs.AppError)
	RefreshTokenExists(refreshToken string) *errs.AppError
}

type AuthRepositoryDb struct {
	client *sqlx.DB
}

func (d AuthRepositoryDb) FindCustomerByEmail(useremail, password string) (*Login, *errs.AppError) {
	var login Login

	postSQLVerify := `SELECT u.email AS email, u.customer_id AS customer_id, u.role AS role, COALESCE(string_agg(a.account_id::TEXT, ','), '') AS account_numbers FROM users u LEFT JOIN accounts a ON a.customer_id = u.customer_id WHERE u.email = $1 AND u.password = $2 GROUP BY u.customer_id, u.email, u.role;`

	// logger.Info("SQL Query: " + postSQLVerify)
	// logger.Info("Parameters: email = " + useremail + " password = " + password)

	err := d.client.Get(&login, postSQLVerify, useremail, password)

	if err != nil {

		logger.Error("Query failed: " + err.Error())

		if err == sql.ErrNoRows {
			// Error cuando no se encuentran credenciales válidas
			return nil, errs.NewAuthenticationError("invalid credentials")
		} else {
			// Error inesperado de base de datos
			logger.Error("Error while verifying login request from database: " + err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}
	}

	return &login, nil
}

func (d AuthRepositoryDb) RefreshTokenExists(refreshToken string) *errs.AppError {
	sqlSelect := "select refresh_token from refresh_token_store where refresh_token = $1"
	var token string
	err := d.client.Get(&token, sqlSelect, refreshToken)
	if err != nil {
		if err == sql.ErrNoRows {
			return errs.NewAuthenticationError("refresh token not registered in the store")
		} else {
			logger.Error("Unexpected database error: " + err.Error())
			return errs.NewUnexpectedError("unexpected database error")
		}
	}
	return nil
}

func (d AuthRepositoryDb) GenerateAndSaveRefreshTokenToStore(authToken AuthToken) (string, *errs.AppError) {
	// generate the refresh token
	var appErr *errs.AppError
	var refreshToken string
	if refreshToken, appErr = authToken.newRefreshToken(); appErr != nil {
		return "", appErr
	}

	// store it in the store
	sqlInsert := "insert into refresh_token_store (refresh_token) values ($1)"
	_, err := d.client.Exec(sqlInsert, refreshToken)
	if err != nil {
		logger.Error("unexpected database error: " + err.Error())
		return "", errs.NewUnexpectedError("unexpected database error")
	}
	return refreshToken, nil
}

func NewAuthRepository(client *sqlx.DB) AuthRepositoryDb {
	return AuthRepositoryDb{client}
}
