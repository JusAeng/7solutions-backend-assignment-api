package services

import (
	"testing"
)

func TestFetchDataFromInvalidApi(t *testing.T){
	url := "fghijk"
	_,err := fetchData(url)
	if err == nil {
		t.Error("Should be no response")
	}

}
func TestRegexStringFromInvalidApi(t *testing.T) {
	url := "abcde"
	_,err := RegexStringFromApi(url,`[\w-]+`)
	if err == nil {
		t.Error("Should be no response")
	}
}

func TestCountWordsEqualToRegexWords(t *testing.T) {
	url := "https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text"
	matches,err := RegexStringFromApi(url,`[\w-]+`)
	if err != nil {
		t.Error("Bad gateway")
	}
	result,err := CountWords(matches)
	if err != nil {
		t.Error("Bad gateway")
	}

	get := 0
    for _, count := range result {
        get += count
    }

	want := len(matches)
	if want != get{
		t.Errorf("want result %d words but get %d words",want,get)
	}
}