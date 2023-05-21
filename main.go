package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"hello": "world",
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // Default port if not specified
	}

	log.Fatal(app.Listen(":" + port))
}