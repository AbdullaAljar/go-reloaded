package goReloaded

import (
	"fmt"
	"regexp"
)

func CheckForWords(sentence string) string {
	targetWords := []string{"\\(bin\\)", "\\(hex\\)", "\\(up, ", "\\(low, ", "\\(cap, ", "\\(up\\)", "\\(low\\)", "\\(cap\\)"}
	for _, word := range targetWords {
		re := regexp.MustCompile(word)
		if re.MatchString(sentence) {
			fmt.Printf("Invalid Sentence: error regarding '%s'!\n", word)
			return "INVALID SENTENCE"
		}
		
	}
	return sentence
}

