package auth

import (
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"
	"tennisy.com/mvp/internal/modules/auth/action/register/dto"
	authv1 "tennisy.com/mvp/pb/auth/v1"
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
		AccessToken:  tokenResp.AccessToken,
		RefreshToken: tokenResp.RefreshToken,
		ExpiresAt:    timestamppb.New(tokenResp.ExpiresAt),
	}, nil
}
