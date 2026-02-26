package main

import (
	"time"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message":   "My name is Carolyn Thai",
			"timestamp": time.Now().UnixMilli(),
			"status":    "I hope this works",
		})
	})

	app.Listen(":3000")
}
