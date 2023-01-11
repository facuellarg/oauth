package service

import (
	"context"
	"oauth/domain/entities"
)

type UserServiceI interface {
	SaveUser(ctx context.Context, user *entities.User) error
	GetUser(ctx context.Context, email string) (*entities.User, error)
}
