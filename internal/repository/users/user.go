package users

import (
	"auth-service/internal/domain/user"
	internalErr "auth-service/internal/errors"
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"log"
)

const (
	userRoleID = 1
)

func (r repo) Create(ctx context.Context, user *user.User) error {
	var (
		userID      string
		insertQuery = `INSERT INTO users(email, password_hash, role_id) values ($1, $2, $3) returning id;`
	)
	if user.RoleID == 0 {
		user.RoleID = userRoleID
	}
	if err := r.pool.QueryRow(ctx, insertQuery, user.Email, user.PasswordHash, user.RoleID).Scan(&userID); err != nil {
		log.Println("users repo at Create():", err)
		return internalErr.ErrUserInsertion
	}

	log.Println("users created with uuid:", userID)
	return nil
}

func (r repo) GetByID(ctx context.Context, id string) (*user.User, error) {
	var (
		usr         user.User
		selectQuery = `SELECT id, email, password_hash, role_id, created_at FROM users WHERE id = $1;`
	)
	if err := r.pool.QueryRow(ctx, selectQuery, id).Scan(&usr.ID, &usr.Email, &usr.PasswordHash, &usr.RoleID, &usr.CreatedAt); err != nil {
		log.Println("users repo at GetByEmail():", err)
		return nil, internalErr.ErrGettingByID
	}
	log.Println("users selected from db with an ID:", usr.ID)
	return &usr, nil
}

func (r repo) GetByEmail(ctx context.Context, email string) (*user.User, error) {
	var (
		usr         user.User
		selectQuery = `SELECT id, email, password_hash, role_id, created_at FROM users WHERE email = $1;`
	)
	if err := r.pool.QueryRow(ctx, selectQuery, email).Scan(&usr.ID, &usr.Email, &usr.PasswordHash, &usr.RoleID, &usr.CreatedAt); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			log.Println("user not found:", err)
			return nil, nil
		}
		log.Println("users repo at GetByEmail():", err)
		return nil, internalErr.ErrGettingByEmail
	}
	log.Println("users selected from db with an email:", usr.Email)
	return &usr, nil
}

func (r repo) IsAdmin(ctx context.Context, userID string) error {
	selectQuery := `SELECT id FROM users WHERE id = $1 IF role='admin';`
	if err := r.pool.QueryRow(ctx, selectQuery, userID).Scan(&userID); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return internalErr.ErrNotAdmin
		}
		log.Println("users repo at IsAdmin():", err)
		return internalErr.ErrSelection
	}
	return nil
}

/*
func (r repo) SetVerified(ctx context.Context, userID string) error {
	updateQuery := `UPDATE users SET verified = TRUE WHERE id = $1;`
	if _, err := r.pool.Exec(ctx, updateQuery, userID); err != nil {
		log.Println("users repo at SetVerified():", err)
		return err
	}
	log.Println("user verified field set TRUE")
	return nil
}

*/
