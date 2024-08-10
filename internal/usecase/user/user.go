package usecase_user

import (
	"context"
	"github.com/MovingTowardsADream/SneakerStore-UserService/internal/entity"
	"log/slog"
)

type UseCaseUser struct {
	log                *slog.Logger
	addUserProfileRepo AddUserProfileRepo
}

func NewUseCaseUser(l *slog.Logger, addUserProfileRepo AddUserProfileRepo) *UseCaseUser {
	return &UseCaseUser{log: l, addUserProfileRepo: addUserProfileRepo}
}

func (ucu *UseCaseUser) AddUserProfile(ctx context.Context, profile *entity.Profile) error {
	return ucu.addUserProfileRepo.UpdateUserProfile(ctx, profile)
}
