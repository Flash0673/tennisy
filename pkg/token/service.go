package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"tennisly.com/mvp/pkg/token/dto"
)

type JWTService struct {
	secret         []byte
	accessTokenTTL time.Duration
}

type AccessClaims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func NewJWTService(secret string, accessTTL time.Duration) *JWTService {
	return &JWTService{
		secret:         []byte(secret),
		accessTokenTTL: accessTTL,
	}
}

func (s *JWTService) Generate(userID, email string) (*dto.TokenPair, error) {
	now := time.Now()

	claims := AccessClaims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(s.accessTokenTTL)),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(s.secret)
	if err != nil {
		return nil, err
	}

	return &dto.TokenPair{
		AccessToken:  accessToken,
		RefreshToken: uuid.NewString(),
		ExpiresAt:    claims.ExpiresAt.Time,
	}, nil
}

func (s *JWTService) Parse(token string) (string, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return s.secret, nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok || !t.Valid {
		return "", fmt.Errorf("invalid token")
	}
	return claims["user_id"].(string), nil
}
