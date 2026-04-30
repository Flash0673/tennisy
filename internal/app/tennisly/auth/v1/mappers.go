package auth

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"tennisly.com/mvp/internal/modules/auth/dto"
	common_token "tennisly.com/mvp/pb/api/common/token"
)

func TokenToPb(token *dto.TokenResponse) *common_token.TokenPair {
	if token == nil {
		return nil
	}
	return &common_token.TokenPair{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		ExpiresAt:    timestamppb.New(token.ExpiresAt),
	}
}
