package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func SplitRanges(s string) []string {
	return strings.Split(strings.TrimSpace(s), ",")
}

func ReadFileLines(path string) ([]string, error) {
	fileBytes, error := os.ReadFile(path)
	if error != nil {
		fmt.Println("Error reading file:", error)
		return nil, error
	}
	fileString := string(fileBytes)

	return SplitRanges(fileString), error
}

func main() {
	var total int = 0

	// Read input.
	ranges, err := ReadFileLines("input.txt")
	if err != nil {
		fmt.Println("Failed to read lines from file.")
		return
	}

	// Process each range.
	for _, idRange := range ranges {
		fmt.Println(idRange)

		rangeLimits := strings.Split(idRange, "-")
		rangeMin, _ := strconv.Atoi(rangeLimits[0])
		rangeMax, _ := strconv.Atoi(rangeLimits[1])

		for number := rangeMin; number <= rangeMax; number++ {
			numberString := fmt.Sprintf("%d", number)
			numberStringLength := len(numberString)
			if numberStringLength%2 != 0 {
				continue
			}

			numberFirstHalf, _ := strconv.Atoi(numberString[:numberStringLength/2])
			numberSecondHalf, _ := strconv.Atoi(numberString[numberStringLength/2:])

			if numberFirstHalf == numberSecondHalf {
				total += number
				fmt.Println("Found invalid ID:", number)
			}
		}
	}

	fmt.Println("Total sum of invalid IDs:", total)
}
