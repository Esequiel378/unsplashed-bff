package main

import (
	"log"

	"github.com/esequiel378/unsplashed-bff/services"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	servicesRoutes := services.GetRoutes()

	app.Mount("/api", servicesRoutes)

	log.Fatal(app.Listen(":4000"))
}
