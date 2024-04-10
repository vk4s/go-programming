package main

import (
	"fmt"
	"io"
	"os"
)

func readFileContents(filename string) (string, error) {
	// Open the file to
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}

	// Read the contents of the file
	fileContents, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(fileContents), nil
}

func main() {
	// Parse the command line arguments and extract filename (path)
	// Second argument is the filename, First argument is the name of the program
	filename := os.Args[1]
	fmt.Println(filename)

	// (1) Read the contents of the file
	fileContents, err := readFileContents(filename)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// (2) Write the contents to terminal
	fmt.Println(fileContents)

	// Combine (1) and (2) and convert into File Interface statement
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Write the contents of the `file` to the stdout (terminal)
	io.Copy(os.Stdout, file)
}
