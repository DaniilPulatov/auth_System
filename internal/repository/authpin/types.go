package authpin

import (
	"auth-service/internal/domain/pin"
	"context"
	"github.com/redis/go-redis/v9"
)

type PinRepository interface {
	Create(context.Context, *pin.Pin) error
	Get(context.Context, string) (string, error)
}

type redisRepo struct {
	client *redis.Client
}

func NewPinRepository(client *redis.Client) PinRepository { return redisRepo{client: client} }
