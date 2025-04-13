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
	"github.com/gofiber/fiber/v2/middleware/favicon"
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
	app := fiber.New()
	app.Use(logger.New())
	// favicon
	app.Use(favicon.New(favicon.Config{
		File: "./public/favicon/favicon.ico",
		URL:  "/favicon.ico",
	}))
	router.SetupRouter(app)
	app.Use(middleware.NotFound)

	return adaptor.FiberApp(app)
}
