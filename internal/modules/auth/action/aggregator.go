package action

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"tennisy.com/mvp/internal/modules/auth/action/register"
)

type Aggregator struct {
	// Сюда собираем экшены
	Register *register.Action
}

func New(pool *pgxpool.Pool) *Aggregator {
	return &Aggregator{
		Register: register.New(pool),
	}
}
