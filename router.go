package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thesamueliyeh/hng-stage-1/controllers"
)

func router(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/", controllers.MainController)
}
