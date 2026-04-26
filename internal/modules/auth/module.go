package auth

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"tennisy.com/mvp/internal/modules/auth/action"
	"tennisy.com/mvp/internal/modules/auth/dal"
)

type Module struct {
	Actions *action.Aggregator
}

func New(pool *pgxpool.Pool) *Module {
	dataAccessLayer := dal.New(pool)
	return &Module{
		Actions: action.New(pool, dataAccessLayer),
	}
}
