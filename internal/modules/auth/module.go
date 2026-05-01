package auth

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"tennisly.com/mvp/internal/modules/auth/action"
	"tennisly.com/mvp/internal/modules/auth/dal"
	"tennisly.com/mvp/pkg/token"
)

type Module struct {
	Actions *action.Aggregator
}

func New(pool *pgxpool.Pool, tokenService *token.JWTService) *Module {
	dataAccessLayer := dal.New(pool)
	return &Module{
		Actions: action.New(pool, dataAccessLayer, tokenService),
	}
}
