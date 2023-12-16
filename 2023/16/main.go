package main

import (
	_ "embed"
	"log"
	"strings"
)

//go:embed input.txt
var input string

type node struct {
	value                                    string
	fromWest, fromEast, fromNorth, fromSouth bool
}

type location [][]node

func (l location) String() string {
	var result string
	for _, row := range l {
		for _, n := range row {
			if n.fromWest || n.fromEast || n.fromNorth || n.fromSouth {
				result += "#"
			} else {
				result += "."
			}
		}
		result += "\n"
	}
	return result
}

func (l location) countEnergized() int {
	var countEnergized int
	for _, row := range l {
		for _, n := range row {
			if n.fromSouth || n.fromNorth || n.fromWest || n.fromEast {
				countEnergized++
			}
		}
	}
	return countEnergized
}

func (l location) moveFromWest(x int, y int) {
	if x < 0 || y < 0 || x > len(l[0])-1 || y > len(l)-1 || l[y][x].fromWest {
		return
	}
	l[y][x].fromWest = true
	switch l[y][x].value {
	case ".", "-":
		l.moveFromWest(x+1, y)
	case "\\":
		l.moveFromNorth(x, y+1)
	case "/":
		l.moveFromSouth(x, y-1)
	case "|":
		l.moveFromNorth(x, y+1)
		l.moveFromSouth(x, y-1)
	}
}

func (l location) moveFromEast(x int, y int) {
	if x < 0 || y < 0 || x > len(l[0])-1 || y > len(l)-1 || l[y][x].fromEast {
		return
	}
	l[y][x].fromEast = true
	switch l[y][x].value {
	case ".", "-":
		l.moveFromEast(x-1, y)
	case "\\":
		l.moveFromSouth(x, y-1)
	case "/":
		l.moveFromNorth(x, y+1)
	case "|":
		l.moveFromNorth(x, y+1)
		l.moveFromSouth(x, y-1)
	}
}

func (l location) moveFromNorth(x int, y int) {
	if x < 0 || y < 0 || x == len(l[0])-1 || y > len(l)-1 || l[y][x].fromNorth {
		return
	}
	l[y][x].fromNorth = true
	switch l[y][x].value {
	case ".", "|":
		l.moveFromNorth(x, y+1)
	case "\\":
		l.moveFromWest(x+1, y)
	case "/":
		l.moveFromEast(x-1, y)
	case "-":
		l.moveFromWest(x+1, y)
		l.moveFromEast(x-1, y)
	}
}

func (l location) moveFromSouth(x int, y int) {
	if x < 0 || y < 0 || x > len(l[0])-1 || y > len(l)-1 || l[y][x].fromSouth {
		return
	}
	l[y][x].fromSouth = true
	switch l[y][x].value {
	case ".", "|":
		l.moveFromSouth(x, y-1)
	case "\\":
		l.moveFromEast(x-1, y)
	case "/":
		l.moveFromWest(x+1, y)
	case "-":
		l.moveFromWest(x+1, y)
		l.moveFromEast(x-1, y)
	}
}

func main() {
	locationMap := getMap(input)
	locationMap.moveFromWest(0, 0)
	log.Printf("Energized tiles: %d", locationMap.countEnergized())
}

func getMap(input string) (locationMap location) {
	for y, s := range strings.Split(input, "\n") {
		locationMap = append(locationMap, make([]node, len(strings.Split(s, ""))))
		for x, c := range strings.Split(s, "") {
			locationMap[y][x] = node{value: c}
		}
	}
	return locationMap
}
