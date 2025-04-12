package router

import (
	"github.com/Satr10/fufufafa-api/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App) {

	app.Get("/", handlers.AllFufufafa)
	api := app.Group("/api")
	api.Get("/", handlers.AllFufufafa)

}
