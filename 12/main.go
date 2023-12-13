package main

import (
	_ "embed"
	"log"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type row struct {
	values  []string
	damaged []int
}

func main() {
	rows := getRows()
	var totalPermutations int
	for _, r := range rows {
		totalPermutations += countPermutations(strings.Join(r.values, ""), r.damaged)
	}
	log.Printf("total permutations: %d", totalPermutations)
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
	if len(beforeQuestionMark) == 1 {
		return slices.Equal(damaged, valuesDamagedPattern)
	}
	// partial match check
	for i, patternValue := range valuesDamagedPattern {
		if len(damaged) <= i || damaged[i] < patternValue {
			return false
		}
	}
	return true
}

func getRows() (rows []row) {
	digitsRegexp := regexp.MustCompile(`\d+`)
	for _, s := range strings.Split(input, "\n") {
		raw := strings.Split(s, " ")
		if len(raw) != 2 {
			log.Panicf("unexpected input, length is not 2 but %d", len(raw))
		}
		rawValues, rawDigits := raw[0], raw[1]
		newRow := row{}
		for _, c := range strings.Split(rawValues, "") {
			newRow.values = append(newRow.values, c)
		}
		for _, d := range digitsRegexp.FindAllStringSubmatch(rawDigits, -1) {
			num, _ := strconv.Atoi(d[0])
			newRow.damaged = append(newRow.damaged, num)
		}
		rows = append(rows, newRow)
	}
	return rows
}
