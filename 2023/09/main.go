package main

import (
	_ "embed"
	"log"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	var firstRowDigits [][]int
	for i, line := range strings.Split(input, "\n") {
		firstRowDigits = append(firstRowDigits, []int{})
		for _, s := range strings.Split(line, " ") {
			num, _ := strconv.Atoi(s)
			firstRowDigits[i] = append(firstRowDigits[i], num)
		}
	}
	sumForward := nextSequencesItemsSum(true, firstRowDigits)
	sumBackward := nextSequencesItemsSum(false, firstRowDigits)
	log.Printf("Sum of rows forward: %d, backward: %d", sumForward, sumBackward)
}

func nextSequencesItemsSum(forward bool, digits [][]int) int {
	var result int
	for _, row := range digits {
		var nextInSequence int
		deep := [][]int{row} // initialise with the first row we got
		for {
			deep = append(deep, []int{})
			if forward {
				for i := 1; i < len(deep[len(deep)-2]); i++ {
					currentRow := deep[len(deep)-2]
					deep[len(deep)-1] = append(deep[len(deep)-1], currentRow[i]-currentRow[i-1])
				}
			} else {
				for i := len(deep[len(deep)-2]) - 1; i > 0; i-- {
					currentRow := deep[len(deep)-2]
					deep[len(deep)-1] = append([]int{currentRow[i] - currentRow[i-1]}, deep[len(deep)-1]...)
				}
			}

			// first and last element are enough to see if the row is all zeroes.
			// I found not-all-zeroes case with two zeroes in a row but none with first and last zeroes.
			allZeroes := deep[len(deep)-1][0] == 0 && deep[len(deep)-1][len(deep[len(deep)-1])-1] == 0
			if allZeroes {
				// nextInSequence starts with 0 as last row is all zeroes
				// for that reason we also skip that last row in the loop below
				for i := len(deep) - 2; i >= 0; i-- {
					if forward {
						lastItemCurrentRow := deep[i][len(deep[i])-1]
						nextInSequence = nextInSequence + lastItemCurrentRow
					} else {
						firstItemCurrentRow := deep[i][0]
						nextInSequence = firstItemCurrentRow - nextInSequence
					}
				}
				break
			}
		}
		result += nextInSequence
	}
	return result
}
