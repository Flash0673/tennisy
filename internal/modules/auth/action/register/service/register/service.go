package register

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"tennisy.com/mvp/internal/modules/auth/action/register/dto"
	"tennisy.com/mvp/internal/modules/auth/domain/refresh_token"
	"tennisy.com/mvp/internal/modules/auth/domain/user"
	dtoToken "tennisy.com/mvp/pkg/token/dto"
)

//go:generate mockgen -destination=mocks/mocks.go -package=mocks . PasswordHasher,TokenService,RefreshTokenRepository,UserRepository

type PasswordHasher interface {
	Hash(password string) (string, error)
}

type TokenService interface {
	Generate(userID, email string) (*dtoToken.TokenPair, error)
}

type RefreshTokenRepository interface {
	Create(ctx context.Context, token *refresh_token.RefreshToken) error
}

type UserRepository interface {
	Create(ctx context.Context, user *user.User) error
}

type Service struct {
	passwordHasher         PasswordHasher
	tokenService           TokenService
	userRepository         UserRepository
	refreshTokenRepository RefreshTokenRepository
}

func New(
	passwordHasher PasswordHasher,
	tokenService TokenService,
	userRepository UserRepository,
	refreshTokenService RefreshTokenRepository,
) *Service {
	return &Service{
		passwordHasher:         passwordHasher,
		tokenService:           tokenService,
		userRepository:         userRepository,
		refreshTokenRepository: refreshTokenService,
	}
}

// Register TODO разбить на 2 сервиса - отдельно юзер, отдельно токен и засунуть под транзакцию
func (s *Service) Register(ctx context.Context, req dto.RegisterRequest) (*dto.TokenResponse, error) {
	hash, err := s.passwordHasher.Hash(req.Password)
	if err != nil {
		return nil, err
	}

	usr := &user.User{
		ID:           uuid.New(),
		Email:        req.Email,
		Username:     req.Username,
		PasswordHash: hash,
		FullName:     lo.ToPtr(req.FullName),
		CreatedAt:    time.Now().UTC(),
		LastLogin:    lo.ToPtr(time.Now()),
		IsActive:     true,
	}

	if err := s.userRepository.Create(ctx, usr); err != nil {
		return nil, err
	}

	tokens, err := s.tokenService.Generate(
		usr.ID.String(),
		usr.Email,
	)
	if err != nil {
		return nil, err
	}

	rt := &refresh_token.RefreshToken{
		ID:        uuid.New(),
		UserID:    usr.ID,
		Token:     tokens.RefreshToken,
		ExpiresAt: tokens.ExpiresAt,
		CreatedAt: time.Now().UTC(),
		IsRevoked: false,
	}

	if err := s.refreshTokenRepository.Create(ctx, rt); err != nil {
		return nil, err
	}

	return &dto.TokenResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
		ExpiresAt:    tokens.ExpiresAt,
	}, nil
}
