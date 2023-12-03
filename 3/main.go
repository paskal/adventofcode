package main

import (
	_ "embed"
	"log"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string
var emptyLine = "."

func main() {
	var sum int
	rawLines := strings.Split(input, "\n")
	lines := make([][]string, len(rawLines))
	for x, line := range rawLines {
		lines[x] = make([]string, len(line))
		for y, c := range strings.Split(line, "") {
			lines[x][y] = c
		}
	}
	isDigit := regexp.MustCompile(`\d`)
	for y, line := range lines {
		var currentNumber string
		var numberHasAdjustedSymbol []int
		for x, c := range line {
			if c != emptyLine && isDigit.MatchString(c) {
				currentNumber += c
				if numberHasAdjustedSymbol == nil {
					numberHasAdjustedSymbol = hasAdjustedSymbols(lines, x, y)
				}
			}
			if !isDigit.MatchString(c) && currentNumber != "" {
				if numberHasAdjustedSymbol != nil {
					digit, _ := strconv.Atoi(currentNumber)
					sum += digit
				}
				currentNumber = ""
				numberHasAdjustedSymbol = nil
			}
		}
		if numberHasAdjustedSymbol != nil && currentNumber != "" {
			digit, _ := strconv.Atoi(currentNumber)
			sum += digit
		}
		currentNumber = ""
		numberHasAdjustedSymbol = nil
	}
	log.Printf("Sum of all detail parts: %d", sum)
}

func hasAdjustedSymbols(lines [][]string, x int, y int) []int {
	isNotDigitOrPoint := regexp.MustCompile(`[^\d.]`)
	notTopLine := x > 0
	notLastLine := x < len(lines[y])-1
	notFirstColumn := y > 0
	notLastColumn := y < len(lines)-1
	if notFirstColumn {
		if notTopLine {
			if isNotDigitOrPoint.MatchString(lines[y-1][x-1]) {
				return []int{x - 1, y - 1}
			}
		}
		if isNotDigitOrPoint.MatchString(lines[y-1][x]) {
			return []int{x, y - 1}
		}
		if notLastLine {
			if isNotDigitOrPoint.MatchString(lines[y-1][x+1]) {
				return []int{x + 1, y - 1}
			}
		}
	}

	if notTopLine {
		if isNotDigitOrPoint.MatchString(lines[y][x-1]) {
			return []int{x - 1, y}
		}
	}
	if notLastLine {
		if isNotDigitOrPoint.MatchString(lines[y][x+1]) {
			return []int{x + 1, y}
		}
	}

	if notLastColumn {
		if notTopLine {
			if isNotDigitOrPoint.MatchString(lines[y+1][x-1]) {
				return []int{x - 1, y + 1}
			}
		}
		if isNotDigitOrPoint.MatchString(lines[y+1][x]) {
			return []int{x, y + 1}
		}
		if notLastLine {
			if isNotDigitOrPoint.MatchString(lines[y+1][x+1]) {
				return []int{x + 1, y + 1}
			}
		}
	}
	return nil
}
