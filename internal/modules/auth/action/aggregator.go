package action

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"tennisly.com/mvp/internal/modules/auth/action/login"
	"tennisly.com/mvp/internal/modules/auth/action/register"
	"tennisly.com/mvp/internal/modules/auth/dal"
	"tennisly.com/mvp/pkg/token"
)

type Aggregator struct {
	// Сюда собираем экшены
	Register *register.Action
	Login    *login.Action
}

func New(pool *pgxpool.Pool, dataAccessLayer *dal.Aggregator, tokenService *token.JWTService) *Aggregator {
	_ = pool
	return &Aggregator{
		Register: register.New(dataAccessLayer, tokenService),
		Login:    login.New(dataAccessLayer, tokenService),
	}
}
