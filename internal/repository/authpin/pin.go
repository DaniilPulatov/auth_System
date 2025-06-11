package authpin

import (
	"auth-service/internal/domain/pin"
	"context"
	"log"
)

func (r redisRepo) Create(ctx context.Context, pin *pin.Pin) error {
	err := r.client.Set(ctx, pin.UserID, pin.LogoutPin, 0).Err()
	if err != nil {
		log.Println("redisDB set error:", err)
		return err
	}
	log.Println("redisDB set pin:", pin.LogoutPin)
	return nil
}

func (r redisRepo) Get(ctx context.Context, userID string) (string, error) {
	logoutPin, err := r.client.Get(ctx, userID).Result()
	if err != nil {
		log.Println("redisDB get error:", err)
		return "", err
	}
	return logoutPin, nil
}
