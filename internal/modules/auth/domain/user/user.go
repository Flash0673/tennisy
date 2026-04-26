package user

import (
	"time"

	"github.com/google/uuid"
)

// User пользователь
type User struct {
	ID           uuid.UUID  // Уникальный идентификатор ползователя
	Email        string     // email ползователя
	Username     string     // username ползователя
	PasswordHash string     // Зашифрованный пароль пользователя
	FullName     *string    // Полное имя ползователя
	CreatedAt    time.Time  // Когда создан
	LastLogin    *time.Time // Последний вход
	IsActive     bool       // Активный пользователь
}
