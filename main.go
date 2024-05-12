package main

import (
	// "encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()
	app.Get("/beef/summary",beefSummaryHandler)
	app.Listen(":8000")
	
}

func beefSummaryHandler(c *fiber.Ctx) error {
	data := countWords("https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text")
	resp := struct {
		Beef map[string]int `json:"beef"`
	}{
		Beef: data,
	}

	return c.JSON(resp)
}

func countWords(url string) map[string]int {
	text,_ := fetchData(url)
	fmt.Print(text)
	wordRegex := regexp.MustCompile(`[\w-]+`)
	matches := wordRegex.FindAllString(text, -1)

	wordCount := make(map[string]int)
	for _, word := range matches {
		word = strings.ToLower(word)
		wordCount[word]++
	}

    return wordCount
}

func fetchData(url string) (string, error) {
	agent := fiber.Get(url)
	statusCode, body, errs := agent.Bytes()
	if len(errs) > 0 {
		return "",errs[0]
	}
	if statusCode != 200 {
		return "",errors.New("status code != 200")
	}

	return string(body),nil
}