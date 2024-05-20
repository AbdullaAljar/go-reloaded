package goReloaded

import (
	"strconv"
	"regexp"
	"strings"
)

func Up(sentence string) string {
    re := regexp.MustCompile(`((?:\b\w+\b\s+){0,100})\(up(?:,\s*(\d+))?\)`)

    output := re.ReplaceAllStringFunc(sentence, func(match string) string {
        groups := re.FindStringSubmatch(match)
        wordsToCapitalize := 1 // default to capitalizing one word
        if len(groups) > 2 && groups[2] != "" {
            n, err := strconv.Atoi(groups[2])
            if err == nil {
                wordsToCapitalize = n
            }
        }
        words := strings.Fields(strings.TrimSpace(groups[1]))
        if len(words) < wordsToCapitalize {
            return match // not enough words to capitalize
        }
        start := len(words) - wordsToCapitalize
        if start < 0 {
            start = 0
        }
        for i := start; i < len(words); i++ {
            words[i] = strings.ToUpper(words[i])
        }
        return strings.Join(words, " ")
    })

    return output
}