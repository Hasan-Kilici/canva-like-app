package routes

import(
	"github.com/gofiber/fiber/v2"
	"practice/handlers"
)

func Service(app fiber.Router) {
	app.Post("/convert", handlers.Convert)
}