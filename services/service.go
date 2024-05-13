package services

import (
	"errors"
	"regexp"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func CountWords(matches []string) (map[string]int, error) {
	wordCount := make(map[string]int)
	for _, word := range matches {
		word = strings.ToLower(word)
		wordCount[word]++
	}

	return wordCount,nil
}

func RegexStringFromApi(url string,pattern string) ([]string,error) {
	text, err := fetchData(url)
	if err !=  nil {
		return nil,err
	}
	wordRegex := regexp.MustCompile(pattern)
	matches := wordRegex.FindAllString(text, -1)
	return matches,nil
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