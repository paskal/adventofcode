package main

import (
	_ "embed"
	"log"
	"math"
	"strings"
)

//go:embed input.txt
var input string

type location [][]string

func (l location) String() string {
	var result string
	for _, row := range l {
		result += strings.Join(row, "") + "\n"
	}
	return result
}

func (l location) moveNorth(x int, y int) int {
	newY := y
	for testY := y - 1; testY >= 0; testY-- {
		if l[testY][x] != "." {
			break
		}
		newY = testY
	}
	if newY != y {
		l[y][x] = "."
		l[newY][x] = "O"
	}
	return int(math.Abs(float64(newY - len(l))))
}

func main() {
	loc := getPattern()
	var weightNorth int
	for y := 0; y < len(loc); y++ {
		for x := 0; x < len(loc[0]); x++ {
			if loc[y][x] == "O" {
				weightNorth += loc.moveNorth(x, y)
			}
		}
	}
	log.Printf("map after:\n%v", loc)
	log.Printf("Sum of weight north: %d", weightNorth)
}

func getPattern() location {
	var loc location
	for _, row := range strings.Split(input, "\n") {
		loc = append(loc, strings.Split(row, ""))
	}
	return loc
}
