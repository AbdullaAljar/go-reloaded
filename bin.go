package goReloaded

import (
	"fmt"
	"log"
	"strconv"
	"regexp"
)


func Bin(sentence string) string {

    // search for pattern (bin) using regexp.

    re := regexp.MustCompile(`(\b[01]+\b) \(bin\)`)

    // for each match, use strconv.ParseInt to convert the bin to int.

    output := re.ReplaceAllStringFunc(sentence, func(match string) string {
        binString := re.FindStringSubmatch(match)[1]
        binNumber, err := strconv.ParseInt(binString, 2, 64)

    // if error, return original
    // If the conversion is successful, it replaces using the fmt.Sprintf.

        if err != nil {
            log.Printf("Error converting bin to decimal: %v\n", err)
            return match
        }
        return fmt.Sprintf("%d", binNumber)
    })
    return output
}
