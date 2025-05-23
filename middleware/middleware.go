package middleware

import "github.com/gofiber/fiber/v2"

func NotFound(c *fiber.Ctx) error {
	return c.Status(404).JSON(fiber.Map{
		"error":   true,
		"message": "Resource Not Found",
	})
}
