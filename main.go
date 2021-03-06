package main

import (
	"log"
	"os"

	"github.com/esequiel378/unsplashed-bff/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	app := fiber.New()

	if origins := os.Getenv("ALLOW_ORIGINS"); origins != "" {
		app.Use(cors.New(cors.Config{
			AllowOrigins: origins,
		}))
	}

	servicesRoutes := services.GetRoutes()

	app.Mount("/api", servicesRoutes)

	log.Fatal(app.Listen(":80"))
}
