package user

import (
	"context"
	"errors"
	"time"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"tennisly.com/mvp/internal/modules/auth/domain/user"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, user *user.User) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	sql := `INSERT INTO users (id, email, username, password_hash, full_name, created_at, last_login, is_active) 
		VALUES
		($1, $2, $3, $4, $5, $6, $7, $8)`

	args := []any{
		user.ID,
		user.Email,
		user.Username,
		user.PasswordHash,
		user.FullName,
		user.CreatedAt,
		user.LastLogin,
		user.IsActive,
	}

	_, err := r.db.Exec(ctx, sql, args...)

	return err
}

func (r *Repository) GetByEmail(ctx context.Context, email string) (*user.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	sql := `SELECT * FROM users WHERE email = $1`

	usr := &user.User{}
	err := pgxscan.Get(ctx, r.db, usr, sql, email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			// TODO сделать обертку над ошибками
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return usr, nil
}
