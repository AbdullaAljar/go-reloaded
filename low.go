package goReloaded

import (
	"strconv"
	"regexp"
	"strings"
)

func Low(sentence string) string {
    re := regexp.MustCompile(`((?:\b\w+\b\s+){0,100})\(low(?:,\s*(\d+))?\)`)

    output := re.ReplaceAllStringFunc(sentence, func(match string) string {
        groups := re.FindStringSubmatch(match)
        wordsToConvert := 1 // default to converting one word
        if len(groups) > 2 && groups[2] != "" {
            n, err := strconv.Atoi(groups[2])
            if err == nil {
                wordsToConvert = n
            }
        }
        words := strings.Fields(strings.TrimSpace(groups[1]))
        if len(words) < wordsToConvert {
            return match // not enough words to convert
        }
        start := len(words) - wordsToConvert
        if start < 0 {
            start = 0
        }
        for i := start; i < len(words); i++ {
            words[i] = strings.ToLower(words[i])
        }
        return strings.Join(words, " ")
    })

    return output
}