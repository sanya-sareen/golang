package main

import "github.com/gofiber/fiber/v2"

func Home(c *fibre.Ctx) error {
	app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World- sanya sareen!")
    })
}