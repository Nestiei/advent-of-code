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
	var totalInvalidPart1, totalInvalidPart2 int = 0, 0

	// Read input.
	ranges, err := ReadFileLines("input.txt")
	if err != nil {
		fmt.Println("Failed to read lines from file.")
		return
	}

	// Process each range.
	for _, idRange := range ranges {
		fmt.Println("Range:", idRange)

		rangeLimits := strings.Split(idRange, "-")
		rangeMin, _ := strconv.Atoi(rangeLimits[0])
		rangeMax, _ := strconv.Atoi(rangeLimits[1])

		for number := rangeMin; number <= rangeMax; number++ {
			// Part 1.
			numberString := fmt.Sprintf("%d", number)
			numberStringLength := len(numberString)

			numberFirstHalf, _ := strconv.Atoi(numberString[:numberStringLength/2])
			numberSecondHalf, _ := strconv.Atoi(numberString[numberStringLength/2:])

			if numberFirstHalf == numberSecondHalf {
				totalInvalidPart1 += number
				fmt.Println("Invalid ID (Part 1):", number)
			}

			// Part 2.
			for indexLength := range numberStringLength / 2 {
				splitLength := indexLength + 1
				splitsFloat := float64(numberStringLength) / float64(splitLength)
				splits := int(splitsFloat)
				if splitsFloat != float64(splits) {
					continue
				}

				numberParts := make([]int, splits)
				for i := range splits {
					splitValue, _ := strconv.Atoi(numberString[i*splitLength : (i+1)*splitLength])
					numberParts[i] = splitValue
				}

				allEqual := true
				for _, numberPart := range numberParts {
					if numberPart != numberParts[0] {
						allEqual = false
						break
					}
				}

				if allEqual {
					totalInvalidPart2 += number
					fmt.Println("Invalid ID (Part 2):", number)
					break
				}
			}
		}
	}

	fmt.Println("Total invalid IDs -> Part 1:", totalInvalidPart1, "-> Part 2:", totalInvalidPart2)
}
