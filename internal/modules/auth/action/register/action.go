package register

import (
	"context"
	"time"

	"golang.org/x/crypto/bcrypt"
	"tennisy.com/mvp/internal/modules/auth/action/register/dto"
	"tennisy.com/mvp/internal/modules/auth/action/register/service/register"
	"tennisy.com/mvp/internal/modules/auth/dal"
	"tennisy.com/mvp/pkg/security"
	"tennisy.com/mvp/pkg/token"
)

type Action struct {
	register *register.Service
}

func New(aggregator *dal.Aggregator) *Action {
	return &Action{
		register: register.New(
			security.NewBcryptHasher(bcrypt.DefaultCost),
			// TODO config
			token.NewJWTService("", 24*time.Hour),
			aggregator.User,
			aggregator.RefreshToken,
		),
	}
}

func (action *Action) Do(ctx context.Context, req dto.RegisterRequest) (*dto.TokenResponse, error) {
	return action.register.Register(ctx, req)
}
