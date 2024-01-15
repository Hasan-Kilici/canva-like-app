package main

import(
	"practice/routes"
	"practice/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/goccy/go-json"
)

func main(){
	app := fiber.New(fiber.Config{
        JSONEncoder: json.Marshal,
        JSONDecoder: json.Unmarshal,
    })

	app.Use(middleware.Compress)
	app.Use(middleware.Security)

	api := app.Group("/api", middleware.RateLimit)
	service := app.Group("/service", middleware.Cors)
	server := app.Group("/server")

	routes.Api(api)
	routes.Server(server)
	routes.Service(service)

	err := app.Listen("127.0.0.1:3000")
    if err != nil {
        panic(err)
    }
}