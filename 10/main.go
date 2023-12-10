package main

import (
	_ "embed"
	"log"
	"strings"
)

//go:embed input.txt
var input string

type node struct {
	value      string
	partOfPath bool
}

type location [][]node

type direction func(int, int) (bool, int, int)

// moves in the first possible location
func (l location) move(startX int, startY int) (x int, y int) {
	l[startY][startX].partOfPath = true
	var possibleDirection []direction
	switch l[startY][startX].value {
	case "S":
		possibleDirection = []direction{l.north, l.south, l.west, l.east}
	case "|":
		possibleDirection = []direction{l.north, l.south}
	case "-":
		possibleDirection = []direction{l.west, l.east}
	case "L":
		possibleDirection = []direction{l.north, l.east}
	case "J":
		possibleDirection = []direction{l.north, l.west}
	case "7":
		possibleDirection = []direction{l.south, l.west}
	case "F":
		possibleDirection = []direction{l.south, l.east}
	}
	var found bool
	for _, d := range possibleDirection {
		if found, x, y = d(startX, startY); found {
			break
		}
	}
	return x, y
}

func (l location) north(x int, y int) (bool, int, int) {
	if y > 0 {
		nextCell := l[y-1][x]
		if !nextCell.partOfPath && strings.ContainsAny(nextCell.value, "|7FS") {
			return true, x, y - 1
		}
	}
	return false, 0, 0
}

func (l location) south(x int, y int) (bool, int, int) {
	if y < len(l)-1 {
		nextCell := l[y+1][x]
		if !nextCell.partOfPath && strings.ContainsAny(nextCell.value, "|LJS") {
			return true, x, y + 1
		}
	}
	return false, 0, 0
}

func (l location) west(x int, y int) (bool, int, int) {
	if x > 0 {
		nextCell := l[y][x-1]
		if !nextCell.partOfPath && strings.ContainsAny(nextCell.value, "-FLS") {
			return true, x - 1, y
		}
	}
	return false, 0, 0
}

func (l location) east(x int, y int) (bool, int, int) {
	if x < len(l[0])-1 {
		nextCell := l[y][x+1]
		if !nextCell.partOfPath && strings.ContainsAny(nextCell.value, "-J7S") {
			return true, x + 1, y
		}
	}
	return false, 0, 0
}

func main() {
	locationMap, startX, startY := getMap()
	x, y := startX, startY
	// make two steps away to not step onto starting position away, and mark it as unvisited
	x, y = locationMap.move(x, y)
	x, y = locationMap.move(x, y)
	locationMap[startY][startX].partOfPath = false // otherwise we won't come back to start
	pathLength := 2
	for {
		x, y = locationMap.move(x, y)
		pathLength++
		if x == startX && y == startY {
			break
		}
	}
	log.Printf("Half of the path length: %d", pathLength/2)
}

func getMap() (locationMap location, animalX int, animalY int) {
	for y, s := range strings.Split(input, "\n") {
		locationMap = append(locationMap, make([]node, len(strings.Split(s, ""))))
		for x, c := range strings.Split(s, "") {
			locationMap[y][x] = node{value: c}
			if c == "S" {
				animalX, animalY = x, y
			}
		}
	}
	return locationMap, animalX, animalY
}
