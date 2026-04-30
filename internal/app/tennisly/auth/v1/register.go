package auth

import (
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"
	"tennisy.com/mvp/internal/modules/auth/action/register/dto"
	authv1 "tennisy.com/mvp/pb/api/auth/v1"
	common_token "tennisy.com/mvp/pb/api/common/token"
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
		TokenPair: &common_token.TokenPair{
			AccessToken:  tokenResp.AccessToken,
			RefreshToken: tokenResp.RefreshToken,
			ExpiresAt:    timestamppb.New(tokenResp.ExpiresAt),
		},
	}, nil
}
