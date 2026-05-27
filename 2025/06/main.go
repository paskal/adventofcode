package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func parseInputPartOne(input string) [][]int {
	var digitsByLines [][]int
	rawLines := strings.Split(" "+input, "\n")
	digitsRegex := regexp.MustCompile(`\d+`)
	for i := range len(rawLines) - 1 {
		digitsByLines = append(digitsByLines, []int{})
		digitsLine := digitsRegex.FindAllString(rawLines[i], -1)
		for _, v := range digitsLine {
			num, _ := strconv.Atoi(v)
			digitsByLines[len(digitsByLines)-1] = append(digitsByLines[len(digitsByLines)-1], num)
		}
	}

	digitsByColumns := make([][]int, len(digitsByLines[0]))
	for i := range digitsByLines {
		for j, num := range digitsByLines[i] {
			digitsByColumns[j] = append(digitsByColumns[j], num)
		}
	}

	return digitsByColumns
}

func parseInputPartTwo(input string) [][]int {
	digits := [][]int{{}}
	rawLines := strings.Split(input, "\n")
	var maxLen int
	for _, line := range rawLines {
		if len(line) > maxLen {
			maxLen = len(line)
		}
	}
	var curColumn int
	for i := range maxLen {
		var rawColumn string
		for j := range len(rawLines) - 1 {
			if len(rawLines[j]) == i {
				continue
			}
			rawColumn += string(rawLines[j][i])
		}
		rawColumn = strings.Trim(rawColumn, " ")
		if rawColumn != "" {
			num, _ := strconv.Atoi(rawColumn)
			digits[curColumn] = append(digits[curColumn], num)
		} else {
			digits = append(digits, []int{})
			curColumn++
		}
	}

	return digits
}

func getOperations(input string, length int) []bool {
	rawLines := strings.Split(" "+input, "\n")
	operationsRegex := regexp.MustCompile(`[*+]`)
	operations := operationsRegex.FindAllString(rawLines[length], -1)
	var operationIsMultiplication []bool
	for _, op := range operations {
		operationIsMultiplication = append(operationIsMultiplication, op == "*")
	}
	return operationIsMultiplication
}

func calculateSum(digits [][]int, operationIsMultiplication []bool) int {
	var sum int
	for i, op := range operationIsMultiplication {
		var localSum int
		if op {
			localSum = 1
		}
		for _, d := range digits[i] {
			if op {
				localSum *= d
			} else {
				localSum += d
			}
		}
		sum += localSum
	}
	return sum
}

func main() {
	digitsPartOne := parseInputPartOne(input)
	digitsPartTwo := parseInputPartTwo(input)
	operationIsMultiplication := getOperations(input, len(digitsPartOne[0]))

	fmt.Printf("First part answer: %d\n", calculateSum(digitsPartOne, operationIsMultiplication))
	partTwoSum := calculateSum(digitsPartTwo, operationIsMultiplication)
	fmt.Printf("Second part answer: %d\n", partTwoSum)
}
