package handlers

import (
	"strconv"

	"github.com/Satr10/fufufafa-api/database"
	"github.com/Satr10/fufufafa-api/helpers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

// AllFufufafa returns all Fufufafa quotes from the database.
// @Summary Get all Fufufafa quotes
// @Description Returns a list of all Fufufafa quotes from the database
// @Tags quotes
// @Accept json
// @Produce json
// @Success 200 {array} model.Post
// @Router /api [get]
func AllFufufafa(c *fiber.Ctx) error {
	c.Set("Cache-Control", "public, max-age=60, s-maxage=60, stale-while-revalidate=30")
	quotes := helpers.AllQuote()
	return c.JSON(quotes)
}

// SingleFufufafa returns a specific Fufufafa quote from the database by its ID.
// @Summary Get a specific Fufufafa quote
// @Description Returns a specific Fufufafa quote from the database using the provided quote ID
// @Tags quotes
// @Accept json
// @Produce json
// @Param quote_id path int true "Quote ID"
// @Success 200 {object} model.Post
// @Failure 400 {object} object "Bad Request"
// @Failure 404 {object} object "Not Found"
// @Router /api/{quote_id} [get]

func SingleFufufafa(c *fiber.Ctx) error {
	quoteID := c.Params("quote_id")

	id, err := strconv.Atoi(quoteID)
	if err != nil {
		log.Error("Error converting string to int: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid quote ID format",
		})
	}

	quote := helpers.QuoteById(id)
	if quote.ID == 0 {
		return c.JSON(fiber.Map{
			"error":   true,
			"message": "Data tidak ditemukan",
		})
	}
	return c.JSON(quote)
}

// RandomQuote returns a random Fufufafa quote from the database.
// @Summary Get a random Fufufafa quote
// @Description Returns a random Fufufafa quote from the database
// @Tags quotes
// @Accept json
// @Produce json
// @Success 200 {object} model.Post
// @Router /api/random [get]
func RandomQuote(c *fiber.Ctx) error {
	pilihan := helpers.Random.Intn(int(database.TotalQuote)) + 1

	quoteRandom := helpers.QuoteById(pilihan)
	return c.JSON(quoteRandom)
}

// Index godoc
// @Summary Index endpoint
// @Description Returns a welcome message with link to API documentation
// @Tags root
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string "Returns status, welcome message and docs link"
// @Router / [get]
func Index(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  "Success",
		"message": "Halo User",
		"docs":    "/api/docs",
	})
}
