package authHandler

import (
	"auth-service/internal/usecase/auth"
	"github.com/gin-gonic/gin"
)

type AuthHandler interface {
	Register(*gin.Context)
	Login(*gin.Context)
	Logout(*gin.Context)
	Refresh(*gin.Context)
}
type handler struct {
	authService auth.AuthenticationService
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func NewAuthHandler(authService auth.AuthenticationService) AuthHandler {
	return handler{authService: authService}
}
