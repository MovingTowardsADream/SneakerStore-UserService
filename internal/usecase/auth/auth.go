package usecase_auth

import (
	"context"
	"github.com/MovingTowardsADream/SneakerStore-UserService/internal/dto"
	"github.com/MovingTowardsADream/SneakerStore-UserService/internal/entity"
	"log/slog"
)

type UseCaseAuth struct {
	log      *slog.Logger
	userRepo UserRepository
}

func NewUseCaseAuth(l *slog.Logger, userRepo UserRepository) *UseCaseAuth {
	return &UseCaseAuth{log: l, userRepo: userRepo}
}

func (uca *UseCaseAuth) CreateUser(ctx context.Context, user *entity.User) (int64, error) {
	return uca.userRepo.CreateUser(ctx, user)
}

func (uca *UseCaseAuth) GenerateToken(ctx context.Context, signIn *dto.SignInRequest) (string, error) {
	_, err := uca.userRepo.GetUser(ctx, signIn)
	return "test", err
}
