package routes

import(
	"github.com/gofiber/fiber/v2"
	"practice/handlers"
)

func Api(app fiber.Router) {
	app.Post("/convert", handlers.Convert)
}