package controller

import (
	"errors"
	"oauth/domain/entities"
	"oauth/use-case/service"
)

type UserController struct {
	UserService service.UserServiceI
}

var (
	ErrEmptyEmail = errors.New("empty email")
)

func NewUserController(userService service.UserServiceI) *UserController {
	return &UserController{
		UserService: userService,
	}
}

func (uc *UserController) SaveUser(ctx Context) error {
	user := &entities.User{}
	if err := ctx.Bind(user); err != nil {
		return err
	}
	return uc.UserService.SaveUser(ctx.Context(), user)
}

func (uc *UserController) GetUser(ctx Context) (*entities.User, error) {
	email := ctx.Query("email")
	if email == "" {
		return nil, ErrEmptyEmail
	}
	return uc.UserService.GetUser(ctx.Context(), email)
}
