package grpc_auth

import (
	"context"
	"github.com/MovingTowardsADream/SneakerStore-UserService/internal/dto"
	"github.com/MovingTowardsADream/SneakerStore-UserService/internal/entity"
)

type Authorization interface {
	CreateUser(ctx context.Context, user *entity.User) (int64, error)
	GenerateToken(ctx context.Context, signIn *dto.SignInRequest) (string, error)
}
