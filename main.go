package main

import (
	"github.com/JusAeng/7solutions-backend-assignment-api/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/beef/summary",handlers.BeefSummaryHandler)
	app.Listen(":8000")
}

