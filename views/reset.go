package views

import "github.com/gofiber/fiber/v2"

func Reset(c *fiber.Ctx) error {
	return c.SendString(c.Query("token"))
}
