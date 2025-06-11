package auth

import (
	"auth-service/internal/domain/pin"
	"auth-service/internal/domain/token"
	"auth-service/internal/domain/user"
	internalErr "auth-service/internal/errors"
	"auth-service/internal/hashing"
	"auth-service/internal/tokenOps"
	"auth-service/pkg/utils"
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

const (
	REFRESH_LIFETIME = 7 * 24 * 3600
	ACCESS_LIFETIME  = 15 * 60
)

func (s service) Register(ctx context.Context, input *user.Input) error {
	if !utils.IsValidEmail(input.Email) {
		return errors.New("invalid email")
	}
	oldUser, err := s.userRepo.GetByEmail(ctx, input.Email)
	if err != nil {

		return err
	}
	if oldUser != nil {
		return internalErr.ErrUserAlreadyExists
	}
	pswdHash, err := hashing.HashPassword(input.Password)
	if err != nil {
		return err
	}
	user := &user.User{
		Email:        input.Email,
		PasswordHash: pswdHash,
		ID:           uuid.New().String(),
	}
	err = s.userRepo.Create(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (s service) Login(ctx context.Context, input *user.Input) (map[string]string, error) {
	if !utils.IsValidEmail(input.Email) {
		return nil, errors.New("invalid email")
	}
	tokens := make(map[string]string)

	existingUser, err := s.userRepo.GetByEmail(ctx, input.Email) // check if user with provided email exists
	if err != nil {
		return tokens, err
	}
	if existingUser == nil {
		return tokens, errors.New("user not found")
	}
	if !hashing.CheckPasswordHash(input.Password, existingUser.PasswordHash) { // check if password is correct
		return tokens, errors.New("invalid password or email")
	}
	logoutPin := uuid.New().String()

	accessToken, err := tokenOps.GenerateToken(existingUser.ID, logoutPin, ACCESS_LIFETIME)
	if err != nil {
		return tokens, err
	}

	refreshToken, err := tokenOps.GenerateToken(existingUser.ID, logoutPin, REFRESH_LIFETIME)
	if err != nil {
		return tokens, err
	}

	oldRefresh, err := s.tokenProvider.Get(ctx, refreshToken.Token) // look up for old refresh token
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return tokens, err
	}
	if oldRefresh != nil {
		if err := s.tokenRepo.Delete(ctx, oldRefresh); err != nil { // if token already in database  then delete
			return tokens, err
		}
	}

	if err := s.tokenRepo.Create(ctx, refreshToken); err != nil { // create new refresh token
		return tokens, err
	}
	err = s.pinRepo.Create(ctx, &pin.Pin{UserID: existingUser.ID, LogoutPin: logoutPin})
	if err != nil {
		return tokens, err
	}

	tokens["access"] = accessToken.Token
	tokens["refresh"] = refreshToken.Token
	return tokens, nil
}

func (s service) Refresh(ctx context.Context, input token.SingleTokenInput) (map[string]string, error) {
	var (
		tokens = make(map[string]string)
		claims = &token.CustomClaims{}
	)
	claims, err := tokenOps.ExtractClaims(input)
	if err != nil {
		return tokens, err
	}
	if claims == nil {
		return tokens, errors.New("invalid token")
	}

	accessToken, err := tokenOps.GenerateToken(claims.ID, claims.LogoutPin, ACCESS_LIFETIME)
	if err != nil {
		return tokens, err
	}

	refreshToken, err := tokenOps.GenerateToken(claims.UserID, claims.LogoutPin, REFRESH_LIFETIME)
	if err != nil {
		return tokens, err
	}

	if err := s.tokenRepo.Delete(ctx, &token.AuthToken{UserID: claims.UserID, Token: input.ProvidedToken}); err != nil {
		return tokens, err
	}
	if err := s.tokenRepo.Create(ctx, refreshToken); err != nil {
		return tokens, err
	}
	tokens["access"] = accessToken.Token
	tokens["refresh"] = refreshToken.Token
	return tokens, nil
}

func (s service) Logout(ctx context.Context, userID string) error {
	newLogoutPin := uuid.New().String()
	if err := s.pinRepo.Create(ctx, &pin.Pin{
		UserID: userID, LogoutPin: newLogoutPin,
	}); err != nil {
		return err
	}
	return nil
}
