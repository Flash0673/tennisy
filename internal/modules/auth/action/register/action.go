package register

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"tennisy.com/mvp/internal/modules/auth/action/register/dal"
	"tennisy.com/mvp/internal/modules/auth/action/register/dto"
)

type Action struct {
	repo *dal.Dal
}

func New(pool *pgxpool.Pool) *Action {
	return &Action{
		repo: dal.NewDal(pool),
	}
}

func (action *Action) Do(ctx context.Context) (dto.UserRow, error) {
	usr, err := action.repo.Register(ctx)
	if err != nil {
		return dto.UserRow{}, err
	}
	return usr, nil
}
