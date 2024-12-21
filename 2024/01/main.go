package main

import (
	_ "embed"
	"log"
	"math"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	numsLeft, numsRight := make([]int, 0), make([]int, 0)
	for i, s := range strings.Split(input, "\n") {
		nums := strings.Fields(s)
		if len(nums) != 2 {
			log.Fatalf("line %d: expected 2 numbers, got %d", i+1, len(nums))
		}
		num, err := strconv.Atoi(nums[0])
		if err != nil {
			log.Fatalf("line %d: error converting %s to integer: %v", i+1, nums[0], err)
		}
		numsLeft = append(numsLeft, num)
		num, err = strconv.Atoi(nums[1])
		if err != nil {
			log.Fatalf("line %d: error converting %s to integer: %v", i+1, nums[1], err)
		}
		numsRight = append(numsRight, num)
	}
	slices.Sort(numsLeft)
	slices.Sort(numsRight)
	var sum int
	for i := range numsLeft {
		sum += int(math.Abs(float64(numsRight[i] - numsLeft[i])))
	}
	log.Printf("Sum: %d", sum)
}
