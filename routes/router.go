package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mythosmystery/typenotes-go-server/controllers"
)

func InitRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/notes", controllers.GetNotes)
	app.Get("/notes/:id", controllers.GetNote)
	app.Post("/notes", controllers.CreateNote)
	app.Put("/notes/:id", controllers.UpdateNote)
	app.Delete("/notes/:id", controllers.DeleteNote)

	app.Get("/users", controllers.GetUsers)
	app.Post("/login", controllers.LoginUser)
	app.Post("/register", controllers.RegisterUser)
}