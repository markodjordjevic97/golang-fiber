package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-example/util"
)

func IsAuthenticated(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	if _, err := util.ParseJWT(cookie); err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthenticated",
		})
	}
	return c.Next() // go to the next routes
}
