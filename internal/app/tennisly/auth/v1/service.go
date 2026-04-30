package auth

import (
	"tennisly.com/mvp/internal/modules/auth"
	authv1 "tennisly.com/mvp/pb/api/auth/v1"
)

// Implementation is a Route implementation
type Implementation struct {
	authv1.UnimplementedAuthServer
	auth *auth.Module
}

// NewAuth return new instance of Implementation.
func NewAuth(auth *auth.Module) *Implementation {
	return &Implementation{
		auth: auth,
	}
}
