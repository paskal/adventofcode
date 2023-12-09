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
	var sum int
	for _, row := range firstRowDigits {
		var nextNumber int
		deep := [][]int{row}
		for nextNumber == 0 {
			deep = append(deep, []int{})
			for i := 1; i < len(deep[len(deep)-2]); i++ {
				currentRow := deep[len(deep)-2]
				deep[len(deep)-1] = append(deep[len(deep)-1], currentRow[i]-currentRow[i-1])
			}

			allZeroes := true
			for _, n := range deep[len(deep)-1] {
				if n != 0 {
					allZeroes = false
					break
				}
			}
			if allZeroes {
				// nextNumber starts with 0 as last row is all zeroes
				for i := len(deep) - 2; i >= 0; i-- {
					lastItemCurrentRow := deep[i][len(deep[i])-1]
					nextNumber = nextNumber + lastItemCurrentRow
				}
			}
		}
		sum += nextNumber
	}
	log.Printf("Sum of rows: %d", sum)
}
