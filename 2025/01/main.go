package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadFileLines(path string) ([]string, error) {
	fileBytes, error := os.ReadFile(path)
	if error != nil {
		fmt.Println("Error reading file:", error)
		return nil, error
	}
	fileString := string(fileBytes)

	return SplitLines(fileString), error
}

func SplitLines(s string) []string {
	return strings.Split(strings.TrimSpace(s), "\n")
}

func GetTurns(line string) (int, error) {
	line = strings.TrimSpace(line)
	number, err := strconv.Atoi(line[1:])
	if err != nil {
		return 0, err
	}

	switch {
	case strings.HasPrefix(line, "L"):
		return -1 * number, nil
	case strings.HasPrefix(line, "R"):
		return number, nil
	}

	return 0, fmt.Errorf("invalid line format")
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func main() {
	var dial int = 50
	var password, passwordClick int = 0, 0

	// Read file lines.
	lines, err := ReadFileLines("input.txt")
	if err != nil {
		fmt.Println("Failed to read lines from file.")
		return
	}

	// Process each line.
	for _, line := range lines {
		fmt.Println(line)

		turns, err := GetTurns(line)
		if err != nil {
			fmt.Println("Error parsing turns:", err)
			return
		}

		var countZeros int = 0
		for i := range Abs(turns) {
			if turns > 0 && (dial+i+1)%100 == 0 {
				countZeros += 1
			} else if turns < 0 && (dial-(i+1))%100 == 0 {
				countZeros += 1
			}
		}

		passwordClick += countZeros
		fmt.Println("Crossed zero", countZeros, "times. Total:", passwordClick)

		dial = (dial + turns) % 100
		fmt.Println("Dial rotated by", turns, "to position:", dial)
		if dial == 0 {
			password += 1
			fmt.Println("Password incremented to:", password)
		}
	}

	fmt.Println("Part 1 Password:", password, "-> Part 2 Password:", passwordClick)
}
