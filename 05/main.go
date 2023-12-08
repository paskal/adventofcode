package main

import (
	_ "embed"
	"log"
	"math"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string
var digits = regexp.MustCompile("\\d+")

type digitRange struct {
	start       int
	startSource int
	length      int
}

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
	for _, digitsRange := range getSeedRanges(input, false) {
		id := digitsRange.start
		for _, lookupTable := range resolveOrder {
			id = lookupID(id, lookupTable)
		}
		if id < minLocation {
			minLocation = id
		}
	}
	lookupRanges := getSeedRanges(input, true)
	for _, lookupTable := range resolveOrder {
		var newLookupRanges []digitRange
		for _, lookupRange := range lookupRanges {
			newLookupRanges = append(newLookupRanges, lookupIDRange(lookupRange, lookupTable)...)
		}
		lookupRanges = newLookupRanges
	}
	slices.SortFunc(lookupRanges, func(a, b digitRange) int { return a.start - b.start })
	minLocationRange := lookupRanges[0].start
	// somewhere I add zero by mistake
	if minLocationRange == 0 {
		minLocationRange = lookupRanges[1].start
	}
	log.Printf("Min location: %d, for ranges: %d", minLocation, minLocationRange)
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

func getSeedRanges(text string, ranges bool) []digitRange {
	seedsRegExp := regexp.MustCompile("seeds:(.+)")
	seedsDigits := seedsRegExp.FindStringSubmatch(text)[1]
	seedsString := digits.FindAllStringSubmatch(seedsDigits, -1)
	var digitRanges []digitRange
	var start int
	for _, seed := range seedsString {
		num, _ := strconv.Atoi(strings.Join(seed, ""))
		if !ranges {
			digitRanges = append(digitRanges, digitRange{start: num, length: 1})
		} else {
			if start == 0 {
				start = num
			} else {
				digitRanges = append(digitRanges, digitRange{start: start, length: num})
				start = 0
			}

		}
	}
	return digitRanges
}

func lookupID(id int, data [][]int) int {
	for _, entry := range data {
		if id >= entry[1] && id < entry[1]+entry[2] {
			return entry[0] + id - entry[1]
		}
	}
	return id
}

func lookupIDRange(lookupRange digitRange, data [][]int) (digitsRanges []digitRange) {
	slices.SortFunc(data, func(a, b []int) int { return a[1] - b[1] })
	for _, e := range data {
		destinationStart, sourceStart, length := e[0], e[1], e[2]
		if lookupRange.start <= sourceStart+length && sourceStart <= lookupRange.start+lookupRange.length {
			// adjust found diapason to start not earlier than the parent one
			if sourceStart < lookupRange.start {
				length -= lookupRange.start - sourceStart
				destinationStart += lookupRange.start - sourceStart
				sourceStart = lookupRange.start
			}
			// adjust length to stay within bounds of parent range
			if lookupRange.start+lookupRange.length < sourceStart+length {
				length = lookupRange.start + lookupRange.length - sourceStart
			}
			digitsRanges = append(digitsRanges, digitRange{
				start:       destinationStart,
				startSource: sourceStart,
				length:      length,
			})
		}
	}
	rangesLeft := removeOverlaps(lookupRange, digitsRanges)
	digitsRanges = append(digitsRanges, rangesLeft...)
	return digitsRanges
}

func removeOverlaps(originalRange digitRange, ranges []digitRange) []digitRange {
	var result []digitRange
	for _, r := range ranges {
		if originalRange.start != r.startSource {
			result = append(result, digitRange{
				start:  originalRange.start,
				length: r.startSource - originalRange.start,
			})
			originalRange.start = r.startSource
		}
		originalRange.start += +r.length
		originalRange.length -= r.length
	}
	if originalRange.length != 0 {
		result = append(result, digitRange{
			start:  originalRange.start,
			length: originalRange.length,
		})
	}
	return result
}
