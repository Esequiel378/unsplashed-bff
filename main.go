package main

import (
	"log"
	"os"
	"time"

	"github.com/esequiel378/unsplashed-bff/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	if origins := os.Getenv("ALLOW_ORIGINS"); origins != "" {
		app.Use(cors.New(cors.Config{
			AllowOrigins: origins,
		}))
	}

	app.Use(cache.New(cache.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.Query("refresh") == "true"
		},
		Expiration:   30 * time.Minute,
		CacheControl: true,
	}))

	servicesRoutes := services.GetRoutes()

	app.Mount("/api", servicesRoutes)

	log.Fatal(app.Listen(":4000"))
}
