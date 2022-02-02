package controller

import "github.com/gofiber/fiber/v2"

func Welcome(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status": "SUCCESS",
		"code":   200,
	})
}
