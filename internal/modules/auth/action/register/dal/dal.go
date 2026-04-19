package dal

import (
	"context"
	"log"
	"time"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jaswdr/faker/v2"
	"tennisy.com/mvp/internal/modules/auth/action/register/dal/dao"
	"tennisy.com/mvp/internal/modules/auth/action/register/dto"
)

type Dal struct {
	pool *pgxpool.Pool
}

func NewDal(pool *pgxpool.Pool) *Dal {
	return &Dal{
		pool: pool,
	}
}

func (d *Dal) Register(ctx context.Context) (dto.UserRow, error) {
	sql := `
		INSERT INTO users (id, email, username, created_at, is_active)
		VALUES ($1, $2, $3, $4, $5)
		returning *;
`
	args := []any{
		uuid.New(),
		faker.New().Internet().Email(),
		faker.New().Person().FirstName(),
		time.Now(),
		true,
	}

	var userRow dao.UserRow
	err := pgxscan.Get(ctx, d.pool, &userRow, sql, args...)
	if err != nil {
		log.Println(err)
		return dto.UserRow{}, err
	}

	return userRow.ToDto(), nil
}
