package main

import (
	_ "embed"
	"log"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	var limits = map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	var goodIDsSum int
	var powerSum int
	for i, s := range strings.Split(input, "\n") {
		games, _ := strings.CutPrefix(s, "Game "+strconv.Itoa(i+1)+": ")
		var gameIsBad bool
		var gameMaximum = map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		for _, game := range strings.Split(games, ";") {
			for _, results := range strings.Split(game, ",") {
				result := strings.Split(strings.Trim(results, " "), " ")
				if len(result) != 2 || limits[result[1]] == 0 {
					log.Panicf("unexpected game result, not digit and color: %q", result)
				}
				observedNum, _ := strconv.ParseInt(result[0], 10, 8)
				color := result[1]
				if int(observedNum) > limits[color] {
					gameIsBad = true
				}
				if gameMaximum[color] < int(observedNum) {
					gameMaximum[color] = int(observedNum)
				}
			}
		}
		if !gameIsBad {
			goodIDsSum += i + 1
		}
		powerSum += gameMaximum["red"] * gameMaximum["green"] * gameMaximum["blue"]
		if i < 10 {
			log.Printf("game %d is bad %v, sum after: %d", i+1, gameIsBad, goodIDsSum)
		}
	}
	log.Printf("Sum of the good games: %d, power sum: %d", goodIDsSum, powerSum)
}
