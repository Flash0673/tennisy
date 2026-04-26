package action

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"tennisy.com/mvp/internal/modules/auth/action/register"
	"tennisy.com/mvp/internal/modules/auth/dal"
)

type Aggregator struct {
	// Сюда собираем экшены
	Register *register.Action
}

func New(pool *pgxpool.Pool, dataAccessLayer *dal.Aggregator) *Aggregator {
	_ = pool
	return &Aggregator{
		Register: register.New(dataAccessLayer),
	}
}
