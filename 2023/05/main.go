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
	start  int
	length int
	offset int
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
		for _, overrides := range resolveOrder {
			id = lookupID(id, overrides)
		}
		// after going through resolveOrder, id is location id
		if id < minLocation {
			minLocation = id
		}
	}
	lookupRanges := getSeedRanges(input, true)
	for _, overrides := range resolveOrder {
		var newLookupRanges []digitRange
		for _, lookupRange := range lookupRanges {
			newLookupRanges = append(newLookupRanges, lookupIDRange(lookupRange, overrides)...)
		}
		lookupRanges = newLookupRanges
	}
	// after going through resolveOrder, lookupRanges contain locations
	slices.SortFunc(lookupRanges, func(a, b digitRange) int { return a.start - b.start })
	minLocationRange := lookupRanges[0].start
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

func lookupIDRange(lookupRange digitRange, overrides []digitRange) (digitsRanges []digitRange) {
	slices.SortFunc(overrides, func(a, b digitRange) int { return a.start - b.start })
	for _, e := range overrides {
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
			// move the original range start to the start of the current offset
			if lookupRange.start != e.start {
				digitsRanges = append(digitsRanges, digitRange{
					start:  lookupRange.start,
					length: e.start - lookupRange.start,
				})
				lookupRange.start = e.start - lookupRange.start
			}
			// add found range to the result
			digitsRanges = append(digitsRanges, digitRange{
				start:  e.start + e.offset,
				length: e.length,
			})
			// cut the part of lookupRange we will just found in overrides
			lookupRange.start += +e.length
			lookupRange.length -= e.length
		}
	}
	// if some lookup range is left after finding overwrites, add it to the result
	if lookupRange.length != 0 {
		digitsRanges = append(digitsRanges, digitRange{
			start:  lookupRange.start,
			length: lookupRange.length,
		})
	}
	return digitsRanges
}
