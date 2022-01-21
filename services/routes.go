package services

import (
	"github.com/esequiel378/unsplashed-bff/services/unsplash"
	"github.com/gofiber/fiber/v2"
)

func GetRoutes() *fiber.App {
	app := fiber.New()

	unsplashRoutes := unsplash.GetRoutes()

	app.Mount("/unsplash", unsplashRoutes)

	return app
}
