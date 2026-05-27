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

func walkTimeline(gameMap [][]bool, start coordinate) int {
	visitedMapCur := make([]int, len(gameMap[0]))
	visitedMapPrev := make([]int, len(gameMap[0]))
	visitedMapPrev[start.x] = 1

	for y, _ := range gameMap {
		if y == start.y {
			continue
		}
		clear(visitedMapCur)
		for x, countAbove := range visitedMapPrev {
			if countAbove == 0 {
				continue
			}
			if !gameMap[y][x] {
				visitedMapCur[x] += countAbove
				continue
			}
			// if we bumped into the split, calculate left and right propagations right away
			startLeftCount := countAbove
			for leftX := x - 1; leftX >= 0; leftX-- {
				// stop only once we hit the first non-split node
				if !gameMap[y][leftX] {
					visitedMapCur[leftX] += startLeftCount
					break
				}
			}
			startRightCount := countAbove
			for rightX := x + 1; rightX < len(gameMap[y]); rightX++ {
				if !gameMap[y][rightX] {
					visitedMapCur[rightX] += startRightCount
					break
				}
			}
		}
		visitedMapPrev, visitedMapCur = visitedMapCur, visitedMapPrev

	}

	var timeSplits int
	for _, n := range visitedMapCur {
		timeSplits += n
	}
	return timeSplits
}

func main() {
	var gameMap [][]bool
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
	}

	fmt.Printf("Number of splits: %d\n", walkAndCountSplits(gameMap, start))
	fmt.Printf("Number of timelines: %d\n", walkTimeline(gameMap, start))
}

func walkAndCountSplits(gameMap [][]bool, start coordinate) int {
	var numberOfSplits int

	var visitedMap [][]bool
	for y := range gameMap {
		visitedMap = append(visitedMap, make([]bool, len(gameMap[y])))
	}

	queue := []coordinate{start}
	var c coordinate
	for len(queue) != 0 {
		c, queue = queue[0], queue[1:]
		if c.x < 0 || c.x == len(gameMap[0]) || c.y == len(gameMap) || visitedMap[c.y][c.x] {
			continue
		}
		visitedMap[c.y][c.x] = true
		if gameMap[c.y][c.x] {
			queue = append(queue, coordinate{c.x - 1, c.y})
			queue = append(queue, coordinate{c.x + 1, c.y})
			numberOfSplits++
			continue
		}
		queue = append(queue, coordinate{c.x, c.y + 1})
	}
	return numberOfSplits
}
