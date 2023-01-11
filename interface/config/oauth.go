package config

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/linkedin"
)

func NewOauthConfigGithub() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     GetConfig().GithubOAuthCredentials.ClientID,
		ClientSecret: GetConfig().GithubOAuthCredentials.ClientSecret,
		Scopes:       []string{"read:user", "user:email"},
		Endpoint:     github.Endpoint,
		RedirectURL:  "http://localhost:8080/oauth/redirect/github",
	}
}

func NewOauthConfigGoogle() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     GetConfig().GoogleOAuthCredentials.ClientID,
		ClientSecret: GetConfig().GoogleOAuthCredentials.ClientSecret,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
		RedirectURL:  "http://localhost:8080/oauth/redirect/google",
	}
}

func NewOAuthConfigLinkedIn() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     GetConfig().LinkedInOAuthCredentials.ClientID,
		ClientSecret: GetConfig().LinkedInOAuthCredentials.ClientSecret,
		Scopes:       []string{"r_liteprofile", "r_emailaddress"},
		Endpoint:     linkedin.Endpoint,
		RedirectURL:  "http://localhost:8080/oauth/redirect/linkedin",
	}
}
