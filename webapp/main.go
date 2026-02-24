package main

import (
	"time"

	"encoding/json"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		payload := fiber.Map{
			"message":   "My name is Carolyn Thai",
			"timestamp": time.Now().UnixMilli(),
		}

		c.Set("Content-Type", "application/json")

		err := json.NewEncoder(c.Response().BodyWriter()).Encode(payload)
		if err != nil {
			return c.Status(500).SendString("Encoder Error")
		}

		return nil
	})

	app.Listen(":3000")
}
