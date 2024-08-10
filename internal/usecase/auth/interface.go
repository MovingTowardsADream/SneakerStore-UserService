package usecase_auth

import (
	"context"
	"github.com/MovingTowardsADream/SneakerStore-UserService/internal/dto"
	"github.com/MovingTowardsADream/SneakerStore-UserService/internal/entity"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *entity.User) (int64, error)
	GetUser(ctx context.Context, signIn *dto.SignInRequest) (*entity.User, error)
}
