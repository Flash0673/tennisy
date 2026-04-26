package dal

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"tennisy.com/mvp/internal/modules/auth/dal/refresh_token"
	"tennisy.com/mvp/internal/modules/auth/dal/user"
)

type Aggregator struct {
	RefreshToken *refresh_token.Repository
	User         *user.Repository
}

func New(pool *pgxpool.Pool) *Aggregator {
	return &Aggregator{
		RefreshToken: refresh_token.NewRefreshTokenRepository(pool),
		User:         user.NewUserRepository(pool),
	}
}
