package main

import (
	"fmt"
	"oauth/external-interface/server"
	"oauth/interface/config"
	"oauth/interface/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
)

func main() {

	htmlEngine := html.New("./public", ".html")
	app := fiber.New(fiber.Config{
		AppName: "oauth-server",
		Views:   htmlEngine,
	})

	gitService := repository.NewGithubOAuth()
	githubServer := server.NewOAuthServer(gitService)

	googleService := repository.NewGoogleOAuthService()
	googleServer := server.NewOAuthServer(googleService)

	config := config.GetConfig()
	fmt.Printf("config: %v\n", config)

	app.Use(logger.New())
	app.Get("/", server.Index)

	group := app.Group("/oauth/redirect")
	group.Get("/github", githubServer.AuthServer)
	group.Get("/google", googleServer.AuthServer)

	app.Listen(":8080")
}
