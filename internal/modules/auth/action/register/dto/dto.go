package dto

import (
	"fmt"
	"time"
)

type UserRow struct {
	ID        string
	Email     string
	Username  string
	FullName  string
	CreatedAt time.Time
	LastLogin time.Time
	IsActive  bool
}

func (u UserRow) String() string {
	return fmt.Sprintf(`UserInfo: {
	%s,
	%s,
	%s,
	%s,
	%s,
	%s,
	%s,
}`, u.ID, u.Email, u.Username, u.FullName, u.CreatedAt, u.LastLogin, u.IsActive)
}
