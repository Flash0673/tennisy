package auth

import (
	"context"

	"tennisly.com/mvp/internal/modules/auth/action/login/dto"
	authv1 "tennisly.com/mvp/pb/api/auth/v1"
)

func (i *Implementation) LogIn(context context.Context, req *authv1.LogInRequest) (*authv1.LogInResponse, error) {
	tokenResp, err := i.auth.Actions.Login.Do(context, dto.LoginRequest{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	})
	if err != nil {
		return nil, err
	}
	return &authv1.LogInResponse{
		TokenPair: TokenToPb(tokenResp),
	}, nil
}
