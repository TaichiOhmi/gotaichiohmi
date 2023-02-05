package handlers

import (
	"net/http"

	"github.com/TaichiOhmi/gotaichiohmi/database"
	"github.com/TaichiOhmi/gotaichiohmi/models"
	"github.com/gofiber/fiber/v2"
)

func ListFacts(c *fiber.Ctx) error {
	facts := []models.Fact{}
	database.DB.DB.Find(&facts)
	return c.Status(http.StatusOK).JSON(facts)
}

func CreateFact(c *fiber.Ctx) error {
	fact := new(models.Fact)
	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	database.DB.DB.Create(&fact)
	return c.Status(http.StatusOK).JSON(fact)
}
