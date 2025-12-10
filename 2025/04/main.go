package main

import (
	"fmt"
	"nestiei/files"
)

func buildMatrix(data []string) [][]int {
	matrix := make([][]int, len(data))
	for x, line := range data {
		matrix[x] = make([]int, len(line))
		for y, char := range line {
			switch char {
			case '@':
				matrix[x][y] = 1
			default:
				matrix[x][y] = 0
			}
		}
	}

	return matrix
}

func calculateRemovedRolls(matrix *[][]int, limit int) (bool, int) {
	totalRemovedRolls := 0

	for x := range *matrix {
		for y := range len((*matrix)[x]) {
			if (*matrix)[x][y] == 0 {
				continue
			}

			adjacentRolls := 0

			for xOffset := -1; xOffset <= 1; xOffset++ {
				for yOffset := -1; yOffset <= 1; yOffset++ {
					xResolved := x + xOffset
					yResolved := y + yOffset
					if xOffset == 0 && yOffset == 0 || xResolved < 0 || xResolved >= len(*matrix) || yResolved < 0 || yResolved >= len((*matrix)[x]) {
						continue
					}

					adjacentRolls += (*matrix)[xResolved][yResolved]
				}
			}

			if adjacentRolls < limit {
				totalRemovedRolls++
				(*matrix)[x][y] = 0
			}
		}
	}

	return totalRemovedRolls > 0, totalRemovedRolls
}

func main() {
	lines, err := files.ReadFileLines("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	matrix := buildMatrix(lines)
	fmt.Println("Matrix:", matrix)

	hasRemovedRolls := true
	totalRemovedRolls := 0
	for index := 0; hasRemovedRolls; index++ {
		hasRemovedRollsIteration, removedRolls := calculateRemovedRolls(&matrix, 4)
		hasRemovedRolls = hasRemovedRollsIteration

		fmt.Println(fmt.Sprintf("Removed rolls iteration #%d:", index+1), removedRolls)
		totalRemovedRolls += removedRolls
	}

	fmt.Println("Total removed rolls:", totalRemovedRolls)
}
