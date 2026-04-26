package dao

import (
	"github.com/samber/lo"
	"tennisy.com/mvp/internal/modules/auth/domain/refresh_token"
	"tennisy.com/mvp/internal/xo"
)

type RefreshTokenDAO struct {
	xo.RefreshToken
}

func (r RefreshTokenDAO) ToDomain() refresh_token.RefreshToken {
	return refresh_token.RefreshToken{
		ID:        r.ID,
		Token:     r.Token,
		ExpiresAt: r.ExpiresAt,
		IsRevoked: r.IsRevoked,
		UserID:    r.UserID,
		CreatedAt: r.CreatedAt,
	}
}

type RefreshTokenDAOs []RefreshTokenDAO

func (rs RefreshTokenDAOs) ToDomain() []refresh_token.RefreshToken {
	return lo.Map(rs, func(item RefreshTokenDAO, _ int) refresh_token.RefreshToken {
		return item.ToDomain()
	})
}
