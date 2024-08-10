package usecase_user

import (
	"context"
	"github.com/MovingTowardsADream/SneakerStore-UserService/internal/entity"
)

type AddUserProfileRepo interface {
	UpdateUserProfile(ctx context.Context, profile *entity.Profile) error
}
