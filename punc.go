package goReloaded

import (
	"strings"
  "regexp"
)

func Punctuation(punct string, sentence string) string {
	// Remove spaces before punctuation using regular expression
	sentence = regexp.MustCompile(`\s*` + regexp.QuoteMeta(punct)).ReplaceAllString(sentence, punct)

	// Check if there is no space after the punctuation, and add one if needed
	sentence = strings.Replace(sentence, punct, punct+" ", -1)

	// Check if there are more than one space after punctuation and replace with a single space
	sentence = strings.Replace(sentence, punct+"  ", punct+" ", -1)

	// Remove multiple consecutive spaces
	sentence = strings.Join(strings.Fields(sentence), " ")

	return sentence
}

func Comma(sentence string) string {
  return Punctuation(",", sentence)
}

func Colon(sentence string) string {
	return Punctuation(":", sentence)
}
  
  func Semicolon(sentence string) string {
	return Punctuation( ";" ,sentence)  
  }




func QuestionExclamation(sentence string) string {
	re := regexp.MustCompile(`\s+([?!])`)
	sentence = re.ReplaceAllString(sentence, "$1")

	re = regexp.MustCompile(`([?!])\s*([?!])`)
	sentence = re.ReplaceAllString(sentence, "$1$2")

	re = regexp.MustCompile(`([?!])\s{2,}`)
	sentence = re.ReplaceAllString(sentence, "$1 ")

	return sentence
}

func Dots(input string) string {
    // Remove spaces before dots or consecutive dots 
    reg := regexp.MustCompile(`\s*\.\s*`)
    output := reg.ReplaceAllString(input, ".")

    // Iterate over the string character by character
    for i := 0; i < len(output); i++ {
        if output[i] == '.' {
            // If the next character is not a dot, add a space
            if i+1 < len(output) && output[i+1] != '.' {
                output = output[:i+1] + " " + output[i+1:]  
            }    
        }  
    }
    
    // Remove any extra spaces that may have been added after the last dot
    output = strings.TrimSuffix(output, " ")
    
    return output 
}

func Quotes(sentence string) string {
	re := regexp.MustCompile(`'\s*([^']+?)\s*'`) // Updated regex with '?' for non-greedy match
	fixedSentence := re.ReplaceAllString(sentence, "'$1'")
	return fixedSentence
}
