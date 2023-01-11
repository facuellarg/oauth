package controller

import (
	"errors"
	"oauth/domain/entities"
	"oauth/use-case/service"
)

type AppController struct {
	UserService        service.UserServiceI
	GoogleOAuthService service.OauthServiceI
	GithubOAuthService service.OauthServiceI
}

var (
	ErrEmptyCode    = errors.New("empty code")
	ErrInvalidCode  = errors.New("invalid code")
	ErrInvalidToken = errors.New("invalid token")
)

func NewAppController(
	userService service.UserServiceI,
	googleOAuthService service.OauthServiceI,
	githubOAuthService service.OauthServiceI,
) *AppController {
	return &AppController{
		UserService:        userService,
		GoogleOAuthService: googleOAuthService,
		GithubOAuthService: githubOAuthService,
	}
}

func (ac *AppController) genericRedirect(ctx Context, oauthService service.OauthServiceI) (*entities.OauthUser, error) {
	code := ctx.Query("code")
	if code == "" {
		return nil, ErrEmptyCode
	}
	userInfo, err := oauthService.GetInfoUser(ctx.Context(), code)
	if err != nil {
		return nil, err
	}
	return userInfo, nil
}

func (ac *AppController) RedirectGoogle(ctx Context) (*entities.OauthUser, error) {
	return ac.genericRedirect(ctx, ac.GoogleOAuthService)
}

func (ac *AppController) RedirectGithub(ctx Context) (*entities.OauthUser, error) {
	return ac.genericRedirect(ctx, ac.GithubOAuthService)
}
