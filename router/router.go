package router

import (
	"github.com/Satr10/fufufafa-api/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App) {

	app.Get("/", handlers.Index)

	// group route for api
	api := app.Group("/api")

	api.Get("/", handlers.AllFufufafa)
	api.Get("/:quote_id<int>?", handlers.SingleFufufafa)

}
