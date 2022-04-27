package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mythosmystery/typenotes-go-server/controllers"
	"github.com/mythosmystery/typenotes-go-server/middleware"
	"github.com/mythosmystery/typenotes-go-server/views"
)

func InitRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	api := app.Group("/api", middleware.Authorized)

	api.Get("/notes", controllers.GetNotes)
	api.Get("/notes/:id", controllers.GetNote)
	api.Post("/notes", controllers.CreateNote)
	api.Put("/notes/:id", controllers.UpdateNote)
	api.Delete("/notes/:id", controllers.DeleteNote)
	api.Get("/users", controllers.GetUsers)
	api.Get("/me", controllers.GetMe)

	app.Post("/login", controllers.LoginUser)
	app.Post("/register", controllers.RegisterUser)
	app.Post("/forgot", controllers.ForgotPassword)
	app.Post("/reset", controllers.ResetPassword)
	app.Get("/reset", views.Reset)
}
