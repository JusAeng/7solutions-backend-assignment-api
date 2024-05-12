package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"
)

func main() {
	http.HandleFunc("/beef/summary", beefSummaryHandler)
	if err := http.ListenAndServe(":8000",nil); err != nil {
		log.Fatal(err)
	}
	
}

func beefSummaryHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/beef/summary" {
		http.Error(w, "404 Not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	data := countWords("https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text")
	resp := struct {
		Beef map[string]interface{} `json:"beef"`
	}{
		Beef: data,
	}

	json.NewEncoder(w).Encode(&resp)
}

func countWords(url string) map[string]interface{} {
	text,_ := fetchData(url)

	wordRegex := regexp.MustCompile(`[\w-]+`)
	matches := wordRegex.FindAllString(text, -1)

	wordCount := make(map[string]int)
	for _, word := range matches {
		word = strings.ToLower(word)
		wordCount[word]++
	}

	res := map[string]interface{}{
		"beef": wordCount,
	}

	// jsonData, err := json.Marshal(res)
	// if err != nil {
	// 	fmt.Println("Error marshaling JSON:", err)
	// 	return
	// }

    return res
}

func fetchData(url string) (string, error) {
    response, err := http.Get(url)
    if err != nil {
        return "", err
    }
    defer response.Body.Close()

    body, err := io.ReadAll(response.Body)
    if err != nil {
        return "", err
    }

    return string(body), nil
}