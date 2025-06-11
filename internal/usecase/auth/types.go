package auth

import (
	"auth-service/internal/domain/token"
	"auth-service/internal/domain/user"
	"auth-service/internal/repository/authpin"
	"auth-service/internal/repository/tokens"
	"auth-service/internal/repository/users"
	"context"
)

type AuthenticationService interface {
	Register(context.Context, *user.Input) error
	Login(context.Context, *user.Input) (map[string]string, error) // first - access, second - refresh
	Logout(context.Context, string) error
	Refresh(context.Context, token.SingleTokenInput) (map[string]string, error)
}

type TokenProvider interface {
	Get(context.Context, string) (*token.AuthToken, error)
	// GetByUserID(context.Context, string) (*token.AuthToken, error)
	// GetByUserID(context.Context, string) (*token.AuthToken, error)
}

type service struct {
	userRepo      users.UsersRepository
	tokenRepo     tokens.AuthenticationRepo
	pinRepo       authpin.PinRepository
	tokenProvider TokenProvider
	//adminChecker AdminChecker
}

func NewAuthService(usrRepo users.UsersRepository,
	tokenRepo tokens.AuthenticationRepo, pinRepo authpin.PinRepository,
	tokenProvider TokenProvider) AuthenticationService {
	return service{
		userRepo:      usrRepo,
		tokenRepo:     tokenRepo,
		pinRepo:       pinRepo,
		tokenProvider: tokenProvider,
	}
}
