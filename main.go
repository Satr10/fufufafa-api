package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/Satr10/fufufafa-api/database"
	"github.com/Satr10/fufufafa-api/helpers"
	"github.com/Satr10/fufufafa-api/middleware"
	"github.com/Satr10/fufufafa-api/model"
	"github.com/Satr10/fufufafa-api/router"
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	database.ConnectDB()

	// menentukan banyak quotes
	database.DB.Model(&model.Post{}).Count(&database.TotalQuote)
	app := fiber.New()

	// init random
	helpers.Random = rand.New(rand.NewSource(time.Now().UnixMicro()))

	// Initialize log
	app.Use(logger.New())

	// gunakan swagger
	swaggerCfg := swagger.Config{
		BasePath: "/api/",
		FilePath: "./docs/swagger.json",
		Path:     "docs",
	}
	app.Use(swagger.New(swaggerCfg))

	// favicon
	app.Use(favicon.New(favicon.Config{
		File: "./public/favicon/favicon.ico",
		URL:  "/favicon.ico",
	}))

	// route
	router.SetupRouter(app)

	// 404 page
	app.Use(middleware.NotFound)

	log.Fatal(app.Listen(":5001"))
}
