package routes

import (
	"text2doc/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/doc", handler.GenerateDoc)
}
