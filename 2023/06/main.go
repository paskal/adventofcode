package main

import (
	_ "embed"
	"log"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type raceEntry struct {
	time     int
	distance int
}

func main() {
	rawLines := strings.Split(input, "\n")
	if len(rawLines) != 2 {
		log.Panicf("format is expected to have just two lines, got %d", len(rawLines))
	}
	rawTime := strings.Split(rawLines[0], ":")
	if len(rawTime) != 2 {
		log.Panicf("time line is not separated by \":\" in two parts as expected: %d", len(rawTime))
	}
	rawDistance := strings.Split(rawLines[1], ":")
	if len(rawDistance) != 2 {
		log.Panicf("distance line is not separated by \":\" in two parts as expected: %d", len(rawDistance))
	}
	var races []raceEntry
	var currentNumber string
	isDigit := regexp.MustCompile(`\d`)
	// add space to verify last digit will be processed within the loop
	for _, c := range strings.Split(rawTime[1]+" ", "") {
		if c != " " && isDigit.MatchString(c) {
			currentNumber += c
		}
		if c == " " && currentNumber != "" {
			num, _ := strconv.Atoi(currentNumber)
			races = append(races, raceEntry{time: num})
			currentNumber = ""
		}
	}
	// add space to verify last digit will be processed within the loop
	var digitIndex int
	for _, c := range strings.Split(rawDistance[1]+" ", "") {
		if c != " " && isDigit.MatchString(c) {
			currentNumber += c
		}
		if c == " " && currentNumber != "" {
			num, _ := strconv.Atoi(currentNumber)
			races[digitIndex].distance = num
			digitIndex++
			currentNumber = ""
		}
	}
	var sumWinningOutcomes int
	for _, race := range races {
		var winningOutcome int
		for time := 0; time <= race.time; time++ {
			if time*(race.time-time) > race.distance {
				winningOutcome += 1
			}
		}
		if sumWinningOutcomes == 0 {
			sumWinningOutcomes += winningOutcome
		} else {
			sumWinningOutcomes *= winningOutcome
		}
	}
	log.Printf("Winning outcomes: %d", sumWinningOutcomes)
}
