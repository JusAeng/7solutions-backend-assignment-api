package handlers

import (
	"github.com/JusAeng/7solutions-backend-assignment-api/services"
	"github.com/gofiber/fiber/v2"
)

func BeefSummaryHandler(c *fiber.Ctx) error {
	data := services.CountWords("https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text")
	resp := struct {
		Beef map[string]int `json:"beef"`
	}{
		Beef: data,
	}

	return c.JSON(resp)
}