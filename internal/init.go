package internal

import (
	"context"
	"fmt"
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"tennisly.com/mvp/internal/modules/auth"
)

func (a *App) initModules(_ context.Context) *App {
	a.modules = modules{
		auth: auth.New(a.postgresConnPool),
	}

	return a
}

func (a *App) initConfig(_ context.Context) *App {
	return a
}

func (a *App) initLogger(_ context.Context) *App {
	return a
}

func (a *App) initPostgres(ctx context.Context) *App {
	var (
		// TODO: вынести в конфиг
		//postgresUser     = os.Getenv("POSTGRES_USER")
		//postgresPassword = os.Getenv("POSTGRES_PASSWORD")
		//postgresHost     = os.Getenv("POSTGRES_HOST")
		//postgresPort     = os.Getenv("POSTGRES_PORT")
		//postgresDB       = os.Getenv("POSTGRES_DB")
		postgresUser     = "postgres"
		postgresPassword = "password"
		postgresHost     = "localhost"
		postgresPort     = "5432"
		postgresDB       = "tennisly"
	)
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		postgresUser, postgresPassword, postgresHost, postgresPort, postgresDB,
	)
	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatal("parse config: %w", err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, cfg)

	a.postgresConnPool = pool

	// TODO: add closure
	//a.postgresConnPool.Close()

	if err != nil {
		log.Fatal("pgxpool connect: %w", err)
	}

	return a
}

func (a *App) initHttp(_ context.Context) *App {
	a.publicServer = chi.NewRouter()
	return a
}
