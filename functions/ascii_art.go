package functions

import (
	// "strconv"
	"fmt"
	"strings"
)

func AsciiArt(input string, fileLine []string, alignment string) string {
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
	terminalWidth := GetTerminalWidth()
	for _, word := range words {
		if word == "" {
			result += "\n"
		} else {
			lines := make([]string, 8)
			for i := 0; i < 8; i++ {
				for j := 0; j < len(word); j++ {
					start := (int(word[j]-' ') * 9) + 1 // calculating the begining of a character based on data from standard.txt
					// result += strconv.Itoa(i)
					//result += fileLine[start+i]
					lines[i] += fileLine[start+i]
				}
				
			}

			for i := 0; i < 8; i++ {
                switch alignment {
				case "right":
                    // padding := terminalWidth - len(lines[i])
                    // if padding > 0 {
                    //     result += strings.Repeat(" ", padding)
                    // }
					var result strings.Builder
					for _, line := range lines {
						padding := terminalWidth - len(line)
						if padding < 0 {
							padding = 0
						}
						result.WriteString(strings.Repeat(" ", padding) + line)
					}
					return result.String()
                case "center":
                    // padding := (terminalWidth - len(lines[i])) / 2
                    // if padding > 0 {
                    //     result += strings.Repeat(" ", padding)
                    // }
					var result strings.Builder
					for _, line := range lines {
						padding := (terminalWidth - len(line)) / 2
						if padding < 0 {
							padding = 0
						}
						result.WriteString(strings.Repeat(" ", padding) + line + "\n")
					}
					return result.String()
                case "justify":
                    wordsInLine := strings.Fields(lines[i])
                    if len(wordsInLine) > 1 {
                        spacesNeeded := terminalWidth - len(strings.Join(wordsInLine, ""))
                        spaceBetweenWords := spacesNeeded / (len(wordsInLine) - 1)
                        extraSpaces := spacesNeeded % (len(wordsInLine) - 1)
                        justifiedLine := strings.Builder{}
                        for j, word := range wordsInLine {
                            if j > 0 {
                                justifiedLine.WriteString(strings.Repeat(" ", spaceBetweenWords))
                                if extraSpaces > 0 {
                                    justifiedLine.WriteString(" ")
                                    extraSpaces--
                                }
                            }
                            justifiedLine.WriteString(word)
                        }
                        lines[i] = justifiedLine.String()
                    }
                case "left":
                    // No additional action needed for left alignment
                }
                result += lines[i] + "\n"
            }
        }
    }
	return result
}

// func AlignRight(text string, width int) string {
// 	var result strings.Builder
// 	lines := strings.Split(text, "\n")
// 	for _, line := range lines {
// 		result.WriteString(fmt.Sprintf("%*s\n", width, line))
// 	}
// 	return result.String()
// }

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
