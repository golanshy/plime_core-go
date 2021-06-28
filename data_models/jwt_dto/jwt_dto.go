package jwt_dto

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	TokenType      string `json:"token_type,omitempty"`
	UserId         string `json:"user_id,omitempty"`
	Email          string `json:"email,omitempty"`
	ClientId       string `json:"client_id,omitempty"`
	EmailVerified  bool   `json:"email_verified"`
	MobileVerified bool   `json:"mobile_verified"`
	jwt.StandardClaims
}
