package service

import (
	"context"
	"oauth/domain/entities"
)

type OauthServiceI interface {
	GetInfoUser(ctx context.Context, code string) (*entities.OauthUser, error)
}
