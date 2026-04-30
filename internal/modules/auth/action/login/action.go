package login

import (
	"context"
	"time"

	"golang.org/x/crypto/bcrypt"
	"tennisly.com/mvp/internal/modules/auth/action/login/dto"
	"tennisly.com/mvp/internal/modules/auth/action/login/service/login"
	"tennisly.com/mvp/internal/modules/auth/dal"
	dtoModule "tennisly.com/mvp/internal/modules/auth/dto"
	"tennisly.com/mvp/pkg/security"
	"tennisly.com/mvp/pkg/token"
)

type Action struct {
	login *login.Service
}

func New(aggregator *dal.Aggregator) *Action {
	return &Action{
		login: login.New(
			security.NewBcryptHasher(bcrypt.DefaultCost),
			// TODO config
			token.NewJWTService("", 24*time.Hour),
			aggregator.User,
			aggregator.RefreshToken,
		),
	}
}

func (action *Action) Do(ctx context.Context, req dto.LoginRequest) (*dtoModule.TokenResponse, error) {
	return action.login.Login(ctx, req)
}
