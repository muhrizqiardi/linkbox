package auth

import "github.com/golang-jwt/jwt/v5"

type TokenClaims struct {
	UserID int `json:"userId"`
	jwt.RegisteredClaims
}
