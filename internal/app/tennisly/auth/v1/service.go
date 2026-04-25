package auth

import (
	"tennisy.com/mvp/internal/modules/auth"
	authv1 "tennisy.com/mvp/pb/auth/v1"
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
