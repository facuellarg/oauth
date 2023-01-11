package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"oauth/domain/entities"
	"oauth/interface/config"

	"github.com/go-resty/resty/v2"
)

type OAuthServiceGoogle struct {
	getterTokenI
}

func NewGoogleOAuthService() *OAuthServiceGoogle {
	return &OAuthServiceGoogle{
		getterTokenI: NewOAuthConfig(config.NewOauthConfigGoogle()),
	}
}

func (o *OAuthServiceGoogle) GetInfoUser(ctx context.Context, accessToken string) (*entities.OauthUser, error) {
	url := fmt.Sprintf("https://www.googleapis.com/oauth2/v2/userinfo?access_token=%s", accessToken)
	client := resty.New()
	var userInfo entities.OauthUser
	response, err := client.R().Get(url)
	if err != nil {
		return nil, err
	}
	userInfoMap := make(map[string]interface{})
	if err := json.Unmarshal(response.Body(), &userInfoMap); err != nil {
		return nil, err
	}
	userInfo.Email = userInfoMap["email"].(string)
	userInfo.Image = userInfoMap["picture"].(string)
	userInfo.Name = userInfoMap["name"].(string)

	return &userInfo, nil

}
