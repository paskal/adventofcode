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
	offset      int
}

func main() {
	resolveOrder := [][]digitRange{
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

func collectMatches(matchString string) []digitRange {
	var results []digitRange
	findSectionRegexp := regexp.MustCompile(matchString + " map:\n(?:\\d+ \\d+ \\d+\n?)+")
	sectionRaw := findSectionRegexp.FindStringSubmatch(input)[0]
	digitsRegexp := regexp.MustCompile("(\\d+) (\\d+) (\\d+)")
	for _, s := range strings.Split(sectionRaw, "\n") {
		match := digitsRegexp.FindStringSubmatch(s)
		if len(match) == 4 {
			destinationRangeStart, _ := strconv.Atoi(match[1])
			sourceRangeStart, _ := strconv.Atoi(match[2])
			rangeLength, _ := strconv.Atoi(match[3])
			results = append(results, digitRange{
				start:  sourceRangeStart,
				length: rangeLength,
				offset: destinationRangeStart - sourceRangeStart,
			})
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

func lookupID(id int, data []digitRange) int {
	for _, e := range data {
		if id >= e.start && id < e.start+e.length {
			return id + e.offset
		}
	}
	return id
}

func lookupIDRange(lookupRange digitRange, data []digitRange) (digitsRanges []digitRange) {
	slices.SortFunc(data, func(a, b digitRange) int { return a.start - b.start })
	for _, e := range data {
		if lookupRange.start <= e.start+e.length && e.start <= lookupRange.start+lookupRange.length {
			// adjust found range to start not earlier than the parent one
			if e.start < lookupRange.start {
				e.length -= lookupRange.start - e.start
				e.start = lookupRange.start
			}
			// adjust length to stay within bounds of parent range
			if lookupRange.start+lookupRange.length < e.start+e.length {
				e.length = lookupRange.start + lookupRange.length - e.start
			}
			digitsRanges = append(digitsRanges, digitRange{
				start:       e.start + e.offset,
				startSource: e.start,
				length:      e.length,
			})
		}
	}
	rangesLeft := removeOverlaps(lookupRange, digitsRanges)
	digitsRanges = append(digitsRanges, rangesLeft...)
	return digitsRanges
}

func removeOverlaps(lookupRange digitRange, ranges []digitRange) []digitRange {
	var result []digitRange
	for _, r := range ranges {
		if lookupRange.start != r.startSource {
			result = append(result, digitRange{
				start:  lookupRange.start,
				length: r.startSource - lookupRange.start,
			})
			lookupRange.start = r.startSource
		}
		lookupRange.start += +r.length
		lookupRange.length -= r.length
	}
	if lookupRange.length != 0 {
		result = append(result, digitRange{
			start:  lookupRange.start,
			length: lookupRange.length,
		})
	}
	return result
}
