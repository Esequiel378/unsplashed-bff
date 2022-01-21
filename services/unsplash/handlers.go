package unsplash

import "github.com/gofiber/fiber/v2"

func greeting(c *fiber.Ctx) error {
	return c.SendString("Hello world")
}
