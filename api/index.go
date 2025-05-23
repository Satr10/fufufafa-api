package api

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/Satr10/fufufafa-api/database"
	"github.com/Satr10/fufufafa-api/helpers"
	"github.com/Satr10/fufufafa-api/middleware"
	"github.com/Satr10/fufufafa-api/model"
	"github.com/Satr10/fufufafa-api/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// Handler is the main entry point of the application. Think of it like the main() method
func Handler(w http.ResponseWriter, r *http.Request) {
	database.ConnectDB()
	// menentukan banyak quotes
	database.DB.Model(&model.Post{}).Count(&database.TotalQuote)

	helpers.Random = rand.New(rand.NewSource(time.Now().UnixMicro()))

	// This is needed to set the proper request path in `*fiber.Ctx`
	r.RequestURI = r.URL.String()

	handler().ServeHTTP(w, r)
}

// building the fiber application
func handler() http.HandlerFunc {
	// THIS CAUSES ERROR IN VERCEL
	// engine := django.New("/views", ".html")
	app := fiber.New(fiber.Config{
		// Views: engine,
	})
	app.Use(logger.New())
	app.Use(cors.New())

	router.SetupRouter(app)
	app.Use(middleware.NotFound)

	return adaptor.FiberApp(app)
}
