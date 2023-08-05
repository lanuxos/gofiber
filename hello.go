package main

import "github.com/gofiber/fiber/v2"

func main() {
    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello world from go fiber version 2")
    })

    // GET http://localhost/hello%20world
    app.Get("/value/:value", func(c *fiber.Ctx) error {
        return c.SendString("value: " + c.Params("value"))
    })

    // GET http://localhost/john
    app.Get("/person/:name?", func(c *fiber.Ctx) error {
        if c.Params("name") != "" {
            return c.SendString("Hello " + c.Params("name"))
        }
        return c.SendString("Hello anonymous")
    })

    // GET http://localhost/api/user/john
    app.Get("/api/*", func(c *fiber.Ctx) error {
        return c.SendString("API path: " + c.Params("*"))
    })

    app.Listen(":80")
}
