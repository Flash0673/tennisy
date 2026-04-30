package auth

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func UserContextInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		if uids := md.Get("user_id"); len(uids) > 0 {
			// Снова кладем в контекст, но уже на стороне gRPC
			ctx = context.WithValue(ctx, "user_id", uids[0])
		}
	}
	return handler(ctx, req)
}
