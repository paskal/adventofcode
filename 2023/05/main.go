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
var digits = regexp.MustCompile("\\d+")

func main() {
	resolveOrder := []map[int]int{
		collectMatches("seed-to-soil"),
		collectMatches("soil-to-fertilizer"),
		collectMatches("fertilizer-to-water"),
		collectMatches("water-to-light"),
		collectMatches("light-to-temperature"),
		collectMatches("temperature-to-humidity"),
		collectMatches("humidity-to-location"),
	}
	minLocation := int(math.Inf(0))
	for _, id := range getSeeds(input) {
		for _, lookupMap := range resolveOrder {
			id = lookupID(id, lookupMap)
		}
		if id < minLocation {
			minLocation = id
		}
	}
	log.Printf("Min location: %d", minLocation)
}

func collectMatches(matchString string) map[int]int {
	var mapResults = map[int]int{}
	findSectionRegexp := regexp.MustCompile(matchString + " map:\n(?:\\d+ \\d+ \\d+\n?)+")
	sectionRaw := findSectionRegexp.FindStringSubmatch(input)[0]
	digitsRegexp := regexp.MustCompile("(\\d+) (\\d+) (\\d+)")
	for _, s := range strings.Split(sectionRaw, "\n") {
		match := digitsRegexp.FindStringSubmatch(s)
		if len(match) == 4 {
			destinationRangeStart, _ := strconv.Atoi(match[1])
			sourceRangeStart, _ := strconv.Atoi(match[2])
			rangeLength, _ := strconv.Atoi(match[3])
			for i := 0; i < rangeLength; i++ {
				mapResults[sourceRangeStart] = destinationRangeStart
				sourceRangeStart++
				destinationRangeStart++
			}
		}
	}
	return mapResults
}

func getSeeds(text string) []int {
	seedsRegExp := regexp.MustCompile("seeds:(.+)")
	seedsDigits := seedsRegExp.FindStringSubmatch(text)[1]
	seedsString := digits.FindAllStringSubmatch(seedsDigits, -1)
	var seeds []int
	for _, seed := range seedsString {
		num, _ := strconv.Atoi(strings.Join(seed, ""))
		seeds = append(seeds, num)
	}
	return seeds
}

func lookupID(id int, data map[int]int) int {
	if value, ok := data[id]; ok {
		return value
	}
	return id
}
