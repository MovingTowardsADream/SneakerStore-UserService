package user

import (
	"context"
	"github.com/MovingTowardsADream/SneakerStore-UserService/internal/model"
)

type UserProfileInfo interface {
	AddUserProfile(ctx context.Context, profile model.Profile) error
}
