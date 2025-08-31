package api

import (
	"github.com/gofiber/fiber/v2"

	"github.com/ardwiinoo/snap-bi-sim/signator-generator/api/handler"
	"github.com/ardwiinoo/snap-bi-sim/signator-generator/pkg/signer"
)

func SetupRoutes(app *fiber.App, svc *signer.Service) {
	app.Post("/v1/sign", handler.NewSignHandler(svc))
}