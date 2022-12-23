package error_handler

import (
	"github.com/gofiber/fiber/v2"
)

func New() func(c *fiber.Ctx, err error) error {
	return func(c *fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError
		return c.Status(code).JSON(map[string]interface{}{
			"code":    code,
			"success": false,
			"message": err.Error(),
		})
	}
}