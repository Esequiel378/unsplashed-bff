package unsplash

import "github.com/gofiber/fiber/v2"

func GetRoutes() *fiber.App {
    app := fiber.New()

    app.Get("/photos", photos)
    app.Get("/search/photos", searchPhotos)

    return app
}
