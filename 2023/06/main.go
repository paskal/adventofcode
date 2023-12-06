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
	sumWinningOutcomesMultipleRaces := multipleRacesWinningOutcomes(rawTime[1], rawDistance[1])
	log.Printf("Winning outcomes with multiple races: %d", sumWinningOutcomesMultipleRaces)
	sumWinningOutcomesSingleRace := oneRaceWinningOutcomes(rawTime[1], rawDistance[1])
	log.Printf("Winning outcomes with single race: %d", sumWinningOutcomesSingleRace)

}

func multipleRacesWinningOutcomes(rawTime string, rawDistance string) int {
	var races []raceEntry
	var currentNumber string
	isDigit := regexp.MustCompile(`\d`)
	// add space to verify last digit will be processed within the loop
	for _, c := range strings.Split(rawTime+" ", "") {
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
	for _, c := range strings.Split(rawDistance+" ", "") {
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
		winningOutcomes := getWinningOutcomes(race)
		if sumWinningOutcomes == 0 {
			sumWinningOutcomes += winningOutcomes
		} else {
			sumWinningOutcomes *= winningOutcomes
		}
	}
	return sumWinningOutcomes
}

func oneRaceWinningOutcomes(rawTime string, rawDistance string) int {
	var race raceEntry
	var currentNumber string
	isDigit := regexp.MustCompile(`\d`)
	for _, c := range strings.Split(rawTime, "") {
		if c != " " && isDigit.MatchString(c) {
			currentNumber += c
		}
	}
	num, _ := strconv.Atoi(currentNumber)
	race.time = num
	currentNumber = ""
	for _, c := range strings.Split(rawDistance, "") {
		if c != " " && isDigit.MatchString(c) {
			currentNumber += c
		}
	}
	num, _ = strconv.Atoi(currentNumber)
	race.distance = num
	return getWinningOutcomes(race)
}

func getWinningOutcomes(race raceEntry) int {
	var winningOutcomes int
	for buttonPressTime := 0; buttonPressTime <= race.time; buttonPressTime++ {
		if buttonPressTime*(race.time-buttonPressTime) > race.distance {
			winningOutcomes += 1
		}
		// formula: buttonPressTime*buttonPressTime-race.time*buttonPressTime+race.distance < 0
	}
	return winningOutcomes
}
