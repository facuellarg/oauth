package service

import (
	"context"
	"oauth/domain/entities"
	"oauth/use-case/repository"
)

type UserService struct {
	UserRepository repository.UserRepositoryI
}

func (us *UserService) SaveUser(ctx context.Context, user *entities.User) error {
	return us.UserRepository.SaveUser(ctx, user)
}

func (us *UserService) GetUser(ctx context.Context, email string) (*entities.User, error) {
	return us.UserRepository.GetUserByEmail(ctx, email)
}
