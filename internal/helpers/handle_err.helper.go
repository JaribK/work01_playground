package helpers

import (
	"github.com/gofiber/fiber/v2"
)

func ErrResponse(c *fiber.Ctx, statusCode int, titleError string, errorDetails string) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"error":   titleError,
		"message": errorDetails,
	})
}
