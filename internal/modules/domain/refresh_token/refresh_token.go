package refresh_token

import (
	"time"

	"github.com/google/uuid"
)

// RefreshToken токен
type RefreshToken struct {
	ID        uuid.UUID // Уникальный идентификатор ползователя
	Token     string    // Значение токена
	ExpiresAt time.Time // Когда истекает токен
	IsRevoked bool      // Является ли отозванным
	UserID    uuid.UUID // Идентификатор ползователя
	CreatedAt time.Time // Когда создан
}
