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

func TestSummaryWordsEqualToRegexWords(t *testing.T) {
	url := "https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text"
	matches,err := RegexStringFromApi(url,`[\w-]+`)
	if err != nil {
		t.Error("Bad gateway")
	}

	testCases := []struct {
		name string
		matches []string
		expected int
	}{
		{"Mock Data",[]string{
			"Magna",
			"short",
			"loin",
			"est",
			"anim",
			"doner",
			"consequat",
			"Cillum",
			"magna",
			"ut",
			"shankle",
			"excepteur",
			"qui",
			"quis",
			"do",
			"bacon",
			"duis",
			"cupim",
			"turkey",
			"in",
			"anim",
			"Magna",
		},22},
		{"Real Data",matches, len(matches)},
	}

	for _,tc := range testCases{
		t.Run(tc.name, func(t *testing.T){
			result,err := CountWords(tc.matches)
			if err != nil {
				t.Error("Bad gateway")
			}
			get := 0
			for _, count := range result {
				get += count
			}
			expected := tc.expected
			if get != expected {
				t.Errorf("Error on %s that expect to %d but get %d",tc.name,expected,get)
			}
		})
	}
}