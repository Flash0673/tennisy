package auth

import (
	"context"
	"errors"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	authv1 "tennisly.com/mvp/pb/api/auth/v1"
)

const authorizationHeader = "Authorization"

var whiteList = map[string]bool{
	authv1.Auth_Register_FullMethodName: true,
	authv1.Auth_LogIn_FullMethodName:    true,
	authv1.Auth_Refresh_FullMethodName:  true,
}

type Parser interface {
	Parse(token string) (string, error)
}

func NewAuthInterceptor(parser Parser) func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

		// Проверяем URL запроса
		if whiteList[info.FullMethod] {
			return handler(ctx, req)
		}

		// Для остальных — логика проверки
		// 1. Extract metadata
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Errorf(codes.Unauthenticated, "metadata is missing")
		}

		// 2. Validate token
		tokens := md.Get(authorizationHeader)
		if len(tokens) == 0 {
			return nil, status.Errorf(codes.Unauthenticated, "invalid header")
		}

		uid, err := parseToken(tokens[0], parser)

		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "invalid token")
		}
		ctx = context.WithValue(ctx, "user_id", uid)
		return handler(ctx, req)
	}
}

func parseToken(token string, parser Parser) (string, error) {

	headerParts := strings.Split(token, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return "", errors.New("invalid auth header")
	}

	if len(headerParts[1]) == 0 {
		return "", errors.New("token is empty")
	}

	return parser.Parse(headerParts[1])
}
