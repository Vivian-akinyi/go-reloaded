package main

import (
	"fmt"
	"os"
)

func main() {
	// Check if the correct number of arguments are provided
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . input.txt output.txt")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	// Read the input file
	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Convert the file contents to a string
	text := string(data)

	// Process the text
	result := ProcessText(text)

	// Write the processed text to the output file
	err = os.WriteFile(outputFile, []byte(result), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Println("Processing completed successfully.")
}
