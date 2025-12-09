package main

import (
	"fmt"
	"math"
	"nestiei/files"
	"slices"
)

func getLargestJoltage(bank string, batteriesNeeded int) int {
	fmt.Print("Bank ", bank)

	bankLength := len(bank)
	joltages := make([]int, bankLength)
	for index, joltageChar := range bank {
		joltage := int(joltageChar - '0')
		joltages[index] = joltage
	}

	largestJoltage := 0
	indexLatestJoltage := 0
	batteriesMissing := batteriesNeeded
	for ; batteriesMissing > 0; batteriesMissing-- {
		joltagesSlice := joltages[indexLatestJoltage : bankLength-batteriesMissing+1]

		highestJoltage := slices.Max(joltagesSlice)
		highestJoltageIndex := slices.Index(joltagesSlice, highestJoltage)
		indexLatestJoltage += highestJoltageIndex + 1

		largestJoltage += highestJoltage * int(math.Pow10(batteriesMissing-1))
	}

	fmt.Println(" -> Largest joltage:", largestJoltage)

	return largestJoltage
}

func main() {
	banks, err := files.ReadFileLines("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	totalJoltagePart1, totalJoltagePart2 := 0, 0
	for _, bank := range banks {
		totalJoltagePart1 += getLargestJoltage(bank, 2)
		totalJoltagePart2 += getLargestJoltage(bank, 12)
	}

	fmt.Println("Total joltage (Part 1):", totalJoltagePart1)
	fmt.Println("Total joltage (Part 2):", totalJoltagePart2)
}
