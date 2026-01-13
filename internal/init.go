package internal

import (
	"context"

	"github.com/go-chi/chi/v5"
	"tennisy.com/mvp/internal/modules/auth"
)

func (a *App) initModules(_ context.Context) *App {
	a.modules = modules{
		auth: auth.New(),
	}

	return a
}

func (a *App) initConfig(_ context.Context) *App {
	return a
}

func (a *App) initLogger(_ context.Context) *App {
	return a
}

func (a *App) initPostgres(_ context.Context) *App {
	return a
}

func (a *App) initHttp(_ context.Context) *App {
	a.publicServer = chi.NewRouter()
	return a
}
