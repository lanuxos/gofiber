# gofiber
[Go Fiber Docs](https://github.com/lanuxos/gofiber.git)

# Installation
`go get github.com/gofiber/fiber/v2`
# Zero allocation
- copy
- utils.CopyString
# Hello world
```
package main

import "github.com/gofiber/fiber/v2"

func main() {
    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello world")
    })

    app.Listen(":80")
}
```
# Basic routing
```
app.METHOD(PATH string, ...func(*fiber.Ctx) error)
// app -- is an instance of Fiber
// METHOD -- is an HTTP request method: GET, POST, PUT, DELETE...
// PATH -- is a virtual path on the server
// func(*fiber.Ctx) error -- is a callback function 
// containing the Context executed when 
// the route is matched

app.Get("/", func(c *func.Ctx) error {
    return c.SendString("Hello world")
})
```
- Parameter
```
// GET http://localhost/hello%20world
app.Get("/:value", func(c *fiber.Ctx) error {
    return c.SendString("value: " + c.Params("value"))
})
```
- Optional Parameter
```
// GET http://localhost/john
    app.Get("/person/:name?", func(c *fiber.Ctx) error {
        if c.Params("name") != "" {
            return c.SendString("Hello " + c.Params("name"))
        }
        return c.SendString("Hello anonymous")
    })
```
- Wildcard
```
// GET http://localhost/api/user/john
    app.Get("/api/*", func(c *fiber.Ctx) error {
        return c.SendString("API path: " + c.Params("*"))
    })
```
# Static files
