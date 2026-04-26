package user

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// User пользователь
type User struct {
	ID           uuid.UUID      // Уникальный идентификатор ползователя
	Email        string         // email ползователя
	Username     string         // username ползователя
	PasswordHash string         // Зашифрованный пароль пользователя
	FullName     sql.NullString // Полное имя ползователя
	CreatedAt    time.Time      // Когда создан
	LastLogin    sql.NullTime   // Последний вход
	IsActive     bool           // Активный пользователь
}
