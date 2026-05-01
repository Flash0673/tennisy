package register

import (
	"context"

	"golang.org/x/crypto/bcrypt"
	"tennisly.com/mvp/internal/modules/auth/action/register/dto"
	"tennisly.com/mvp/internal/modules/auth/action/register/service/register"
	"tennisly.com/mvp/internal/modules/auth/dal"
	dtoModule "tennisly.com/mvp/internal/modules/auth/dto"
	"tennisly.com/mvp/pkg/security"
	"tennisly.com/mvp/pkg/token"
)

type Action struct {
	register *register.Service
}

func New(aggregator *dal.Aggregator, tokenService *token.JWTService) *Action {
	return &Action{
		register: register.New(
			security.NewBcryptHasher(bcrypt.DefaultCost),
			tokenService,
			aggregator.User,
			aggregator.RefreshToken,
		),
	}
}

func (action *Action) Do(ctx context.Context, req dto.RegisterRequest) (*dtoModule.TokenResponse, error) {
	return action.register.Register(ctx, req)
}
