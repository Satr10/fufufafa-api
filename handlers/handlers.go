package handlers

import (
	"github.com/Satr10/fufufafa-api/database"
	"github.com/Satr10/fufufafa-api/model"
	"github.com/gofiber/fiber/v2"
)

func AllFufufafa(c *fiber.Ctx) error {
	var allFufufafa []model.Post
	database.DB.Find(&allFufufafa)
	return c.JSON(allFufufafa)

}
