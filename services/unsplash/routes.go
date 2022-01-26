package unsplash

import "github.com/gofiber/fiber/v2"

func GetRoutes() *fiber.App {
    app := fiber.New()

    app.Get("/photos", photos)
    app.Get("/search/photos", searchPhotos)
    app.Get("/topics", topics)
    app.Get("/topics/:slug/photos", topicsPhotos)

    return app
}
