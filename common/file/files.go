package files

import (
	"fmt"
	"os"
	"strings"
)

func splitLines(data string) []string {
	lines := strings.Split(strings.TrimSpace(data), "\n")
	for i, line := range lines {
		lines[i] = strings.TrimSpace(line)
	}
	return lines
}

func ReadFile(path string) (string, error) {
	fileBytes, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return "", err
	}

	return string(fileBytes), nil
}

func ReadFileLines(path string) ([]string, error) {
	fileBytes, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}
	fileString := string(fileBytes)

	return splitLines(fileString), nil
}
