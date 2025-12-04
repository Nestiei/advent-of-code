package main

import (
	"fmt"
	"os"
	"strings"
)

func SplitLines(s string) []string {
	return strings.Split(strings.TrimSpace(s), "\n")
}

func ReadFileLines(path string) ([]string, error) {
	fileBytes, error := os.ReadFile(path)
	if error != nil {
		fmt.Println("Error reading file:", error)
		return nil, error
	}
	fileString := string(fileBytes)

	return SplitLines(fileString), error
}

func main() {
	// Read file lines.
	lines, err := ReadFileLines("input.txt")
	if err != nil {
		fmt.Println("Failed to read lines from file.")
		return
	}
}
