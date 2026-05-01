package internal

import (
	"context"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
	"tennisly.com/mvp/internal/modules/auth"
	"tennisly.com/mvp/pkg/token"
)

const (
	grpcAddress = "localhost:7002"
	httpAddress = "localhost:8080"
)

type modules struct {
	auth *auth.Module
}

type App struct {
	httpServer *runtime.ServeMux
	grpcServer *grpc.Server

	modules modules

	postgresConnPool *pgxpool.Pool

	tokenService *token.JWTService

	// Всякие коннекшены
}

// New создает новое приложение
func New(ctx context.Context) *App {
	a := &App{}

	a.initPostgres(ctx).
		initServices(ctx).
		initModules(ctx).
		initHttpServer(ctx).
		initGrpcServer(ctx)

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

	list, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		return err
	}

	log.Printf("gRPC server listening at %v\n", grpcAddress)

	return a.grpcServer.Serve(list)
}

// runPublicHTTP is called in New function
func (a *App) runPublicHTTP(_ context.Context) error {

	log.Printf("http server listening at %v\n", httpAddress)

	return http.ListenAndServe(httpAddress, a.httpServer)
}
