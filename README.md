# ASCII Art Generator

## Description
This is a simple command-line tool written in Go that generates ASCII art based on standard text. It reads characters from a standard text file and converts them into ASCII art representation.

## Usage
1. **Clone Repository**: Clone this repository to your local machine.
2. **Build Program**: Ensure you have Go installed on your machine. Navigate to the directory containing the `main.go` file and run the following command to build the program:
    ```bash
    go build main.go
    ```
3. **Run Program**: Execute the program by providing the input text as a command-line argument. For example:
    ```bash
    go run . "Hello"
    ```
    This will generate ASCII art representing the word "Hello".
4. **Line Breaks**: You can provide a newline character (`\n`) to create a line break in the ASCII art. For example:
    ```bash
    go run . "Hello\nWorld"
    ```
    This will generate ASCII art for "Hello" followed by a newline, and then ASCII art for "World".

## Arguments
- **Input Text**: The program takes one argument, which is the input text for which ASCII art is to be generated. This text can contain alphanumeric characters, special characters, and newline characters (`\n`).

## Files
- **main.go**: Contains the main program logic, including argument parsing and reading from the standard text file.
- **ascii_art.go**: Contains the function for generating ASCII art based on the input text.

## Standard Text File
The ASCII art characters are stored in a standard text file named `standard.txt`. Each character's representation occupies 8 lines in the file.

## Errors
- Only one argument is allowed to be input by the user
- If too many or too few arguments are provided, the program will display an error message and exit.
- If there is an error reading the `standard.txt` file, the program will display an error message and exit.

## Dependencies
- This program does not have any external dependencies.

