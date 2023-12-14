package main

import (
	_ "embed"
	"fmt"
	"log"
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

func (l *locationMap) calculateSymmetry(withSmudges int) {
	l.verticalSymmetry = findVerticalSymmetry(l.location, withSmudges)
	l.horizontalSymmetry = findHorizontalSymmetry(l.location, withSmudges)
}

func findHorizontalSymmetry(location [][]string, withSmudges int) int {
	var result int
	var symmetry []int
	var symmetrySmudges []int
	// calculate first row to learn all the possibilities and check only them after
	for y := 1; y < len(location); y++ {
		allowedSmudges := withSmudges
		var noSymmetry bool
		var j int
		for yBack := y - 1; yBack >= 0; yBack-- {
			for x := 0; x < len(location[0]) && y+j < len(location); x++ {
				if location[y+j][x] != location[yBack][x] {
					if allowedSmudges > 0 {
						allowedSmudges--
						continue
					}
					noSymmetry = true
					break
				}
				if noSymmetry {
					break
				}
			}
			j++
		}
		if !noSymmetry {
			symmetry = append(symmetry, y)
			symmetrySmudges = append(symmetrySmudges, allowedSmudges)
		}
	}
	for i := 0; i < len(symmetry); i++ {
		// only allow results when no smudges are left unaccounted for
		if symmetrySmudges[i] == 0 {
			result = symmetry[i]
		}
	}
	return result
}

func findVerticalSymmetry(location [][]string, withSmudges int) int {
	var result int
	var symmetry []int
	var symmetrySmudges []int
	// calculate first row to learn all the possibilities and check only them after
	for x := 1; x < len(location[0]); x++ {
		allowedSmudges := withSmudges
		var noSymmetry bool
		var j int
		for xBack := x - 1; xBack >= 0; xBack-- {
			if x+j < len(location[0]) && location[0][x+j] != location[0][xBack] {
				if allowedSmudges > 0 {
					allowedSmudges--
					j++
					continue
				}
				noSymmetry = true
				break
			}
			j++
		}
		if !noSymmetry {
			symmetry = append(symmetry, x)
			symmetrySmudges = append(symmetrySmudges, allowedSmudges)
		}
	}

	// skip first line
	for rowNumber := 1; rowNumber < len(location); rowNumber++ {
		var newSymmetry []int
		var newSmudges []int
		for i, x := range symmetry {
			var j int
			var noSymmetry bool
			for xBack := x - 1; xBack >= 0; xBack-- {
				if x+j < len(location[rowNumber]) && location[rowNumber][x+j] != location[rowNumber][xBack] {
					if symmetrySmudges[i] > 0 {
						symmetrySmudges[i]--
						j++
						continue
					}
					noSymmetry = true
					break
				}
				j++
			}
			if !noSymmetry {
				newSymmetry = append(newSymmetry, x)
				newSmudges = append(newSmudges, symmetrySmudges[i])
			}
		}
		symmetry = newSymmetry
		symmetrySmudges = newSmudges
	}
	for i := 0; i < len(symmetry); i++ {
		// only allow results when no smudges are left unaccounted for
		if symmetrySmudges[i] == 0 {
			result = symmetry[i]
		}
	}
	return result
}

func main() {
	patterns := getPattern()

	var sumWithNoSmudges int
	var sumWithOneSmudge int
	for i, _ := range patterns {
		patterns[i].calculateSymmetry(0)
		sumWithNoSmudges += patterns[i].horizontalSymmetry*100 + patterns[i].verticalSymmetry
		fmt.Printf("with no smudges:\n%s\n", patterns[i])
		patterns[i].calculateSymmetry(1)
		fmt.Printf("with one smudge:\n%s\n", patterns[i])
		if patterns[i].horizontalSymmetry+patterns[i].verticalSymmetry == 0 {
			log.Printf("no one smudge simmetry found for pattern %d", i+1)
		}
		sumWithOneSmudge += patterns[i].horizontalSymmetry*100 + patterns[i].verticalSymmetry
	}
	log.Printf("Sum with zero smudges: %d , with one smudge: %d", sumWithNoSmudges, sumWithOneSmudge)
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
