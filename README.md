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
`app.Static(prefix, root string, config ...Static)`
```
app := fiber.New()
app.Static("/", "./static")
app.Listen(":3000")
```
# API
## Fiber
- New
```
func New(config ...Config) *App

app := fiber.New()
```
- Config
```
app := fiber.New(fiber.Config{
    AppName: "Test Go Fiber App V 0.0.1",
    Prefork: true,
    StrictRouting: true,
    CaseSensitive: true,
    ServerHearder: "Fiber",
})
```
- NewError
```
func NewError(code int, message ...string) *Error

app.Get("/", func(c *fiber.Ctx) error {
    return fiber.NewError(782, "CUSTOM ERROR MESSAGE")
})
```
- IsChild [determines if the current process is a result of Prefork]
```
func IsChild() bool

// Prefork will spawn child processes
app := fiber.New(fiber.Config{
    Prefork: true,
})
if !fiber.IsChild() {
    fmt.PrintLn("I am the parent process")
} else {
    fmt.Println("I am a child process")
}
```
## App
- Static
```
func (app *App) Static(prefix, root string, config ...Static) Router

app.Static("/", "./static")
app.Static("/", "./public")

// or using virtual path prefix
app.Static("./static", "./public")

// Custom config
// Static defines configuration options when defining static assets.
type Static struct {
    // When set to true, the server tries minimizing CPU usage by caching compressed files.
    // This works differently than the github.com/gofiber/compression middleware.
    // Optional. Default value false
    Compress bool `json:"compress"`

    // When set to true, enables byte range requests.
    // Optional. Default value false
    ByteRange bool `json:"byte_range"`

    // When set to true, enables directory browsing.
    // Optional. Default value false.
    Browse bool `json:"browse"`

    // When set to true, enables direct download.
    // Optional. Default value false.
    Download bool `json:"download"`

    // The name of the index file for serving a directory.
    // Optional. Default value "index.html".
    Index string `json:"index"`

    // Expiration duration for inactive file handlers.
    // Use a negative time.Duration to disable it.
    //
    // Optional. Default value 10 * time.Second.
    CacheDuration time.Duration `json:"cache_duration"`

    // The value for the Cache-Control HTTP-header
    // that is set on the file response. MaxAge is defined in seconds.
    //
    // Optional. Default value 0.
    MaxAge int `json:"max_age"`

    // ModifyResponse defines a function that allows you to alter the response.
    //
    // Optional. Default: nil
    ModifyResponse Handler

    // Next defines a function to skip this middleware when returned true.
    //
    // Optional. Default: nil
    Next func(c *Ctx) bool
}

app.Static("/", "./public", fiber.Static{
    Compress: true,
    ByteRange: true,
    Browser: true,
    Index: "john.html",
    CacheDuration: 10 * time.Second,
    MaxAge: 3600,
})
```
- Route Handlers
```
// HTTP methods
func (app *App) Get(path string, handlers ...Handler) Router
func (app *App) Head(path string, handlers ...Handler) Router
func (app *App) Post(path string, handlers ...Handler) Router
func (app *App) Put(path string, handlers ...Handler) Router
func (app *App) Delete(path string, handlers ...Handler) Router
func (app *App) Connect(path string, handlers ...Handler) Router
func (app *App) Options(path string, handlers ...Handler) Router
func (app *App) Trace(path string, handlers ...Handler) Router
func (app *App) Patch(path string, handlers ...Handler) Router

// Add allows you to specify a method as value
func (app *App) Add(method, path string, handlers ...Handler) Router

// All will register the route on all HTTP methods
// Almost the same as app.Use but not bound to prefixes
func (app *App) All(path string, handlers ...Handler) Router

// Simple GET handler
app.Get("/api/list", func(c *fiber.Ctx) error{
    return c.SendString("I am GET request")
})

// Simple POST handler
app.get("/api/register", func(c *fiber.Ctx) error {
    return c.SendString("I am POST request)
})

// Use can be used for middleware packages and prefix catchers
func (app *App) Use(args ...interface{}) Router

// Match any request
app.Use(func(c *fiber.Ctx) error {
    return c.Next()
})

// Match request starting with /api [prefix]
app.Use("/api", func(c *fiber.Ctx) error {
    return c.Next()
})

// Match request starting with /api or /home [multiple-prefix support]
app.Use([]string{"/api", "/home"}, func(c *fiber.Ctx) error {
    return c.Next()
})

// Attach multiple handlers
app.Use("/api", func(c *fiber.Ctx) error {
    c.Set("X-Custom-Header", random.String(32))
    return c.Next()
}, func(c fiber.Ctx) error {
    return c.Next()
})
```
- Mount
`func (a *App) Mount(prefix string, app *App) Router`
```
app := fiber.New()
micro := fiber.New()
app.Mount("/mountJohn", micro) // GET /mountJohn/doe -> 200 OK
micro.Get("/doe", func(c *fiber.Ctx) error {
    return c.SendStatus(fiber.StatusOK)
})
```
- MountPath
`func (app *App) MountPath() string`
```
// MountPath [contains one or more path patterns on which a sub-app was mounted]
mountPathApp := fiber.New()
one := fiber.New()
two := fiber.New()
three := fiber.New()

two.Mount("/three", three)
one.Mount("two", two)
mouthPathApp.Mount("/one", one)

one.MouthPath() // "/one"
two.MouthPate() // "/one/two"
three.MouthPath() // "/one/two/three"
mouthPathApp.MouthPath() // ""
// mouthing order is important, mouth the deepest app first
```
- Group [create routes group]
`func (app *App) Group(prefix string, handlers ...Handler) Router`
```
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
```
- Route [define routes with a common prefix inside the common function]
`func (app *App) Route(prefix string, fn func(router Router), name ...string) Router`
```
app.Route("/routePrefix", func(api fiber.Router){
    api.Get("/foo", func(c *fiber.Ctx) error { // routePrefix/foo (name:routePrefix.foo)
        return c.SendString("route prefix with foo")
    }).Name("foo")
    api.Get("/bar", func(c *fiber.Ctx) error { // routePrefix/bar (name:routePrefix.bar)
        return c.SendString("route prefix with bar")
    }).Name("bar")
    }, "routeTest")
```
- Server
`func (app *App) Server()`
`app.Server().MaxConnsPerIP=1`
- Server Shutdown
```
func (app *App) Shutdown() error
func (app *App) ShutdownWithTimeout(timeout time.Duration) error
func (app *App) ShutdownWithContext(ctx content.Context) error
```
- HandlersCount [return amount of registered handlers]
`func (app *App) HandlersCount() uint32`
- Stack [return original router stack]
`func (app *App) Stack() [][]*Route`
```
import (
    "fmt"
    "encoding/json"
    )
var handler = func(c *fiber.Ctx) error {return nil}
    app.Get("/john/:age", handler)
    app.Post("/register", handler)

    data, _ := json.MarshalIndent(app.Stack(), "", " ")
    fmt.Println(string(data))
```
- Name [assign the name of created router]
`func (app *App) Name(name string) Router`
```
app.Get("/appName", handler)
app.Name("index")
app.Get("/nameDoe", handler).Name("name doe")
app.Trace("/tracer", handler).Name("tracer")
app.Delete("/delete", handler).Name("delete")

a := app.Group("/a")
a.Name("fd.")

a.Get("/test", handler).Name("test")

nameData, _ := json.MarshalIndent(app.Stack(), "", " ")
fmt.Print(string(nameData))
```
- GetRoute [get route by name]
`func (app *App) GetRoute(name string) Route`
```
app.Get("/", handler).Name("getRoute")
getRouteData, _ := json.MarshalIndent(app.GetRoute("getRoute"), "", " ")
fmt.Print("GET_ROUTE", string(getRouteData))
```
- GetRoutes [get all routes]
`func (app *App) GetRoutes(filterUseOption ...bool) []Route`
```
app.Post("/", func (c *fiber.Ctx) error {
    return c.SendString("GetRoutesMethod")
}).Name("index")
getRoutesData, _ := json.MarshalIndent(app.GetRoutes(true), "", " ")
fmt.Print("GET_ROUTES", string(getRoutesData))
```
- Config [return read-only app configuration]
`func (app *App) Config() Config`
- Handler [return server handler that can be used to serve custom *fasthttp.RequestCtx requests]
`func (app *App) Handler() fasthttp.RequestHandler`
- Listen [listen serves HTTP requests from the given address]
`func (app *App) Listen(addr string) error`
```
// listen on port :8080
app.Listen(":8080")

// custom host
app.Listen("127.0.0.1:8080")
```
- ListenTLS [serves HTTPs requests from the given address using certFile and keyFile paths to as TLS certificate and key file]
`func (app *App) ListenTLS(addr, certFile, keyFile string) error`
`app.ListenTLS(":443", "./cert.pen", "./cert.key");`
- ListenTLSWithCertificate
`func (app *App) ListenTLSWithCertificate(addr string, cert tls.Certificate) error`
`app.ListenTLSWithCertificate(":443", cert);`
- ListenMutualTLS
`func (app *App) ListenMutualTLS(addr, certFile, keyFile, clientCertFile string) error`
`app.ListenMutualTLS(":443", "./cert.pen", "./cert.key", "./ca-chain-cert.pen");`
- ListenMutualTLSWithCertificate
`func (app *App) ListenMutualTLSWithCertificate(addr string, cert tls.Certificate, clientCertPool *x509.CertPool) error`
`app.ListenMutualTLSWithCertificate(":443", cert, clientCertPool);`
- Listener
`func (app *App) Listener(ln net.Listener) error`
```
ln, _ := net.Listen("tcp", ":3000")
cer, _ := lts.LoadX509KeyPar("server.crt", "server.key")
ln = tls.NewListener(ln, &tls.Config{Certificates: []tls.Certificate{cer}})
app.Listener(ln)
```
- Test
`func (app *App) Test(req *http.Request, msTimeout ...int) (*http.Response, error)`
```
// Create route with GET method for test:
app.Get("/", func(c *fiber.Ctx) error {
  fmt.Println(c.BaseURL())              // => http://google.com
  fmt.Println(c.Get("X-Custom-Header")) // => hi

  return c.SendString("hello, World!")
})

// http.Request
req := httptest.NewRequest("GET", "http://google.com", nil)
req.Header.Set("X-Custom-Header", "hi")

// http.Response
resp, _ := app.Test(req)

// Do something with results:
if resp.StatusCode == fiber.StatusOK {
  body, _ := ioutil.ReadAll(resp.Body)
  fmt.Println(string(body)) // => Hello, World!
}
```
- Hooks [return hooks property]
`func (app *App) Hooks() *Hooks`
## Ctx
## Constants
## Client
## Log
## Middleware
# Guide
## Routing
## Grouping
## Templates
## Error Handling
## Validation
## Hooks
## Make Fiber Faster
# Extra
## FAQ
## Benchmarks