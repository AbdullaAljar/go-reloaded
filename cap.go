package goReloaded

import (
	"strconv"
	"regexp"
	"strings"
)

func Cap(sentence string) string {
	re := regexp.MustCompile(`((?:\b\w+\b\s+){0,100})\(cap(?:,\s*(\d+))?\)`)

	output := re.ReplaceAllStringFunc(sentence, func(match string) string {
		groups := re.FindStringSubmatch(match)
		wordsToCapitalize := 1 // default to capitalizing the first word
		if len(groups) > 2 && groups[2] != "" {
			n, err := strconv.Atoi(groups[2])
			if err != nil {
				
				return match // Return the original match if there's an error converting the number
			}
			wordsToCapitalize = n
		}
		words := strings.Fields(strings.TrimSpace(groups[1]))
		if len(words) < wordsToCapitalize {
			
			return match // Return the original match if there are not enough words
		}
		start := len(words) - wordsToCapitalize
		if start < 0 {
			start = 0
		}
		for i := start; i < len(words); i++ {
			words[i] = strings.Title(strings.ToLower(words[i]))
		}
		return strings.Join(words, " ")
	})

	return output
}