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

func (l location) moveStones(direction string) {
	switch direction {
	case "north":
		for y := 0; y < len(l); y++ {
			for x := 0; x < len(l[0]); x++ {
				if l[y][x] == "O" {
					l.moveStoneNorth(x, y)
				}
			}
		}
	case "south":
		for y := len(l) - 1; y >= 0; y-- {
			for x := 0; x < len(l[0]); x++ {
				if l[y][x] == "O" {
					l.moveStoneSouth(x, y)
				}
			}
		}
	case "west":
		for y := 0; y < len(l); y++ {
			for x := 0; x < len(l[0]); x++ {
				if l[y][x] == "O" {
					l.moveStoneWest(x, y)
				}
			}
		}
	case "east":
		for y := 0; y < len(l); y++ {
			for x := len(l[0]) - 1; x >= 0; x-- {
				if l[y][x] == "O" {
					l.moveStoneEast(x, y)
				}
			}
		}
	}
}

func (l location) moveStoneNorth(x int, y int) {
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
}

func (l location) moveStoneSouth(x int, y int) {
	newY := y
	for testY := y + 1; testY < len(l); testY++ {
		if l[testY][x] != "." {
			break
		}
		newY = testY
	}
	if newY != y {
		l[y][x] = "."
		l[newY][x] = "O"
	}
}

func (l location) moveStoneWest(x int, y int) {
	newX := x
	for testX := x - 1; testX >= 0; testX-- {
		if l[y][testX] != "." {
			break
		}
		newX = testX
	}
	if newX != x {
		l[y][x] = "."
		l[y][newX] = "O"
	}
}

func (l location) moveStoneEast(x int, y int) {
	newX := x
	for testX := x + 1; testX < len(l[0]); testX++ {
		if l[y][testX] != "." {
			break
		}
		newX = testX
	}
	if newX != x {
		l[y][x] = "."
		l[y][newX] = "O"
	}
}

func main() {
	loc := getPattern()
	loc.moveStones("north")
	log.Printf("map after moving stones north:\n%v", loc)
	weightNorth := calculateLoad(loc)

	cyclesNumber := 1000000000
	scoreAfterManyRotation := getScoreAfterFullRotations(cyclesNumber, getPattern())
	log.Printf("Sum of weight north %d , after %d cycles: %d", weightNorth, cyclesNumber, scoreAfterManyRotation)
}

func getScoreAfterFullRotations(rotations int, loc location) int {
	var locScores []int
	var locComparableList []string
	var firstMatch, repetitionStart int
	var foundFirstMatch, foundFullCycle bool
	for i := 0; i < rotations && !foundFullCycle; i++ {
		loc.moveStones("north")
		loc.moveStones("west")
		loc.moveStones("south")
		loc.moveStones("east")
		var foundMatch bool
		for j, _ := range locComparableList {
			if locComparableList[j] == loc.String() {
				foundMatch = true
				if firstMatch == 0 {
					firstMatch = j
					repetitionStart = i + 1
				}
				if firstMatch != 0 && j == 0 {
					foundFullCycle = true
				}
			}
		}
		if !foundFirstMatch && firstMatch != 0 {
			locScores = locScores[firstMatch:]
			locComparableList = locComparableList[firstMatch:]
			foundFirstMatch = true
		}
		if !foundMatch {
			locScores = append(locScores, calculateLoad(loc))
			locComparableList = append(locComparableList, loc.String())
		}
	}
	// special case from 0 to first match (10 in given example)
	if repetitionStart == 0 && !foundFirstMatch {
		return locScores[len(locScores)-1]
	}
	if !foundFullCycle {
		log.Printf("found it")
	}
	// special case from foundFirstMatch to
	locationIndexAtRotations := (rotations - repetitionStart) % len(locComparableList)
	return locScores[locationIndexAtRotations]
}

func calculateLoad(loc location) int {
	var totalLoad int
	for y := 0; y < len(loc); y++ {
		for x := 0; x < len(loc[0]); x++ {
			if loc[y][x] == "O" {
				totalLoad += int(math.Abs(float64(y - len(loc))))
			}
		}
	}
	return totalLoad
}

func getPattern() location {
	var loc location
	for _, row := range strings.Split(input, "\n") {
		loc = append(loc, strings.Split(row, ""))
	}
	return loc
}
