package env

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestNewEnv(t *testing.T) {
	t.Run("successful env loading", func(t *testing.T) {
		file, err := os.Create(".env")
		defer func() {
			file.Close()
			os.Remove(file.Name())
		}()
		require.NoError(t, err)
		err = NewEnv(".env")
		require.NoError(t, err)
	})
}
