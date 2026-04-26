package dto

import "time"

type RegisterRequest struct {
	Username string
	FullName string
	Email    string
	Password string
}

type TokenResponse struct {
	AccessToken  string
	RefreshToken string
	ExpiresAt    time.Time
}
