package services

import (
	"errors"
	"regexp"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func CountWords(url string) map[string]int {
	text, _ := fetchData(url)
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
		return "", errs[0]
	}
	if statusCode != 200 {
		return "", errors.New("status code != 200")
	}

	return string(body), nil
}