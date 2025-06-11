package tokens

import (
	"auth-service/internal/domain/token"
	"auth-service/pkg/postgresql"
	"context"
)

type AuthenticationRepo interface {
	Create(context.Context, *token.AuthToken) error
	Get(context.Context, string) (*token.AuthToken, error)
	GetByUserID(context.Context, string) (*token.AuthToken, error)
	Update(context.Context, *token.AuthToken) error
	Delete(context.Context, *token.AuthToken) error
}

type postgresRepo struct {
	pool postgresql.Pool
}

func NewPostgresAuthRepo(pool postgresql.Pool) AuthenticationRepo {
	return postgresRepo{pool: pool}
}
