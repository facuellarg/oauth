package repository

import (
	"context"

	"golang.org/x/oauth2"
)

type getterTokenI interface {
	GetAccessToken(ctx context.Context, code string) (string, error)
}

type oAuthService struct {
	config *oauth2.Config
}

func NewOAuthConfig(config *oauth2.Config) *oAuthService {

	return &oAuthService{
		config: config,
	}
}

func (o *oAuthService) GetAccessToken(ctx context.Context, code string) (string, error) {
	token, err := o.config.Exchange(ctx, code)
	if err != nil {
		return "", err
	}
	return token.AccessToken, nil
}
