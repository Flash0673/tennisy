package refresh_token

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"tennisy.com/mvp/internal/modules/auth/domain/refresh_token"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRefreshTokenRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, token *refresh_token.RefreshToken) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := r.db.Exec(ctx,
		`INSERT INTO refresh_tokens (id, token, expires_at, is_revoked, user_id, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)`,
		token.ID, token.Token, token.ExpiresAt, token.IsRevoked, token.UserID, token.CreatedAt,
	)
	return err
}
