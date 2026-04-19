package internal

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"

	"tennisy.com/mvp/internal/modules/auth"
)

type modules struct {
	auth *auth.Module
}

type App struct {
	publicServer chi.Router
	adminServer  chi.Router
	grpcServer   *grpc.Server

	opts Options

	lis listeners

	modules modules

	// TODO Posgres conn
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

	// TODO здесь запускаем приложение

	a.runPublicHTTP()
	a.runGRPC()
}

// runGRPC in called in Run function
func (a *App) runGRPC() {
	if a.grpcServer != nil {
		go func() {
			if err := a.grpcServer.Serve(a.lis.grpc); err != nil {
				// TODO log + closeer
			}
		}()
		// TODO graceful shoutdown
	}
}

// runAdminHTTP is called in New function
func (a *App) runAdminHTTP() {
	adminServer := &http.Server{Handler: a.adminServer}
	go func() {
		if err := adminServer.Serve(a.lis.httpAdmin); err != nil && !errors.Is(err, http.ErrServerClosed) {
			// TODO log + closeer
		}
	}()
	// TODO graceful shoutdown
}

// runPublicHTTP is called in New function
func (a *App) runPublicHTTP() {
	// We don't need to start actual server
	//if !a.publicHTTPEnabled() {
	//	return
	//}

	//publicServer := &http.Server{Handler: a.publicServer}
	//go func() {
	//	if err := publicServer.Serve(a.lis.http); err != nil && !errors.Is(err, http.ErrServerClosed) {
	//		// TODO log + closeer
	//	}
	//}()
	//// TODO graceful shoutdown

	router := chi.NewRouter()
	router.Get("/", func(writer http.ResponseWriter, request *http.Request) {

		res, err := a.modules.auth.Actions.Register.Do(context.Background())

		_, err = writer.Write(res)
		if err != nil {
			// TODO log
		}
	})

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	srv := &http.Server{
		Addr:    "localhost:8080",
		Handler: router,
		//ReadTimeout:  cfg.HTTPServer.Timeout,
		//WriteTimeout: cfg.HTTPServer.Timeout,
		//IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {

		}
	}()

	<-done

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		// TODO log

		return
	}

}

func (a *App) publicHTTPEnabled() bool {
	return !a.opts.DisabledPublicHTTP
}
