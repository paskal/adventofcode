package main

import (
	_ "embed"
	"log"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

//go:embed input.txt
var input string

type row struct {
	values  []string
	damaged []int
}

func main() {
	rowsOne := getRows(1)
	var totalPermutationsOne atomic.Int32
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(len(rowsOne))
	for _, r := range rowsOne {
		go func(r row) {
			totalPermutationsOne.Add(int32(countPermutations(strings.Join(r.values, ""), r.damaged)))
			waitGroup.Done()
		}(r)
	}
	waitGroup.Wait()
	log.Printf("total permutations for multiplier 1: %d", totalPermutationsOne.Load())

	rowsFive := getRows(5)
	var totalPermutationsFive atomic.Int32
	waitGroup.Add(len(rowsFive))
	for _, r := range rowsFive {
		go func(r row) {
			totalPermutationsFive.Add(int32(countPermutations(strings.Join(r.values, ""), r.damaged)))
			waitGroup.Done()
		}(r)
	}
	waitGroup.Wait()
	log.Printf("total permutations for multiplier 5: %d", totalPermutationsFive.Load())
}

func countPermutations(values string, damagedPattern []int) (total int) {
	damagedRegexp := regexp.MustCompile(`#+`)
	for i, c := range strings.Split(values, "") {
		if c == "?" {
			var valueStart string
			if i > 0 {
				valueStart = strings.Join(strings.Split(values, "")[:i], "")
			}
			valueEnd := strings.Join(strings.Split(values, "")[i+1:], "")
			var hashTagPattern []int
			for _, match := range damagedRegexp.FindAllStringSubmatch(valueStart+"#", -1) {
				hashTagPattern = append(hashTagPattern, len(match[0]))
			}
			tryHashTag := valueStart + "#" + valueEnd
			if patternMatches(hashTagPattern, damagedPattern, false) {
				total += countPermutations(tryHashTag, damagedPattern)
			}
			// reduce hashtag count by one for trying the dot
			if len(hashTagPattern) > 0 {
				hashTagPattern[len(hashTagPattern)-1]--
				if hashTagPattern[len(hashTagPattern)-1] == 0 {
					hashTagPattern = hashTagPattern[:len(hashTagPattern)-1]
				}
			}
			tryDot := valueStart + "." + valueEnd
			if patternMatches(hashTagPattern, damagedPattern, false) {
				total += countPermutations(tryDot, damagedPattern)
			}
			break
		}
		if i == len(values)-1 {
			var fullPattern []int
			for _, match := range damagedRegexp.FindAllStringSubmatch(values, -1) {
				fullPattern = append(fullPattern, len(match[0]))
			}
			if patternMatches(fullPattern, damagedPattern, true) {
				total += 1
			}
		}
	}
	return total
}

func patternMatches(checkPattern []int, damaged []int, full bool) bool {
	// length of what we got is bigger than damaged pattern
	if len(damaged) < len(checkPattern) {
		return false
	}
	// case of full match should be checked completely
	if full && len(checkPattern) != len(damaged) {
		return false
	}
	// partial match check, only for last two items as previous are checked before
	for i, patternValue := range checkPattern {
		if len(checkPattern) > 3 && i < 3 {

		}
		// 1. damaged value for given index is lower than discovered pattern value
		// 2. it's not the last element of the list of patterns, and damaged value doesn't completely match discovered pattern value
		// (2. is more strict check than 1. only for cases when we know number won't increase)
		// 3. last item check if no question marks left
		if damaged[i] < patternValue || (i < len(checkPattern)-1 && damaged[i] != patternValue) || (full && i == len(checkPattern)-1 && damaged[i] != patternValue) {
			return false
		}
	}
	return true
}

func getRows(multiplier int) (rows []row) {
	digitsRegexp := regexp.MustCompile(`\d+`)
	for _, s := range strings.Split(input, "\n") {
		raw := strings.Split(s, " ")
		if len(raw) != 2 {
			log.Panicf("unexpected input, length is not 2 but %d", len(raw))
		}
		multipliedRawValues, multipliedRawDigits := raw[0], raw[1]
		for i := 1; i < multiplier; i++ {
			multipliedRawValues += "?" + raw[0]
			multipliedRawDigits += "," + raw[1]
		}
		newRow := row{}
		for _, c := range strings.Split(multipliedRawValues, "") {
			newRow.values = append(newRow.values, c)
		}
		for _, d := range digitsRegexp.FindAllStringSubmatch(multipliedRawDigits, -1) {
			num, _ := strconv.Atoi(d[0])
			newRow.damaged = append(newRow.damaged, num)
		}
		rows = append(rows, newRow)
	}
	return rows
}
