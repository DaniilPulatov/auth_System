package tokens

import (
	"auth-service/internal/domain/token"
	internalErr "auth-service/internal/errors"
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"log"
)

func (r postgresRepo) Create(ctx context.Context, token *token.AuthToken) error {
	insertQuery := `insert into tokens(token, user_id, issued_at, expires_at, logout_pin) values($1, $2, $3, $4, $5);`
	if _, err := r.pool.Exec(ctx, insertQuery, token.Token, token.UserID, token.IssuedAt, token.ExpiresAt, token.LogoutPin); err != nil {
		log.Printf("failed to insert tokens: %v", err)
		return internalErr.ErrTokenInsertion
	}

	log.Println("tokens created")
	return nil
}

func (r postgresRepo) GetByUserID(ctx context.Context, id string) (*token.AuthToken, error) {
	var (
		selectQuery = `select token, user_id, expires_at, issued_at, logout_pin from tokens where user_id = $1;`
		rToken      token.AuthToken
	)
	if err := r.pool.QueryRow(ctx, selectQuery, id).Scan(&rToken.Token, &rToken.UserID, &rToken.ExpiresAt, &rToken.IssuedAt, &rToken.LogoutPin); err != nil {
		log.Printf("failed to get tokens: %v", err)
		return nil, internalErr.ErrGettingByToken
	}
	log.Println("tokens found")
	return &rToken, nil
}

func (r postgresRepo) Get(ctx context.Context, refreshToken string) (*token.AuthToken, error) {
	var (
		selectQuery = `select token, user_id, expires_at, issued_at, logout_pin from tokens where token = $1;`
		rToken      token.AuthToken
	)
	if err := r.pool.QueryRow(ctx, selectQuery, refreshToken).Scan(&rToken.Token, &rToken.UserID, &rToken.ExpiresAt, &rToken.IssuedAt, &rToken.LogoutPin); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, pgx.ErrNoRows
		}
		log.Printf("failed to get tokens: %v", err)
		return nil, internalErr.ErrGettingByToken
	}
	log.Println("tokens found")
	return &rToken, nil
}

func (r postgresRepo) Update(ctx context.Context, token *token.AuthToken) error {
	updateQuery := `update tokens set logout_pin = $1 where token = $2;`
	if _, err := r.pool.Exec(ctx, updateQuery, token.LogoutPin, token.Token); err != nil {
		log.Printf("failed to update tokens's authpin authpin: %v", err)
		return internalErr.ErrUpdatingLogoutPin
	}
	log.Println("tokens's authpin authpin updated")
	return nil
}

func (r postgresRepo) Delete(ctx context.Context, token *token.AuthToken) error {
	deleteQuery := `delete from tokens where user_id = $1;`
	if _, err := r.pool.Exec(ctx, deleteQuery, token.UserID); err != nil {
		log.Printf("failed to delete tokens: %v", err)
		return internalErr.ErrTokenDeletion
	}
	log.Println("tokens deleted")
	return nil
}
