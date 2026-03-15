package routes

import (
	"msg_queue/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/health", handlers.GetHealth)
}
