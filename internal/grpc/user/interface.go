package grpc_user

import (
	"context"
	"github.com/MovingTowardsADream/SneakerStore-UserService/internal/entity"
)

type UserProfileInfo interface {
	AddUserProfile(ctx context.Context, profile *entity.Profile) error
}
