package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func surroundedBylessThanFourRolls(grid [][]int, x, y int) bool {
	// start with -1 as we will count the roll itself and need to deduct it from the sum
	surroudingRolls := -1

	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			surroudingRolls += grid[i][j]
		}
	}

	return surroudingRolls < 4
}

func calculateFirstPart(grid [][]int) int {
	var accessibleRolls int

	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[i])-1; j++ {
			if grid[i][j] == 0 {
				continue
			}
			if surroundedBylessThanFourRolls(grid, i, j) {
				accessibleRolls++
			}
		}
	}

	return accessibleRolls
}

func calculateAndGetNewGrid(grid [][]int) (int, [][]int) {
	var accessibleRolls int
	var newGrid = grid

	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[i])-1; j++ {
			if grid[i][j] == 0 {
				continue
			}
			if surroundedBylessThanFourRolls(grid, i, j) {
				newGrid[i][j] = 0
				accessibleRolls++
			}
		}
	}

	return accessibleRolls, newGrid
}

func calculateSecondPart(grid [][]int) int {
	var accessibleRolls, newlyAccessibleRolls int
	newGrid := grid

	for {
		newlyAccessibleRolls, newGrid = calculateAndGetNewGrid(newGrid)
		accessibleRolls += newlyAccessibleRolls
		if newlyAccessibleRolls == 0 {
			break
		}
	}

	return accessibleRolls
}

func main() {
	lines := strings.Count(input, "\n") + 1
	// making grid one size bigger than it is to avoid handling edges
	grid := make([][]int, lines+2)

	for i, s := range strings.Split(input, "\n") {
		grid[0] = make([]int, len(s)+2)
		grid[i+1] = make([]int, len(s)+2)
		for j := 0; j < len(s); j++ {
			if string(s[j]) == "@" {
				grid[i+1][j+1] = 1
			}
		}
	}
	grid[lines+1] = grid[0]

	fmt.Printf("Accessible rolls for part 1: %d\n", calculateFirstPart(grid))
	fmt.Printf("Accessible rolls for part 2: %d\n", calculateSecondPart(grid))
}
