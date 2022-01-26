package unsplash

import (
	"github.com/esequiel378/unsplashed-bff/services/unsplash/api"
	"github.com/gofiber/fiber/v2"
)

func photos(c *fiber.Ctx) error {
	queryString := c.Request().URI().QueryString()
	params := "?" + string(queryString)

	photos, err := api.PhotosList(params)

	if err != nil {
		c.Status(fiber.StatusBadRequest).SendString(err.Error())
		return err
	}

	return c.Status(fiber.StatusOK).JSON(photos)
}

func topics(c *fiber.Ctx) error {
	queryString := c.Request().URI().QueryString()
	params := "?" + string(queryString)

	topics, err := api.TopicsList(params)

	if err != nil {
		c.Status(fiber.StatusBadRequest).SendString(err.Error())
		return err
	}

	return c.Status(fiber.StatusOK).JSON(topics)
}

func topicsPhotos(c *fiber.Ctx) error {
	queryString := c.Request().URI().QueryString()
	params := "?" + string(queryString)
    slug := c.Params("slug")

	topics, err := api.TopicsPhotos(slug, params)

	if err != nil {
		c.Status(fiber.StatusBadRequest).SendString(err.Error())
		return err
	}

	return c.Status(fiber.StatusOK).JSON(topics)
}



func searchPhotos(c *fiber.Ctx) error {
	queryString := c.Request().URI().QueryString()
	params := "?" + string(queryString)

	photos, err := api.PhotosSearch(params)

	if err != nil {
		c.Status(fiber.StatusBadRequest).SendString(err.Error())
		return err
	}

	return c.Status(fiber.StatusOK).JSON(photos)
}
