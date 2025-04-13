package router

import (
	"time"

	"github.com/Satr10/fufufafa-api/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
)

func SetupRouter(app *fiber.App) {

	app.Get("/", handlers.Index)

	// group route for api
	api := app.Group("/api")

	api.Get("/", handlers.AllFufufafa)
	api.Get("/:quote_id<int>?", cache.New(cache.Config{Expiration: 30 * time.Minute, CacheControl: true}), handlers.AllFufufafa)
	api.Get("/random", handlers.RandomQuote)

}
