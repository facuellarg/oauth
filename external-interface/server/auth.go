package server

import (
	"oauth/interface/repository"

	"github.com/gofiber/fiber/v2"
)

type oAuthServer struct {
	oauthService repository.OauthServiceI
}

func NewOAuthServer(oauthService repository.OauthServiceI) *oAuthServer {
	return &oAuthServer{
		oauthService: oauthService,
	}
}

func (s *oAuthServer) AuthServer(ctx *fiber.Ctx) error {
	code := ctx.Query("code")
	if code == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":             "invalid_request",
			"error_description": "Missing or invalid code",
		})
	}
	token, err := s.oauthService.GetAccessToken(ctx.Context(), code)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":             "invalid_request",
			"error_description": err.Error(),
		})
	}
	user, err := s.oauthService.GetInfoUser(ctx.Context(), token)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":             "invalid_request",
			"error_description": err.Error(),
		})
	}
	return ctx.JSON(user)

}
