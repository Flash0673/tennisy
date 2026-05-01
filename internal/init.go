package internal

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	authController "tennisly.com/mvp/internal/app/tennisly/auth/v1"
	authv1 "tennisly.com/mvp/pb/api/auth/v1"
	authInterceptor "tennisly.com/mvp/pkg/middleware/grpc/auth"
	authMiddleware "tennisly.com/mvp/pkg/middleware/http/auth"
	metadataInterceptor "tennisly.com/mvp/pkg/middleware/metadata"
	"tennisly.com/mvp/pkg/token"

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

func (a *App) initHttpServer(ctx context.Context) *App {
	a.httpServer = runtime.NewServeMux(
		runtime.WithMiddlewares(
			authMiddleware.NewAuthMiddleware(a.tokenService),
		),
		runtime.WithMetadata(metadataInterceptor.UserMetadata),
	)

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	// Здесь регаем http ендпоинты
	err := authv1.RegisterAuthHandlerFromEndpoint(ctx, a.httpServer, grpcAddress, opts)
	if err != nil {
		log.Fatal("register grpc gateway: %w", err)
	}

	return a
}

func (a *App) initGrpcServer(ctx context.Context) *App {
	a.grpcServer = grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
		grpc.ChainUnaryInterceptor(
			//authInterceptor.NewAuthInterceptor(a.tokenService),
			authInterceptor.UserContextInterceptor,
		),
	)

	reflection.Register(a.grpcServer)

	// Здесь регистрируем наши grpc контроллеры
	authv1.RegisterAuthServer(a.grpcServer, authController.NewAuth(a.modules.auth))

	return a
}

func (a *App) initServices(_ context.Context) *App {
	a.tokenService = token.NewJWTService("", 24*time.Hour)
	return a
}
