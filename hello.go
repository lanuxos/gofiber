package main

import (
    "github.com/gofiber/fiber/v2"
    "fmt"
    "encoding/json"
)

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
    app.Get("/apis/*", func(c *fiber.Ctx) error {
        return c.SendString("API path: " + c.Params("*"))
    })

    // Static
    app.Static("/", "./static")

    // Router handlers
    app.Get("/api/simpleGet", func(c *fiber.Ctx) error {
        return c.SendString("I am simple get request")
    })
    app.Post("/api/simplePost", func(c *fiber.Ctx) error {
        return c.SendString("I am simple post request")
    })
    app.Use("/api", func(c *fiber.Ctx) error {
        return c.Next()
    })

    // Mount [mouthing fiber instance]
    micro := fiber.New()
    app.Mount("/mountJohn", micro) // GET /mountJohn/doe -> 200 OK
    micro.Get("/doe", func(c *fiber.Ctx) error {
        return c.SendStatus(fiber.StatusOK)
    })

    // MountPath [contains one or more path patterns on which a sub-app was mounted]
    mountPathApp := fiber.New()
    one := fiber.New()
    two := fiber.New()
    three := fiber.New()

    two.Mount("/three", three)
    one.Mount("two", two)
    mountPathApp.Mount("/one", one)

    one.MountPath() // "/one"
    two.MountPath() // "/one/two"
    three.MountPath() // "/one/two/three"
    mountPathApp.MountPath() // ""

    // Group [route grouping]
    apiGroup := app.Group("/apigroup", func(c *fiber.Ctx) error{
        return c.Next()
    })
    g1 := apiGroup.Group("/g1", func(c *fiber.Ctx) error{
        return c.Next()
    })
    g1.Get("/list", func(c *fiber.Ctx) error{ // /apiGroup/g1/list
        return c.SendString("group1list")
    })
    g1.Get("/user", func(c *fiber.Ctx) error{ // /apiGroup/g1/user
        return c.SendString("group1user")
    })
    g2 := apiGroup.Group("/g2", func(c *fiber.Ctx) error{
        return c.Next()
    })
    g2.Get("/list", func(c *fiber.Ctx) error{ // /apiGroup/g2/list
        return c.SendString("group2list")
    })
    g2.Get("/user", func(c *fiber.Ctx) error{ // /apiGroup/g2/user
        return c.SendString("group2user")
    }) 

    // Route
    app.Route("/routePrefix", func(api fiber.Router){
        api.Get("/foo", func(c *fiber.Ctx) error { // routePrefix/foo (name:routePrefix.foo)
            return c.SendString("route prefix with foo")
        }).Name("foo")
        api.Get("/bar", func(c *fiber.Ctx) error { // routePrefix/bar (name:routePrefix.bar)
            return c.SendString("route prefix with bar")
        }).Name("bar")
    }, "routeTest")

    // Server
    app.Server().MaxConnsPerIP=1

    // Stack
    var handler = func(c *fiber.Ctx) error {return nil}
    app.Get("/john/:age", handler)
    app.Post("/register", handler)

    stackData, _ := json.MarshalIndent(app.Stack(), "", " ")
    fmt.Println("STACK", string(stackData))

    // Name
    app.Get("/appName", handler)
    app.Name("index")
    app.Get("/nameDoe", handler).Name("name doe")
    app.Trace("/tracer", handler).Name("tracer")
    app.Delete("/delete", handler).Name("delete")

    a := app.Group("/a")
    a.Name("fd.")

    a.Get("/test", handler).Name("test")

    nameData, _ := json.MarshalIndent(app.Stack(), "", " ")
    fmt.Print("NAME", string(nameData))

    // GetRoute
    app.Get("/", handler).Name("getRoute")
    getRouteData, _ := json.MarshalIndent(app.GetRoute("getRoute"), "", " ")
    fmt.Print("GET_ROUTE", string(getRouteData))

    // GetRoutes
    app.Post("/", func (c *fiber.Ctx) error {
        return c.SendString("GetRoutesMethod")
    }).Name("index")
    getRoutesData, _ := json.MarshalIndent(app.GetRoutes(true), "", " ")
    fmt.Print("GET_ROUTES", string(getRoutesData))

    // Listen
    app.Listen(":80")
}
