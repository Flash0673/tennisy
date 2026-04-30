package login

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"tennisly.com/mvp/internal/modules/auth/action/login/dto"
	"tennisly.com/mvp/internal/modules/auth/domain/refresh_token"
	"tennisly.com/mvp/internal/modules/auth/domain/user"
	dtoModule "tennisly.com/mvp/internal/modules/auth/dto"
	dtoToken "tennisly.com/mvp/pkg/token/dto"
)

//go:generate mockgen -destination=mocks/mocks.go -package=mocks . PasswordHasher,TokenService,RefreshTokenRepository,UserRepository

type PasswordHasher interface {
	Compare(password, hash string) bool
}

type TokenService interface {
	Generate(userID, email string) (*dtoToken.TokenPair, error)
}

type RefreshTokenRepository interface {
	Create(ctx context.Context, token *refresh_token.RefreshToken) error
}

type UserRepository interface {
	GetByEmail(ctx context.Context, email string) (*user.User, error)
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

// Login TODO разбить на 2 сервиса - отдельно юзер, отдельно токен и засунуть под транзакцию
func (s *Service) Login(ctx context.Context, req dto.LoginRequest) (*dtoModule.TokenResponse, error) {

	usr, err := s.userRepository.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	if !s.passwordHasher.Compare(req.Password, usr.PasswordHash) {
		return nil, errors.New("invalid password")
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

	return &dtoModule.TokenResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
		ExpiresAt:    tokens.ExpiresAt,
	}, nil
}
