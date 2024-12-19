package dto

type NewLoginRequestDto struct {
	// Username string `json:"username"`
	Useremail string `json:"email"`
	Password  string `json:"password"`
}

// type NewLoginRequestDto struct {
// 	Name     string `json:"full_name"`
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// 	// Role     string `json:"role"`
// }

// type NewLoginResponseDto struct {
// 	Name  string `json:"full_name"`
// 	Token string `json:"token"`
// }

// type LoginResponse struct {
// 	AccessToken  string `json:"access_token"`
// 	RefreshToken string `json:"refresh_token,omitempty"`
// }

// type RefreshTokenRequest struct {
// 	AccessToken  string `json:"access_token"`
// 	RefreshToken string `json:"refresh_token"`
// }

// func (r RefreshTokenRequest) IsAccessTokenValid() *jwt.ValidationError {

// 	// 1. invalid token.
// 	// 2. valid token but expired
// 	_, err := jwt.Parse(r.AccessToken, func(token *jwt.Token) (interface{}, error) {
// 		return []byte(domain.HMAC_TKN_SECRET), nil
// 	})
// 	if err != nil {
// 		var vErr *jwt.ValidationError
// 		if errors.As(err, &vErr) {
// 			return vErr
// 		}
// 	}
// 	return nil
// }
