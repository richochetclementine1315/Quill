package routes

import (
	"QuillBackend/controller"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controller.Register)
	app.Post("/api/login", controller.Login)
	// Any API below this line will be protected and require authentication**
	// app.Use(middleware.isAuthenticate)
}
