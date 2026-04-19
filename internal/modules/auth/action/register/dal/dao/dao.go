package dao

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"tennisy.com/mvp/internal/modules/auth/action/register/dto"
)

type UserRow struct {
	ID        uuid.UUID      `db:"id"`
	Email     string         `db:"email"`
	Username  string         `db:"username"`
	FullName  sql.NullString `db:"full_name"`
	CreatedAt time.Time      `db:"created_at"`
	LastLogin sql.NullTime   `db:"last_login"`
	IsActive  bool           `db:"is_active"`
}

func (u UserRow) ToDto() dto.UserRow {
	return dto.UserRow{
		ID:        u.ID.String(),
		Email:     u.Email,
		Username:  u.Username,
		FullName:  u.FullName.String,
		CreatedAt: u.CreatedAt,
		LastLogin: u.LastLogin.Time,
		IsActive:  u.IsActive,
	}
}
