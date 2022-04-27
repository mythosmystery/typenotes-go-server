package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mythosmystery/typenotes-go-server/config"
	"github.com/mythosmystery/typenotes-go-server/middleware"
	"github.com/mythosmystery/typenotes-go-server/models"
)

func GetNotes(c *fiber.Ctx) error {
	var notes []models.Note
	config.DB.Joins("User").Find(&notes)
	return c.JSON(notes)
}

func GetNote(c *fiber.Ctx) error {
	id := c.Params("id")
	var note models.Note
	config.DB.Joins("User").First(&note, id)
	return c.JSON(note)
}

func CreateNote(c *fiber.Ctx) error {
	user := c.Locals("user").(middleware.AuthUser)
	var note models.Note
	note.UserId = user.ID
	if err := c.BodyParser(&note); err != nil {
		return err
	}
	config.DB.Create(&note)
	config.DB.Joins("User").First(&note, note.ID)
	return c.JSON(note)
}

func UpdateNote(c *fiber.Ctx) error {
	id := c.Params("id")
	var note models.Note
	if err := c.BodyParser(&note); err != nil {
		return err
	}
	config.DB.Model(&note).Where("id = ?", id).Updates(note)
	config.DB.Joins("User").First(&note, id)
	return c.JSON(note)
}

func DeleteNote(c *fiber.Ctx) error {
	id := c.Params("id")
	var note models.Note
	config.DB.First(&note, id)
	config.DB.Delete(&note)
	return c.JSON(note)
}