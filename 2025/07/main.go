package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

type coordinate struct {
	x, y int
}

func main() {
	var gameMap [][]bool
	var visitedMap [][]bool
	var start coordinate
	for y, line := range strings.Split(input, "\n") {
		gameMap = append(gameMap, []bool{})
		for x := range len(line) {
			if string(line[x]) == "S" {
				start.x, start.y = x, y
			}
			if string(line[x]) == "^" {
				gameMap[y] = append(gameMap[y], true)
			} else {
				gameMap[y] = append(gameMap[y], false)
			}

		}
		visitedMap = append(visitedMap, make([]bool, len(gameMap[y])))
	}

	var numberOfSplits int
	queue := []coordinate{start}
	for len(queue) != 0 {
		start, queue = queue[0], queue[1:]
		if start.x < 0 || start.x == len(gameMap[0]) || start.y == len(gameMap) || visitedMap[start.y][start.x] {
			continue
		}
		visitedMap[start.y][start.x] = true
		if gameMap[start.y][start.x] {
			queue = append(queue, coordinate{start.x - 1, start.y})
			queue = append(queue, coordinate{start.x + 1, start.y})
			numberOfSplits++
			continue
		}
		queue = append(queue, coordinate{start.x, start.y + 1})
	}

	fmt.Printf("Number of splits: %d\n", numberOfSplits)
}
