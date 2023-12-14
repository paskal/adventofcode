package main

import (
	_ "embed"
	"fmt"
	"log"
	"slices"
	"strings"
)

//go:embed input.txt
var input string

type locationMap struct {
	location                             [][]string
	verticalSymmetry, horizontalSymmetry int
}

func (l *locationMap) String() string {
	var result string
	for y, row := range l.location {
		if y != 0 && y == l.horizontalSymmetry {
			result += strings.Repeat("—", len(row)) + "\n"
		}
		for x, c := range row {
			if x != 0 && x == l.verticalSymmetry {
				result += "|"
			}
			switch c {
			case "#":
				result += "#"
			default:
				result += " "
			}
		}
		result += "\n"
	}
	return result
}

func (l *locationMap) calculateSymmetry() {
	l.verticalSymmetry = findVerticalSymmetry(l.location)
	l.horizontalSymmetry = findHorizontalSymmetry(l.location)
}

func findHorizontalSymmetry(location [][]string) int {
	var result int
	var symmetry []int
	// calculate first row to learn all the possibilities and check only them after
	for y := 1; y < len(location); y++ {
		var noSymmetry bool
		var i int
		for yBack := y - 1; yBack >= 0; yBack-- {
			if y+i < len(location) && !slices.Equal(location[y+i], location[yBack]) {
				noSymmetry = true
				break
			}
			i++
		}
		if !noSymmetry {
			symmetry = append(symmetry, y)
		}
	}
	if len(symmetry) == 1 {
		result = symmetry[0]
	}
	return result
}

func findVerticalSymmetry(location [][]string) int {
	var result int
	var symmetry []int
	// calculate first row to learn all the possibilities and check only them after
	for x := 1; x < len(location[0]); x++ {
		var noSymmetry bool
		var i int
		for xBack := x - 1; xBack >= 0; xBack-- {
			if x+i < len(location[0]) && location[0][x+i] != location[0][xBack] {
				noSymmetry = true
				break
			}
			i++
		}
		if !noSymmetry {
			symmetry = append(symmetry, x)
		}
	}
	for _, row := range location {
		var newSymmetry []int
		for _, x := range symmetry {
			var i int
			var noSymmetry bool
			for xBack := x - 1; xBack >= 0; xBack-- {
				if x+i < len(row) && row[x+i] != row[xBack] {
					noSymmetry = true
					break
				}
				i++
			}
			if !noSymmetry {
				newSymmetry = append(newSymmetry, x)
			}
		}
		symmetry = newSymmetry
	}
	if len(symmetry) == 1 {
		result = symmetry[0]
	}
	return result
}

func main() {
	patterns := getPattern()

	var sumPartOne int
	for i, _ := range patterns {
		patterns[i].calculateSymmetry()
		fmt.Printf("%s\n", patterns[i])
		sumPartOne += patterns[i].horizontalSymmetry*100 + patterns[i].verticalSymmetry

	}
	log.Printf("Sum for part one: %d", sumPartOne)
}

func getPattern() []*locationMap {
	var patternIndex int
	patterns := []*locationMap{{}}
	for _, row := range strings.Split(input, "\n") {
		if row == "" {
			patternIndex++
			patterns = append(patterns, &locationMap{})
			continue
		}
		patterns[patternIndex].location = append(patterns[patternIndex].location, strings.Split(row, ""))
	}
	return patterns
}
