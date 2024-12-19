package domain

import (
	"database/sql"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Login struct {
	// Username   string         `db:"username"`
	// Email      string         `json:"email" db:"email"`
	Email      string         ` db:"email"`
	CustomerId sql.NullString `db:"customer_id"`
	Accounts   sql.NullString `db:"account_numbers"`
	Role       string         `db:"role"`
}

func (l Login) ClaimsForAccessToken() AccessTokenClaims {
	if l.Accounts.Valid && l.CustomerId.Valid {
		return l.claimsForUser()
	} else {
		return l.claimsForAdmin()
	}
}

func (l Login) claimsForUser() AccessTokenClaims {
	accounts := strings.Split(l.Accounts.String, ",")
	return AccessTokenClaims{
		CustomerId: l.CustomerId.String,
		Accounts:   accounts,
		// Username:   l.Username,
		Useremail: l.Email,
		Role:      l.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ACCESS_TOKEN_DURATION).Unix(),
		},
	}
}

func (l Login) claimsForAdmin() AccessTokenClaims {
	return AccessTokenClaims{
		// Username: l.Username,
		Useremail: l.Email,
		Role:      l.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ACCESS_TOKEN_DURATION).Unix(),
		},
	}
}
