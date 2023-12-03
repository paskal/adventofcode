package main

import (
	_ "embed"
	"fmt"
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
	var gears = map[string][]int{}
	isDigit := regexp.MustCompile(`\d`)
	for y, line := range lines {
		var currentNumber string
		var numberHasAdjustedSymbol []int
		var symbol string
		var adjustedToGear bool
		for x, c := range line {
			if c != emptyLine && isDigit.MatchString(c) {
				currentNumber += c
				if numberHasAdjustedSymbol == nil {
					numberHasAdjustedSymbol, symbol = hasAdjustedSymbols(lines, x, y)
					if symbol == "*" {
						adjustedToGear = true
					}
				}
			}
			if !isDigit.MatchString(c) && currentNumber != "" {
				if numberHasAdjustedSymbol != nil {
					digit, _ := strconv.Atoi(currentNumber)
					sum += digit
					if adjustedToGear {
						gearCoordinates := fmt.Sprintf("%d%d", numberHasAdjustedSymbol[0], numberHasAdjustedSymbol[1])
						gears[gearCoordinates] = append(gears[gearCoordinates], digit)
					}
				}
				currentNumber = ""
				numberHasAdjustedSymbol = nil
				adjustedToGear = false
			}
		}
		if numberHasAdjustedSymbol != nil && currentNumber != "" {
			digit, _ := strconv.Atoi(currentNumber)
			sum += digit
			if adjustedToGear {
				gearCoordinates := fmt.Sprintf("%d%d", numberHasAdjustedSymbol[0], numberHasAdjustedSymbol[1])
				gears[gearCoordinates] = append(gears[gearCoordinates], digit)
			}
		}
		currentNumber = ""
		numberHasAdjustedSymbol = nil
		adjustedToGear = false
	}
	var sumGears int
	for _, parts := range gears {
		if len(parts) == 2 {
			sumGears += parts[0] * parts[1]
		}
	}
	log.Printf("Sum of all detail parts: %d, sum of gear and multiplications: %d", sum, sumGears)
}

func hasAdjustedSymbols(lines [][]string, x int, y int) (xy []int, symbol string) {
	isNotDigitOrPoint := regexp.MustCompile(`[^\d.]`)
	notTopLine := x > 0
	notLastLine := x < len(lines[y])-1
	notFirstColumn := y > 0
	notLastColumn := y < len(lines)-1
	if notFirstColumn {
		if notTopLine {
			if isNotDigitOrPoint.MatchString(lines[y-1][x-1]) {
				return []int{x - 1, y - 1}, lines[y-1][x-1]
			}
		}
		if isNotDigitOrPoint.MatchString(lines[y-1][x]) {
			return []int{x, y - 1}, lines[y-1][x]
		}
		if notLastLine {
			if isNotDigitOrPoint.MatchString(lines[y-1][x+1]) {
				return []int{x + 1, y - 1}, lines[y-1][x+1]
			}
		}
	}

	if notTopLine {
		if isNotDigitOrPoint.MatchString(lines[y][x-1]) {
			return []int{x - 1, y}, lines[y][x-1]
		}
	}
	if notLastLine {
		if isNotDigitOrPoint.MatchString(lines[y][x+1]) {
			return []int{x + 1, y}, lines[y][x+1]
		}
	}

	if notLastColumn {
		if notTopLine {
			if isNotDigitOrPoint.MatchString(lines[y+1][x-1]) {
				return []int{x - 1, y + 1}, lines[y+1][x-1]
			}
		}
		if isNotDigitOrPoint.MatchString(lines[y+1][x]) {
			return []int{x, y + 1}, lines[y+1][x]
		}
		if notLastLine {
			if isNotDigitOrPoint.MatchString(lines[y+1][x+1]) {
				return []int{x + 1, y + 1}, lines[y+1][x+1]
			}
		}
	}
	return nil, ""
}
