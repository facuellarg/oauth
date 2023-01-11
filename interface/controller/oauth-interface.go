package controller

import "oauth/domain/entities"

type ControllerI interface {
	RedirectGoogle(ctx Context) (*entities.OauthUser, error)
	RedirectGithub(ctx Context) (*entities.OauthUser, error)
}
