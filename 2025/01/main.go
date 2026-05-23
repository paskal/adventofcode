package main

import (
	_ "embed"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	lineRegexp := regexp.MustCompile(`^([LR])(\d+)$`)
	var currentPosition = 50
	var ticks int64
	var numberOfZeroPositions int
	for s := range strings.SplitSeq(input, "\n") {
		r := lineRegexp.FindStringSubmatch(s)
		ticks, _ = strconv.ParseInt(r[2], 10, 64)
		if r[1] == "L" {
			currentPosition += int(math.Floor(float64(ticks)/100)) * 100 // append full circles to total for passing zero calculation
			ticks = int64(math.Mod(float64(ticks), 100))                 // drop full circles from ticks for clarity
			if int64(math.Mod(float64(currentPosition), 100)) != 0 {     // from 0, we won't pass 0 again no matter which direction we go.
				if int64(math.Mod(float64(currentPosition), 100))-ticks < 0 {
					currentPosition += 100
				}
				if int64(math.Mod(float64(currentPosition), 100))-ticks > 0 {
					currentPosition -= 100
				}
			}
			ticks = 100 - ticks
		}
		currentPosition += int(ticks)
		if math.Mod(float64(currentPosition), 100) == 0 {
			numberOfZeroPositions++
		}
	}
	log.Printf("Number of zero positions: %d", numberOfZeroPositions)
	log.Printf("Number of zero passes: %v", math.Floor(float64(currentPosition)/100))
}
