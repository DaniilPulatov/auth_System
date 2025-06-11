package users

import (
	"auth-service/internal/domain/user"
	"auth-service/pkg/postgresql"
	"context"
)

type UserRepository interface {
	Create(context.Context, *user.User) error
	GetByID(context.Context, string) (*user.User, error)
	GetByEmail(context.Context, string) (*user.User, error)
	//SetVerified(context.Context, string) error
}

type repo struct {
	pool postgresql.Pool
}

func NewUserRepo(pool postgresql.Pool) UserRepository {
	return repo{
		pool: pool,
	}
}
