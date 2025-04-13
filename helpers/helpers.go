package helpers

import (
	"math/rand"

	"github.com/Satr10/fufufafa-api/database"
	"github.com/Satr10/fufufafa-api/model"
)

var Random *rand.Rand

// mendapatkan semua quote dari db
func AllQuote() (quotes []model.Post) {
	database.DB.Find(&quotes)
	return quotes
}
func QuoteById(id int) (quote model.Post) {
	database.DB.First(&quote, id)
	return quote
}
