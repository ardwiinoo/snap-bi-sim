package handler

import (
	"github.com/gofiber/fiber/v2"

	"github.com/ardwiinoo/snap-bi-sim/signator-generator/pkg/signer"
)

func NewSignHandler(s *signer.Service) fiber.Handler {
	return func (c *fiber.Ctx) error {
		var req signer.SignRequest

		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(
				fiber.Map{
					"error": err.Error(),
				},
			)
		}

		resp, err := s.GenerateSignature(req)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(
				fiber.Map{
					"error": err.Error(),
				},
			)
		}

		return c.JSON(resp)
	}
}