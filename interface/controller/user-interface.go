package controller

import "oauth/domain/entities"

type UserControllerI interface {
	SaveUser(ctx Context) error
	GetUser(ctx Context) (*entities.User, error)
}
