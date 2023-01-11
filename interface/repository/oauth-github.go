package repository

import (
	"context"
	"fmt"
	"oauth/domain/entities"
	"oauth/interface/config"

	githubClient "github.com/google/go-github/v45/github"
	"golang.org/x/oauth2"
)

type OAuthServiceGithub struct {
	getterTokenI
}

func NewGithubOAuth() *OAuthServiceGithub {
	return &OAuthServiceGithub{
		getterTokenI: NewOAuthConfig(config.NewOauthConfigGithub()),
	}
}

func (o *OAuthServiceGithub) GetInfoUser(ctx context.Context, accessToken string) (*entities.OauthUser, error) {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := githubClient.NewClient(tc)
	gitUser, _, err := client.Users.Get(ctx, "")
	if err != nil {
		return nil, err
	}

	userInfo := new(entities.OauthUser)

	if gitUser.Email != nil {
		userInfo.Email = *gitUser.Email
	} else {
		emails, _, err := client.Users.ListEmails(ctx, &githubClient.ListOptions{})
		if err != nil {
			return nil, err
		}
		if len(emails) == 0 {
			return nil, fmt.Errorf("no email found")
		}
		for _, email := range emails {
			if email.GetPrimary() {
				userInfo.Email = email.GetEmail()
				break
			}
		}

	}
	userInfo.Image = gitUser.GetAvatarURL()
	userInfo.Name = gitUser.GetLogin()
	return userInfo, nil

}
