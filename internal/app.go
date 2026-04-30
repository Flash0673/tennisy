package internal

import (
	"context"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	authController "tennisly.com/mvp/internal/app/tennisly/auth/v1"
	authv1 "tennisly.com/mvp/pb/api/auth/v1"

	"tennisly.com/mvp/internal/modules/auth"
)

const (
	grpcAddress = "localhost:7002"
	httpAddress = "localhost:8080"
)

type modules struct {
	auth *auth.Module
}

type App struct {
	publicServer chi.Router
	adminServer  chi.Router
	grpcServer   *grpc.Server

	modules modules

	postgresConnPool *pgxpool.Pool

	// Всякие коннекшены
}

// New создает новое приложение
func New(ctx context.Context) *App {
	a := &App{}

	a.initPostgres(ctx).
		initModules(ctx)

	return a
}

// Run переопределяет конфигурацию для запуска сервиса локально
func (a *App) Run(ctx context.Context) {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()

		if err := a.runGRPC(); err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		defer wg.Done()

		if err := a.runPublicHTTP(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	wg.Wait()
}

// runGRPC in called in Run function
func (a *App) runGRPC() error {
	grpcServer := grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
	)

	reflection.Register(grpcServer)

	// Здесь регистрируем наши grpc контроллеры
	authv1.RegisterAuthServer(grpcServer, authController.NewAuth(a.modules.auth))

	list, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		return err
	}

	log.Printf("gRPC server listening at %v\n", grpcAddress)

	return grpcServer.Serve(list)
}

// runPublicHTTP is called in New function
func (a *App) runPublicHTTP(ctx context.Context) error {
	mux := runtime.NewServeMux(
		runtime.WithMiddlewares(),
	)

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	// Здесь регаем http ендпоинты
	err := authv1.RegisterAuthHandlerFromEndpoint(ctx, mux, grpcAddress, opts)
	if err != nil {
		return err
	}

	log.Printf("http server listening at %v\n", httpAddress)

	return http.ListenAndServe(httpAddress, mux)
}
