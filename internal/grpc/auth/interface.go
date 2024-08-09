package auth

import (
	"context"
	"github.com/MovingTowardsADream/SneakerStore-UserService/internal/model"
)

type Authorization interface {
	CreateUser(ctx context.Context, user model.User) (int64, error)
	GenerateToken(ctx context.Context, phone, password string) (string, error)
}
