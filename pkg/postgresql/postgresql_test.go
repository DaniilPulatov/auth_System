package postgresql

import (
	"github.com/lpernett/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"log"
	"os"
	"testing"
)

func TestNewPostgresDB(t *testing.T) {
	t.Run("invalid dsn", func(t *testing.T) {
		_, err := NewPostgresDB("invalid dsn")
		assert.Error(t, err)
	})

	t.Run("valid dsn", func(t *testing.T) {
		err := godotenv.Load("../../.env")
		require.NoError(t, err)
		_, err = NewPostgresDB(os.Getenv("DATABASE_URL"))
		if err != nil {
			log.Println(err)
		}
		assert.NoError(t, err)
	})
}
