package main

import (
	_ "embed"
	"log"
	"math"
	"slices"
	"strings"
)

//go:embed input.txt
var input string

type galaxy struct {
	x, y        int
	pathLengths []int
}

func main() {
	galaxies := getGalaxies(2 - 1)
	var totalLength int
	for i, source := range galaxies {
		for _, destination := range galaxies[i+1:] {
			length := findPathLength(source.x, source.y, destination.x, destination.y)
			galaxies[i].pathLengths = append(galaxies[i].pathLengths, length)
			totalLength += length
		}
	}
	log.Printf("length of shortest paths between all galaxies with expand 2x: %d", totalLength)
	galaxies = getGalaxies(1000000 - 1)
	totalLength = 0
	for i, source := range galaxies {
		for _, destination := range galaxies[i+1:] {
			length := findPathLength(source.x, source.y, destination.x, destination.y)
			galaxies[i].pathLengths = append(galaxies[i].pathLengths, length)
			totalLength += length
		}
	}
	log.Printf("length of shortest paths between all galaxies with expand 1000000x: %d", totalLength)

}

func getGalaxies(expandValue int) (galaxies []galaxy) {
	var filledRows, filledColumns []int
	for y, s := range strings.Split(input, "\n") {
		for x, c := range strings.Split(s, "") {
			if c == "#" {
				filledColumns = append(filledColumns, y)
				filledRows = append(filledRows, x)
				galaxies = append(galaxies, galaxy{x: x, y: y})
			}
		}
	}
	for i, _ := range galaxies {
		var incX int
		// no need to check current row as it has galaxy
		for x := 0; x < galaxies[i].x; x++ {
			if !slices.Contains(filledRows, x) {
				incX += expandValue
			}
		}
		// no need to check current column as it has galaxy
		var incY int
		for y := 0; y < galaxies[i].y; y++ {
			if !slices.Contains(filledColumns, y) {
				incY += expandValue
			}
		}
		galaxies[i].x += incX
		galaxies[i].y += incY
	}
	return galaxies
}

func findPathLength(x int, y int, x2 int, y2 int) (length int) {
	return int(math.Abs(float64(x2-x)) + math.Abs(float64(y2-y)))
}
