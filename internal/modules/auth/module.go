package auth

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"tennisy.com/mvp/internal/modules/auth/action"
)

type Module struct {
	Actions *action.Aggregator
}

func New(pool *pgxpool.Pool) *Module {
	return &Module{
		Actions: action.New(pool),
	}
}
