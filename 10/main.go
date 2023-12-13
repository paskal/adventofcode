package main

import (
	_ "embed"
	"fmt"
	"log"
	"strings"
)

//go:embed input.txt
var input string

type node struct {
	value                            string
	partOfPath, left, right, visited bool
}

func (n node) String() string {
	if n.left {
		return "░"
	}
	if n.right {
		return "▓"
	}
	if !n.partOfPath {
		return " "
	}
	switch n.value {
	case "L":
		return "└"
	case "F":
		return "┌"
	case "J":
		return "┘"
	case "7":
		return "┐"
	case "-":
		return "─"
	case "|":
		return "│"
	}
	return n.value
}

type location [][]node

type direction func(int, int) (bool, int, int)

func (l location) String() string {
	var result string
	for _, row := range l {
		for _, n := range row {
			result += n.String()
		}
		result += "\n"
	}
	return result
}

// moves in the first possible location
func (l location) move(startX int, startY int) (x int, y int) {
	l[startY][startX].partOfPath = true
	l[startY][startX].left = false
	l[startY][startX].right = false
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
			if x > 0 && !l[y-1][x-1].partOfPath {
				l[y-1][x-1].left = true
			}
			if x < len(l[0])-1 && !l[y-1][x+1].partOfPath {
				l[y-1][x+1].right = true
			}
			return true, x, y - 1
		}
	}
	return false, 0, 0
}

func (l location) south(x int, y int) (bool, int, int) {
	if y < len(l)-1 {
		nextCell := l[y+1][x]
		if !nextCell.partOfPath && strings.ContainsAny(nextCell.value, "|LJS") {
			if x > 0 && !l[y+1][x-1].partOfPath {
				l[y+1][x-1].right = true
			}
			if x < len(l[0])-1 && !l[y+1][x+1].partOfPath {
				l[y+1][x+1].left = true
			}
			return true, x, y + 1
		}
	}
	return false, 0, 0
}

func (l location) west(x int, y int) (bool, int, int) {
	if x > 0 {
		nextCell := l[y][x-1]
		if !nextCell.partOfPath && strings.ContainsAny(nextCell.value, "-FLS") {
			if y > 0 && !l[y-1][x-1].partOfPath {
				l[y-1][x-1].right = true
			}
			if y < len(l)-1 && !l[y+1][x-1].partOfPath {
				l[y+1][x-1].left = true
			}
			if nextCell.value == "F" {
				if x > 1 && !l[y][x-2].partOfPath {
					l[y][x-2].right = true
				}
			}
			if nextCell.value == "L" {
				if x > 1 && !l[y][x-2].partOfPath {
					l[y][x-2].left = true
				}
			}
			return true, x - 1, y
		}
	}
	return false, 0, 0
}

func (l location) east(x int, y int) (bool, int, int) {
	if x < len(l[0])-1 {
		nextCell := l[y][x+1]
		if !nextCell.partOfPath && strings.ContainsAny(nextCell.value, "-J7S") {
			if y > 0 && !l[y-1][x+1].partOfPath {
				l[y-1][x+1].left = true
			}
			if y < len(l)-1 && !l[y+1][x+1].partOfPath {
				l[y+1][x+1].right = true
			}
			if nextCell.value == "J" {
				if x < len(l[0])-2 && !l[y][x+2].partOfPath {
					l[y][x+2].right = true
				}
			}
			if nextCell.value == "7" {
				if x < len(l[0])-2 && !l[y][x+2].partOfPath {
					l[y][x+2].left = true
				}
			}
			return true, x + 1, y
		}
	}
	return false, 0, 0
}

// marks all nodes as visited and counts dots in them.
// must be called on a node which is not "partOfPath"
func (l location) mark(x int, y int) (area int) {
	l[y][x].visited = true
	area++
	if y > 0 && !l[y-1][x].partOfPath && !l[y-1][x].visited {
		area += l.mark(x, y-1)
	}
	if y < len(l)-1 && !l[y+1][x].partOfPath && !l[y+1][x].visited {
		area += l.mark(x, y+1)
	}
	if x > 0 && !l[y][x-1].partOfPath && !l[y][x-1].visited {
		area += l.mark(x-1, y)
	}
	if x < len(l[0])-1 && !l[y][x+1].partOfPath && !l[y][x+1].visited {
		area += l.mark(x+1, y)
	}

	return area
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
	locationMap[startY][startX].partOfPath = true
	locationMap[startY][startX].left = false
	locationMap[startY][startX].right = false

	var leftArea int
	var rightArea int
	for markY, row := range locationMap {
		for markX, _ := range row {
			if locationMap[markY][markX].left && !locationMap[markY][markX].visited {
				leftArea += locationMap.mark(markX, markY)
			}
			if locationMap[markY][markX].right && !locationMap[markY][markX].visited {
				rightArea += locationMap.mark(markX, markY)
			}
		}
	}

	log.Printf("Half of the path length: %d, area left %d, area right %d", pathLength/2, leftArea, rightArea)
	fmt.Printf("Map:\n%v", locationMap)
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
