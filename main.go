package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Static("/", "./public")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/user/:name", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("Hello %s, Welcome to the Fiber!", c.Params("name"))
		return c.SendString(msg)
	})

	app.Listen(":3000")
}
