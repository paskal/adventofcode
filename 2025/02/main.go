package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func isInvalidPart1(d int) bool {
	numberOfDigits := int(math.Log10(float64(d))) + 1
	// uneven number of digits can't contain a pair
	if numberOfDigits%2 != 0 {
		return false
	}
	numAsString := strconv.Itoa(d)
	if numAsString[:numberOfDigits/2] == numAsString[numberOfDigits/2:] {
		return true
	}
	return false
}

func isInvalidPart2(d int) bool {
	numberOfDigits := int(math.Log10(float64(d))) + 1
	numAsString := strconv.Itoa(d)
	for i := 1; i <= numberOfDigits/2; i++ {
		if numberOfDigits%i != 0 {
			continue
		}
		gotGoodSegment := true
		segments := numberOfDigits / i
		for j := 1; j < segments; j++ {
			if numAsString[:i] != numAsString[i*j:i*j+i] {
				gotGoodSegment = false
				break
			}
		}
		if gotGoodSegment {
			return true
		}
	}
	return false
}

func main() {
	r := make([]string, 2)
	var invalidIdsSumPart1, invalidIdsSumPart2 int
	for raw := range strings.SplitSeq(input, ",") {
		r = strings.Split(raw, "-")
		beginning, _ := strconv.ParseInt(r[0], 10, 64)
		end, _ := strconv.ParseInt(r[1], 10, 64)
		for d := int(beginning); d <= int(end); d++ {
			if isInvalidPart1(d) {
				invalidIdsSumPart1 += d
			}
			if isInvalidPart2(d) {
				invalidIdsSumPart2 += d
			}
		}
	}
	fmt.Printf("Sum of invalid ids part 1: %v\n", invalidIdsSumPart1)
	fmt.Printf("Sum of invalid ids part 2: %v\n", invalidIdsSumPart2)
}
