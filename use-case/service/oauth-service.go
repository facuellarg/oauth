package service

import (
	"context"
	"oauth/domain/entities"
	"oauth/use-case/repository"
)

type OauthService struct {
	OauthRepository repository.OAuthI
}

func NewOauthService(oauthRepository repository.OAuthI) *OauthService {
	return &OauthService{
		OauthRepository: oauthRepository,
	}
}

func (s *OauthService) GetInfoUser(ctx context.Context, code string) (*entities.OauthUser, error) {

	token, err := s.OauthRepository.GetAccessToken(ctx, code)
	if err != nil {
		return nil, err
	}
	return s.OauthRepository.GetInfoUser(ctx, token)
}
