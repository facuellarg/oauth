package repository

import (
	"context"
	"oauth/domain/entities"
)

type OauthServiceI interface {
	GetAccessToken(ctx context.Context, code string) (string, error)
	GetInfoUser(ctx context.Context, accessToken string) (*entities.OauthUser, error)
}
