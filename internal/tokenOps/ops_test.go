package tokenOps

import (
	"auth-service/internal/domain/token"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"log"
	"os"
	"testing"
	"time"
)

const secret = "secret"

func TestNewToken(t *testing.T) {
	os.Setenv("JWT_SECRET", secret)
	defer os.Unsetenv("JWT_SECRET")
	userId := uuid.New().String()
	t.Run("Successful generation", func(t *testing.T) {
		obj := &token.AuthToken{
			UserID:    userId,
			ExpiresAt: time.Now().Local().Add(time.Hour),
			IssuedAt:  time.Now().Local(),
			LogoutPin: "random authpin authpin",
		}
		token, err := newToken(obj)
		require.NoError(t, err)
		log.Println("Generated tokens:", token)
		require.NotNil(t, token)
	})
}

func TestValidateToken(t *testing.T) {
	os.Setenv("JWT_SECRET", secret)
	defer os.Unsetenv("JWT_SECRET")
	userId := uuid.New().String()
	t.Run("Successful hashing", func(t *testing.T) {
		obj := &token.AuthToken{
			UserID:    userId,
			ExpiresAt: time.Now().Local().Add(time.Hour),
			IssuedAt:  time.Now().Local(),
			LogoutPin: "random authpin authpin",
		}
		token, err := newToken(obj)
		require.NoError(t, err)
		log.Println("Generated tokens:", token)
		require.NotNil(t, token)

		err = ValidateToken(token)
		require.NoError(t, err)
	})
	t.Run("tokens expired", func(t *testing.T) {
		obj := &token.AuthToken{
			UserID:    userId,
			ExpiresAt: time.Now().Local().Add(-time.Hour),
			IssuedAt:  time.Now().Local(),
			LogoutPin: "random authpin authpin",
		}
		token, err := newToken(obj)
		require.NoError(t, err)
		log.Println("Generated tokens:", token)
		require.NotNil(t, token)
		time.Sleep(1 * time.Second)
		err = ValidateToken(token)
		require.NotNil(t, err)
		require.Contains(t, err.Error(), "tokens is expired")
	})
}

func TestExtractClaims(t *testing.T) {
	os.Setenv("JWT_SECRET", secret)
	defer os.Unsetenv("JWT_SECRET")
	userId := uuid.New().String()
	logoutPin := "random authpin authpin"
	t.Run("Successful extract claims", func(t *testing.T) {
		obj := &token.AuthToken{
			UserID:    userId,
			ExpiresAt: time.Now().Local().Add(time.Hour),
			IssuedAt:  time.Now().Local(),
			LogoutPin: logoutPin,
		}
		newToken, err := newToken(obj)
		require.NoError(t, err)
		log.Println("Generated tokens:", newToken)
		require.NotNil(t, newToken)

		claims, err := ExtractClaims(token.SingleTokenInput{ProvidedToken: newToken})
		require.NoError(t, err)
		require.Equal(t, claims.UserID, userId)
		require.Equal(t, claims.LogoutPin, logoutPin)
	})
}

func TestGenerateAccessToken(t *testing.T) {
	os.Setenv("JWT_SECRET", secret)
	defer os.Unsetenv("JWT_SECRET")
	userId := uuid.New().String()
	logoutPin := "random authpin authpin"
	t.Run("Successful generate access token", func(t *testing.T) {
		accessToken, err := GenerateToken(userId, logoutPin, 180*60)
		require.NoError(t, err)
		log.Println("Generated access tokens:", accessToken.Token)
		require.NotNil(t, accessToken)
		require.Equal(t, accessToken.UserID, userId)
		require.Equal(t, accessToken.LogoutPin, logoutPin)

	})
}
