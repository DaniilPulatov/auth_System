package main

import (
	"auth-service/internal/di"
	"auth-service/pkg/env"
	"go.uber.org/fx"
	"log"
)

func main() {
	if err := env.NewEnv(".env"); err != nil {
		log.Fatal(err)
	}
	fx.New(di.NewModule()).Run()
}
