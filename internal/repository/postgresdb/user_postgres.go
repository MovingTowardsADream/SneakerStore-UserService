package postgresdb

import (
	"context"
	"github.com/MovingTowardsADream/SneakerStore-UserService/internal/dto"
	"github.com/MovingTowardsADream/SneakerStore-UserService/internal/entity"
	"github.com/MovingTowardsADream/SneakerStore-UserService/pkg/postgres"
)

type UserRepo struct {
	db *postgres.Postgres
}

func NewUserRepo(db *postgres.Postgres) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (ur *UserRepo) CreateUser(ctx context.Context, user *entity.User) (int64, error) {
	return 1, nil
}

func (ur *UserRepo) GetUser(ctx context.Context, signIn *dto.SignInRequest) (*entity.User, error) {
	return nil, nil
}

func (ur *UserRepo) UpdateUserProfile(ctx context.Context, profile *entity.Profile) error {
	return nil
}
