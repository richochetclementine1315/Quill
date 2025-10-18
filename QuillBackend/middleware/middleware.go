package middleware

import (
	"QuillBackend/utils"

	"github.com/gofiber/fiber/v2"
)

func isAuthenticate(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	if _, err := utils.ParseJWT(cookie); err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{"message": "unauthenticated"})
	}
	return c.Next()

}
