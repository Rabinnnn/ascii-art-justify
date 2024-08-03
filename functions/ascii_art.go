package functions

import (
	// "strconv"
	"fmt"
	"strings"
)

func AsciiArt(input string, fileLine []string) string {
	result := ""

	// replacing every instance of new line with the symbol \\n
	input = strings.Replace(input, "\n", "\\n", -1)

	if !validSentence(input) {
		return ""
	}

	// slicing the input base on the presence of the string "\n"
	words := strings.Split(input, "\\n")

	empty := EmptyArray(words)
	if empty != "false" {
		fmt.Print(empty)
		return ""
	}

	for _, word := range words {
		if word == "" {
			result += "\n"
		} else {
			for i := 0; i < 8; i++ {
				for j := 0; j < len(word); j++ {
					start := (int(word[j]-' ') * 9) + 1 // calculating the begining of a character based on data from standard.txt
					// result += strconv.Itoa(i)
					result += fileLine[start+i]
				}
				result += "\n"
			}
		}
	}
	return result
}

func validSentence(word string) bool {
	for _, letter := range word {
		if !(letter >= ' ' && letter <= '~') {
			fmt.Println("Error, character", string(letter), "is an invalid character!!!!")
			return false
		}
	}
	return true
}

func EmptyArray(words []string) string {
	result := ""

	for i, word := range words {
		if word != "" {
			result = "false"
			return result
		}
		if i != 0 {
			result += "\n"
		}
	}
	return result
}
