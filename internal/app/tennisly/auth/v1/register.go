package auth

import (
	"context"

	"tennisly.com/mvp/internal/modules/auth/action/register/dto"
	authv1 "tennisly.com/mvp/pb/api/auth/v1"
)

func (i *Implementation) Register(ctx context.Context, req *authv1.RegisterRequest) (resp *authv1.RegisterResponse, err error) {
	tokenResp, err := i.auth.Actions.Register.Do(ctx, dto.RegisterRequest{
		Username: req.GetUserName(),
		FullName: req.GetFullName(),
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	})
	if err != nil {
		return nil, err
	}
	return &authv1.RegisterResponse{
		TokenPair: TokenToPb(tokenResp),
	}, nil
}
