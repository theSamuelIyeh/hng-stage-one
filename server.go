package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	_ "github.com/joho/godotenv/autoload"
	"os"
)

func setup() {
	app := fiber.New()
	router(app)
	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
