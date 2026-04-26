package dao

import (
	"github.com/samber/lo"
	"tennisy.com/mvp/internal/modules/auth/domain/user"
	"tennisy.com/mvp/internal/xo"
)

type UserDAO struct {
	xo.User
}

func (u UserDAO) ToDomain() user.User {
	return user.User{
		ID:           u.ID,
		Email:        u.Email,
		Username:     u.Username,
		PasswordHash: u.PasswordHash,
		FullName:     u.FullName,
		CreatedAt:    u.CreatedAt,
		LastLogin:    u.LastLogin,
		IsActive:     u.IsActive,
	}
}

type UserDAOs []UserDAO

func (us UserDAOs) ToDomain() []user.User {
	return lo.Map(us, func(item UserDAO, _ int) user.User {
		return item.ToDomain()
	})
}
