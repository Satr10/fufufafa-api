package handlers

import (
	"strconv"

	"github.com/Satr10/fufufafa-api/database"
	"github.com/Satr10/fufufafa-api/helpers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func AllFufufafa(c *fiber.Ctx) error {
	c.Set("Cache-Control", "public, max-age=60, s-maxage=60, stale-while-revalidate=30")
	quotes := helpers.AllQuote()
	return c.JSON(quotes)
}

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

func RandomQuote(c *fiber.Ctx) error {
	pilihan := helpers.Random.Intn(int(database.TotalQuote)) + 1

	quoteRandom := helpers.QuoteById(pilihan)
	return c.JSON(quoteRandom)
}

func Index(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  "Success",
		"message": "Halo User",
	})
}
