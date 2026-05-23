package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func biggestNumber(s string, budget int) int {
	lastIndex := -1
	var result string
	remainingBudget := budget
	for len(result) < budget {
		currentIndex := lastIndex + 1
		for i := currentIndex + 1; i < len(s)-(remainingBudget-1); i++ {
			if s[currentIndex] < s[i] {
				currentIndex = i
			}
		}
		result += string(s[currentIndex])
		lastIndex = currentIndex
		remainingBudget--
	}

	r, _ := strconv.ParseInt(result, 10, 64)
	return int(r)
}

func main() {
	var partOneSum, partTwoSum int
	for s := range strings.SplitSeq(input, "\n") {
		partOneSum += biggestNumber(s, 2)
		partTwoSum += biggestNumber(s, 12)
	}
	fmt.Printf("Sum jolts part 1: %v\n", partOneSum)
	fmt.Printf("Sum jolts part 2: %v\n", partTwoSum)
}
