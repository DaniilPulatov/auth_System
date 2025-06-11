package di

import (
	"auth-service/internal/migrations"
	"auth-service/internal/repository/authpin"
	"auth-service/internal/repository/tokens"
	"auth-service/internal/repository/users"
	"auth-service/internal/rest"
	"auth-service/internal/rest/handlers/authHandler"
	"auth-service/internal/rest/middleware"
	"auth-service/internal/usecase/auth"
	"auth-service/pkg/postgresql"
	"auth-service/pkg/redisDB"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

func NewMux() *gin.Engine {
	return gin.New()
}

func NewHTTPServer(lc fx.Lifecycle, server *rest.Server) *http.Server {
	srv := &http.Server{
		Addr:              net.JoinHostPort(os.Getenv("HOST"), os.Getenv("PORT")),
		Handler:           server,
		ReadHeaderTimeout: 10 * time.Second,
	}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
					log.Printf("Error starting server: %s\n", err)
				}
			}()
			server.Run()
			log.Println("Server started")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("Shutting down server")
			return srv.Close()
		},
	})
	return srv
}

func PostgresProvider(lc fx.Lifecycle) (postgresql.Pool, error) {
	pool, err := postgresql.NewPostgresDB(os.Getenv("POSTGRES_URL"))
	if err != nil {
		log.Println("Error connecting to postgres", err)
		return nil, err
	}
	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			pool.Close()
			return nil
		},
	})
	return pool, nil
}

func RedisProvider() (*redis.Client, error) {
	client, err := redisDB.NewRedisDB()
	if err != nil {
		return nil, err
	}
	return client, nil
}

func NewModule() fx.Option {
	return fx.Module("auth service",
		fx.Provide(
			NewMux,
			PostgresProvider,
			RedisProvider,
			middleware.NewMiddleware,
			fx.Annotate(
				authHandler.NewAuthHandler,
				fx.As(new(authHandler.AuthHandler)),
			),

			fx.Annotate(
				users.NewUserRepo,
				fx.As(new(users.UsersRepository)),
			),
			fx.Annotate(
				tokens.NewPostgresAuthRepo,
				fx.As(new(tokens.AuthenticationRepo)),
				fx.As(new(auth.TokenProvider)),
			),
			fx.Annotate(
				authpin.NewPinRepository,
				fx.As(new(authpin.PinRepository)),
			),
			fx.Annotate(
				auth.NewAuthService,
				fx.As(new(auth.AuthenticationService)),
			),
			rest.NewServer,
			http.NewServeMux,
			NewHTTPServer,
		),
		fx.Invoke(
			func(lc fx.Lifecycle) {
				lc.Append(fx.Hook{
					OnStart: func(ctx context.Context) error {
						return migrations.ApplyMigrations(os.Getenv("MIGRATIONS_DIR"), os.Getenv("POSTGRES_URL"))
					},
				})
			},
			func(*http.Server) {},
		),
	)
}
