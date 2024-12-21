package main

import (
	_ "embed"
	"log"
	"math"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type queueItem struct {
	x, y, stepsInDirection int
	direction              string
}

type location [][]int

func (l location) findPathLength(sourceX, sourceY, destinationX, destinationY int) int {
	var distanceFromStart [][]int
	distanceFromStart[sourceX][sourceY] = 0
	for y, _ := range l {
		for x, _ := range l {
			if x != sourceX && y != sourceY {
				l[y][x] = int(math.Inf(0))

			}
		}
	}

	queue := []queueItem{{x: sourceX, y: sourceY, stepsInDirection: 0, direction: ""}}
	var current queueItem

	for len(queue) != 0 {

		current = queue[len(queue)-1]
		queue = queue[:len(queue)-1]

		for _, i := range l.neighbours(current.x, current.y) {
			x, y := i[0], i[1]

		}
		//for each neighbor u of v:           // where neighbor u has not yet been removed from Q.
		//	  alt := dist[v] + length(v, u)
		//	  if alt < dist[u]:               // A shorter path to u has been found
		//		  dist[u]  := alt            // Update distance of u
		//
		//  return dist[]

	}
	return distanceFromStart[destinationX][destinationY]
}

func (l location) neighbours(x int, y int) [][]int {
	var neighbours [][]int
	if x > 0 {
		if y > 0 {
			neighbours = append(neighbours, []int{x - 1, y - 1})
		}
		neighbours = append(neighbours, []int{x - 1, y})
		if y < len(l.distanceFromStart) {
		}
	}
	return neighbours
}

func main() {
	locationMap := getMap(input)
	shortestPathLength := locationMap.findPathLength(0, 0, len(locationMap)-1, len(locationMap[0])-1)
	log.Printf("Shortest length from start to finish: %d", shortestPathLength)
}

func getMap(input string) (locationMap location) {
	for y, s := range strings.Split(input, "\n") {
		row := strings.Split(s, "")
		locationMap = append(locationMap, make([]int, len(row)))
		for x, c := range row {
			num, _ := strconv.Atoi(c)
			locationMap[y][x] = num
		}
	}
	return locationMap
}
