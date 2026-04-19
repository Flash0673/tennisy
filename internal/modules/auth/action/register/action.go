package register

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"tennisy.com/mvp/internal/modules/auth/action/register/dal"
)

type Action struct {
	repo *dal.Dal
}

func New(pool *pgxpool.Pool) *Action {
	return &Action{
		repo: dal.NewDal(pool),
	}
}

func (action *Action) Do(ctx context.Context) ([]byte, error) {
	usr, err := action.repo.Register(ctx)
	if err != nil {
		return nil, err
	}
	return []byte(usr.String()), nil
}
