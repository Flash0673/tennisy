package auth

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	authv1 "tennisly.com/mvp/pb/api/auth/v1"
)

func (i *Implementation) Dummy(ctx context.Context, req *authv1.DummyMsg) (*authv1.DummyMsg, error) {

	return nil, status.Error(codes.Unimplemented, fmt.Sprintf("%v", ctx.Value("user_id")))
}
