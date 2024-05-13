package handlers

import (
	"github.com/JusAeng/7solutions-backend-assignment-api/services"
	"github.com/gofiber/fiber/v2"
)

func BeefSummaryHandler(c *fiber.Ctx) error {
	url := "https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text"
	matches,err := services.RegexStringFromApi(url,`[\w-]+`)
	if err != nil {
		return c.SendStatus(fiber.StatusBadGateway)
	}

	summary,err := services.CountWords(matches)
	if err != nil {
		return c.SendStatus(fiber.StatusBadGateway)
	}
	resp := struct {
		Beef map[string]int `json:"beef"`
	}{
		Beef: summary,
	}

	return c.JSON(resp)
}