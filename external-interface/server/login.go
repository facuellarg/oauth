package server

import (
	"oauth/interface/config"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
)

type loginInfo struct {
	URL         string
	Description string
	Icon        string
}

func NewOauthInfo(description, icon string, config *oauth2.Config) loginInfo {
	info := loginInfo{
		Description: description,
		URL:         config.AuthCodeURL(""),
		Icon:        icon,
	}
	return info
}

func Index(c *fiber.Ctx) error {
	info := make([]loginInfo, 0)
	info = append(info, NewOauthInfo("Login with Github", "github-square", config.NewOauthConfigGithub()))
	info = append(info, NewOauthInfo("Login with Google", "google", config.NewOauthConfigGoogle()))
	return c.Render("index", info)
}
