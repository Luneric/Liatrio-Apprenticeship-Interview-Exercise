package main

import (
	"time"

	"encoding/json"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
	})

	app.Get("/", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message":   "My name is Carolyn Thai",
			"timestamp": time.Now().UnixMilli(),
		})
	})

	app.Listen(":3000")
}
