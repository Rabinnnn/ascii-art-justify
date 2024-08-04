package main

import (
	"ascii-art-justify/functions"
	"flag"
	"fmt"
	"os"
	"runtime"
	"os/exec"
	//"os/signal"
	"strings"
	//"syscall"
    "time"
	 
)

func main() {
	// Define flag that will be used to specify the output file
    alignFlag := flag.String("align", "left", "Alignment type: center, left, right, justify")
	flag.Parse()

	if len(flag.Args()) > 2 || len(flag.Args()) < 1 {
		fmt.Println("Usage1: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --align=right something standard")
		return
	}

	stringInput := flag.Args()[0]

	// Variable to track if the flag was set
	var nameSet bool
	var flagSet bool 

	// Enforce the flag format to be used to be --output=<filename.txt>
	flag.Visit(func(f *flag.Flag) {
		if f.Name == "align" {
			flagSet = true
			result := strings.Replace(os.Args[1], *alignFlag, "", 1)
			if (result == "--align=") {
				nameSet = true
			}
		}
	})
	// defining usage and error handling
	if !nameSet && len(flag.Args()) == 2 || !flagSet && len(flag.Args()) == 2{
		fmt.Println("Usage2: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --align=right something standard")
		return
	}

	// if !(strings.HasSuffix(*alignFlag, ".txt")) {
	// 	fmt.Println("Usage3: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --align=right something standard")
	// 	return
	// }

	BannerFile := "standard.txt"

	if len(flag.Args()) == 2 {
		banner := strings.Replace(flag.Args()[1], ".txt", "", 1)
		BannerFile = banner + ".txt"
	}

	// read banner file specified
	file, err := os.ReadFile(BannerFile)
	if err != nil {
		fmt.Println("Error openning", BannerFile, err)
		return
	}
	file = []byte(strings.Replace(string(file), "\r\n", "\n", -1))

	fileLine := strings.Split(string(file), "\n")

	link := ""
	switch BannerFile {
	case "standard.txt":
		link = "https://learn.zone01kisumu.ke/git/root/public/src/branch/master/subjects/ascii-art/standard.txt"
	case "shadow.txt":
		link = "https://learn.zone01kisumu.ke/git/root/public/src/branch/master/subjects/ascii-art/shadow.txt"
	case "thinkertoy.txt":
		link = "https://learn.zone01kisumu.ke/git/root/public/src/branch/master/subjects/ascii-art/thinkertoy.txt"
	}

	if len(fileLine) != 856 {
		fmt.Println("The file", BannerFile, "is not correctly formated, please use the correct version", link, "!!!")
		return
	}

	//Write the results to the output file specified by user then print the results.
	// asciiOutput := functions.AsciiArt(stringInput, fileLine)
	// error := os.WriteFile(*alignFlag, []byte(asciiOutput), 0644)
	// if error != nil {
	// 	fmt.Println("Error:", error)
	// } else if flagSet == false {
	// 	fmt.Print(functions.AsciiArt(stringInput, fileLine))
	// }

	prevWidth := functions.GetTerminalWidth()
	go func() {
        for {
            time.Sleep(500 * time.Millisecond) // Adjust the sleep duration as needed
            currWidth := functions.GetTerminalWidth()
            if currWidth != prevWidth {
                prevWidth = currWidth
              //  fmt.Print("\033[H\033[2J") // Clear terminal
			  clearScreen()
			  if flagSet {
					//width := functions.GetTerminalWidth()
					//alignedText := functions.AlignText(stringInput, *alignFlag, width)
					//fmt.Println(alignedText)
					fmt.Print(functions.AsciiArt(stringInput, fileLine, *alignFlag))
				}else {
					fmt.Print(functions.AsciiArt(stringInput, fileLine, "left"))
				}
            }
        }
    }()
	 // Initial display
	 clearScreen()
	// fmt.Print("\033[H\033[2J") // Clear terminal

	if flagSet {
		//width := functions.GetTerminalWidth()
		//alignedText := functions.AlignText(stringInput, *alignFlag, width)
		//fmt.Println(alignedText)
		fmt.Print(functions.AsciiArt(stringInput, fileLine, *alignFlag))
	}else {
		fmt.Print(functions.AsciiArt(stringInput, fileLine, "left"))
	}

	// Keep the program running to capture resize events
    select {}
}

func clearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		fmt.Print("\033[H\033[2J") // Clear screen for Unix-like systems
	}
}