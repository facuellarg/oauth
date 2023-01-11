package repository

import (
	"context"
	"oauth/domain/entities"
)

type UserRepositoryI interface {
	GetUserByEmail(ctx context.Context, email string) (*entities.User, error)
	SaveUser(ctx context.Context, user *entities.User) error
}
