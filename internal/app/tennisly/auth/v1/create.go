package auth

import (
	"context"

	authv1 "tennisy.com/mvp/pb/auth/v1"
)

func (i *Implementation) Create(ctx context.Context, req *authv1.CreateRequest) (resp *authv1.CreateResponse, err error) {
	user, err := i.auth.Actions.Register.Do(ctx)
	if err != nil {
		return nil, err
	}
	return &authv1.CreateResponse{
		UserId: user.ID,
	}, nil
}
