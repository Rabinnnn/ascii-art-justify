package functions

import (
	"fmt"
    "runtime"
  //  "flag"
//	"os"
	"os/exec"
	"strings"
    "strconv"
)


func GetTerminalWidth() int {
    if runtime.GOOS == "windows" {
        cmd := exec.Command("cmd", "/c", "mode con")
        out, err := cmd.Output()
        if err != nil {
            fmt.Println("Error:", err)
            return 80 // Default width if detection fails
        }
        output := string(out)
        for _, line := range strings.Split(output, "\n") {
            if strings.Contains(line, "Columns:") {
                parts := strings.Split(line, ":")
                if len(parts) > 1 {
                    width, err := strconv.Atoi(strings.TrimSpace(parts[1]))
                    if err == nil {
                        return width
                    }
                }
            }
        }
    } else {
        cmd := exec.Command("tput", "cols")
        out, err := cmd.Output()
        if err != nil {
            fmt.Println("Error:", err)
            return 80 // Default width if detection fails
        }
        var width int
        fmt.Sscanf(string(out), "%d", &width)
        return width
    }
    return 80 // Default width if detection fails
}

func AlignText(text, align string, width int) string {
    switch align {
    case "left":
        return fmt.Sprintf("%-*s", width, text)
    case "right":
        return fmt.Sprintf("%*s", width, text)
    case "center":
        return fmt.Sprintf("%*s", (width+len(text))/2, text)
    case "justify":
        words := strings.Fields(text)
        if len(words) == 1 {
            return fmt.Sprintf("%*s", width, text)
        }
        spaces := width - len(text) + len(words) - 1
        spaceBetweenWords := spaces / (len(words) - 1)
        extraSpaces := spaces % (len(words) - 1)
        var result strings.Builder
        for i, word := range words {
            if i > 0 {
                result.WriteString(strings.Repeat(" ", spaceBetweenWords))
                if extraSpaces > 0 {
                    result.WriteString(" ")
                    extraSpaces--
                }
            }
            result.WriteString(word)
        }
        return result.String()
    default:
        return "Usage: go run . [OPTION] [STRING] [BANNER]"
    }
}



/*
func applyAlignment(line, align string, width int) string {
	switch align {
	case "center":
		padding := (width - len(line)) / 2
		return strings.Repeat(" ", padding) + line
	case "right":
		padding := width - len(line)
		return strings.Repeat(" ", padding) + line
	case "justify":
		words := strings.Fields(line)
		if len(words) == 1 {
			return line
		}
		spaces := width - len(strings.Join(words, ""))
		spaceBetween := spaces / (len(words) - 1)
		extraSpaces := spaces % (len(words) - 1)
		var result strings.Builder
		for i, word := range words {
			if i > 0 {
				result.WriteString(strings.Repeat(" ", spaceBetween))
				if extraSpaces > 0 {
					result.WriteString(" ")
					extraSpaces--
				}
			}
			result.WriteString(word)
		}
		return result.String()
	case "left":
		fallthrough
	default:
		return line
	}
} */


/*
spaceCount := 0
// Iterate over each line of the input.
for _, word := range Input {
    if word == "" {
        spaceCount++
        if spaceCount < len(Input) {
            fmt.Println()
        }
    } else {
        // Print the banner for non-empty strings with the specified alignment.
        Ascii.PrintBanner(word, *alignment)
    }
} */