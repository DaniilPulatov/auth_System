package env

import (
	pkgErrors "auth-service/pkg/errors"
	"github.com/lpernett/godotenv"
	"log"
)

func NewEnv(filepath string) error {
	if err := godotenv.Load(filepath); err != nil {
		log.Println("Error loading .env file:", err)
		return pkgErrors.ErrNewEnv
	}
	log.Println("environmental variables are loaded successfully")
	return nil
}
