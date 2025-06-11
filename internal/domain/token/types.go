package token

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type AuthToken struct {
	ExpiresAt time.Time
	IssuedAt  time.Time
	LogoutPin string
	Token     string
	UserID    string
}

type SingleTokenInput struct {
	ProvidedToken string `json:"provided_token" binding:"required"`
}

// CustomClaims - struct for tokens generation
type CustomClaims struct {
	jwt.RegisteredClaims
	LogoutPin string `json:"logout_pin"`
	UserID    string `json:"user_id"`
}
