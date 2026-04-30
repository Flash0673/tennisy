package dto

import "time"

type TokenResponse struct {
	AccessToken  string
	RefreshToken string
	ExpiresAt    time.Time
}
