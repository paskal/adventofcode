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
	resolveOrder := [][][]int{
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
		for _, lookupTable := range resolveOrder {
			id = lookupID(id, lookupTable)
		}
		if id < minLocation {
			minLocation = id
		}
	}
	log.Printf("Min location: %d", minLocation)
}

func collectMatches(matchString string) [][]int {
	var results = [][]int{}
	findSectionRegexp := regexp.MustCompile(matchString + " map:\n(?:\\d+ \\d+ \\d+\n?)+")
	sectionRaw := findSectionRegexp.FindStringSubmatch(input)[0]
	digitsRegexp := regexp.MustCompile("(\\d+) (\\d+) (\\d+)")
	for _, s := range strings.Split(sectionRaw, "\n") {
		match := digitsRegexp.FindStringSubmatch(s)
		if len(match) == 4 {
			destinationRangeStart, _ := strconv.Atoi(match[1])
			sourceRangeStart, _ := strconv.Atoi(match[2])
			rangeLength, _ := strconv.Atoi(match[3])
			results = append(results, []int{destinationRangeStart, sourceRangeStart, rangeLength})
		}
	}
	return results
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

func lookupID(id int, data [][]int) int {
	for _, entry := range data {
		if id > entry[1] && id < entry[1]+entry[2] {
			return entry[0] + id - entry[1]
		}
	}
	return id
}
