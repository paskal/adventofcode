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
	rowsFive := getRows(5)
	var totalPermutationsOne atomic.Int32
	var totalPermutationsFive atomic.Int32
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

func countPermutations(values string, damaged []int) (total int) {
	for i, c := range strings.Split(values, "") {
		if c == "?" {
			var valueStart string
			if i > 0 {
				valueStart = strings.Join(strings.Split(values, "")[:i], "")
			}
			valueEnd := strings.Join(strings.Split(values, "")[i+1:], "")
			tryHashTag := valueStart + "#" + valueEnd
			tryDot := valueStart + "." + valueEnd
			if patternMatches(tryHashTag, damaged) {
				total += countPermutations(tryHashTag, damaged)
			}
			if patternMatches(tryDot, damaged) {
				total += countPermutations(tryDot, damaged)
			}
			break
		}
		if i == len(values)-1 && patternMatches(values, damaged) {
			total += 1
		}
	}
	return total
}

func patternMatches(values string, damaged []int) bool {
	var valuesDamagedPattern []int
	beforeQuestionMark := strings.Split(values, "?")
	damagedRegexp := regexp.MustCompile(`#+`)
	for _, c := range damagedRegexp.FindAllStringSubmatch(beforeQuestionMark[0], -1) {
		valuesDamagedPattern = append(valuesDamagedPattern, len(c[0]))
	}
	// case of full match should be checked completely
	if len(beforeQuestionMark) == 1 && len(valuesDamagedPattern) != len(damaged) {
		return false
	}
	// partial match check
	for i, patternValue := range valuesDamagedPattern {
		// 1. length of damaged is lower than resulting pattern length
		// 2. damaged value for given index is lower than discovered pattern value
		// 3. it's not the last element of the list of patterns, and damaged value doesn't completely match discovered pattern value
		// (3. is more strict check than 2. only for cases when we know number won't increase)
		// 4. last item check if no question marks left
		if len(damaged) <= i || damaged[i] < patternValue || (i < len(valuesDamagedPattern)-1 && damaged[i] != patternValue) || (len(beforeQuestionMark) == 1 && i == len(valuesDamagedPattern)-1 && damaged[i] != patternValue) {
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
