package unsplash

import "github.com/gofiber/fiber/v2"

func GetRoutes() *fiber.App {
    app := fiber.New()

    app.Get("/greeting", greeting)

    return app
}
