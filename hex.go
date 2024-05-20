package goReloaded

import (
	"fmt"
	"log"
	"strconv"
	"regexp"
)

func Hex(sentence string) string {

    // find pattern (hex)

    re := regexp.MustCompile(`(\b[0-9A-Fa-f]+\b) \(hex\)`)

    // replace matches with hex using strconv.ParseInt

    output := re.ReplaceAllStringFunc(sentence, func(match string) string {
        hexString := re.FindStringSubmatch(match)[1]
        hexNumber, err := strconv.ParseInt(hexString, 16, 64)

    // if error return match and statement otherwise return replacement

        if err != nil {
            log.Printf("Error converting hex to decimal: %v\n", err)
            return match
        }
        return fmt.Sprintf("%d", hexNumber)
    })
    return output
}