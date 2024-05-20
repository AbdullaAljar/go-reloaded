package goReloaded

import (
	"strings"
)

func Vowels(sentence string) string {
	words := strings.Fields(sentence)
	vowels := "aeiouAEIOU"
	correctedWords := make([]string, len(words))

	for i, word := range words {
		// Check if the word is "a" or "A"
		if word == "a" || word == "A" {
			// Check if the word after the current one starts with a vowel
			if i < len(words)-1 {
				nextWord := words[i+1]
				if len(nextWord) > 0 && strings.Contains(vowels, string(nextWord[0])) {
					// Change "a" or "A" to "an" or "An" if the next word starts with a vowel
					if word == "A" {
						correctedWords[i] = "An"
					} else {
						correctedWords[i] = "an"
					}
					continue
				}
			}
		}
		// Use the original word if it doesn't meet the condition
		correctedWords[i] = word
	}

	return strings.Join(correctedWords, " ")
}
