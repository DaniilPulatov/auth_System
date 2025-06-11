package users

import (
	"auth-service/internal/domain/user"
	"auth-service/pkg/postgresql"
	"context"
)

type UsersRepository interface {
	Create(context.Context, *user.User) error
	GetByID(context.Context, string) (*user.User, error)
	GetByEmail(context.Context, string) (*user.User, error)
}

type repo struct {
	pool postgresql.Pool
}

func NewUserRepo(pool postgresql.Pool) UsersRepository {
	return repo{
		pool: pool,
	}
}
