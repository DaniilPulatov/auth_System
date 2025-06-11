package tokenOps

import (
	"auth-service/internal/domain/token"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"os"
	"time"
)

func newToken(rToken *token.AuthToken) (string, error) {
	payload := token.CustomClaims{
		UserID:    rToken.UserID,
		LogoutPin: rToken.LogoutPin,
	}
	payload.IssuedAt = jwt.NewNumericDate(rToken.IssuedAt)
	payload.ExpiresAt = jwt.NewNumericDate(rToken.ExpiresAt)
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return newToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func ValidateToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			err := fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			return nil, errors.New(err.Error())
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		log.Println(err)
		return errors.New(err.Error())
	}
	if !token.Valid {
		log.Println("Invalid tokens")
		return errors.New("invalid tokens")
	}
	return nil
}

func ExtractClaims(input token.SingleTokenInput) (*token.CustomClaims, error) {
	parsed, err := jwt.ParseWithClaims(input.ProvidedToken, &token.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		err := ValidateToken(input.ProvidedToken)
		if err != nil {
			return nil, err
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		log.Println(err)
		return nil, errors.New("invalid token")
	}
	claims := parsed.Claims.(*token.CustomClaims)
	return claims, nil
}

func GenerateToken(userID, logoutPin string, lifeteme int) (*token.AuthToken, error) {
	accessObj := token.AuthToken{
		UserID:    userID,
		ExpiresAt: time.Now().Local().Add(time.Second * time.Duration(lifeteme)),
		IssuedAt:  time.Now().Local(),
		LogoutPin: logoutPin,
	}
	newAccessToken, err := newToken(&accessObj)
	if err != nil {
		return nil, err
	}
	accessObj.Token = newAccessToken
	return &accessObj, nil
}
