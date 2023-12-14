package main

import (
	_ "embed"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/go-pkgz/lcw"
)

//go:embed input.txt
var input string

type row struct {
	values string
	groups string
}

type cache struct {
	cache lcw.LoadingCache
}

func main() {
	waitGroup := sync.WaitGroup{}
	c := cache{}
	c.cache, _ = lcw.NewLruCache()
	defer c.cache.Close()

	oneRow := getRows(1)
	var totalPermutationsOne atomic.Uint64
	waitGroup.Add(len(oneRow))
	for _, r := range oneRow {
		go func(r row) {
			totalPermutationsOne.Add(uint64(c.countPermutations(r.values, r.groups)))
			waitGroup.Done()
		}(r)
	}
	waitGroup.Wait()
	log.Printf("total permutations for multiplier 1: %d", totalPermutationsOne.Load())

	fiveRows := getRows(5)
	var totalPermutationsFive atomic.Uint64
	waitGroup.Add(len(fiveRows))
	for _, r := range fiveRows {
		go func(r row) {
			totalPermutationsFive.Add(uint64(c.countPermutations(r.values, r.groups)))
			waitGroup.Done()
		}(r)
	}
	waitGroup.Wait()
	log.Printf("total permutations for multiplier 5: %d", totalPermutationsFive.Load())
}

// that solution is originated from Reddit user u/StaticMoose, I didn't come up with it myself
// https://www.reddit.com/r/adventofcode/comments/18hbbxe/2023_day_12python_stepbystep_tutorial_with_bonus/
func (c cache) countPermutations(values string, rawGroups string) (total int) {
	// no groups but possible still more values to check
	if rawGroups == "" {
		if strings.Contains(values, "#") {
			return 0
		}
		return 1
	}

	// no values but groups still present
	if values == "" {
		return 0
	}

	switch strings.Split(values, "")[0] {
	case "#":
		return c.pound(values, rawGroups)
	case ".":
		return c.dot(values, rawGroups)
	}
	// case "?"
	return c.pound(values, rawGroups) + c.dot(values, rawGroups)
}

func (c cache) pound(record, rawGroups string) int {
	//cachedGroups, _ := c.cache.Get(rawGroups, func() (interface{}, error) {
	//	return parseGroups(rawGroups), nil
	//})
	//groups := cachedGroups.([]int)
	groups := parseGroups(rawGroups)

	nextGroup := groups[0]

	if len(record) < nextGroup {
		return 0
	}

	recordArray := strings.Split(record, "")
	// if the first is a pound, then the first N characters must be
	// able to be treated as a pound, where N is the first group number
	thisGroup := strings.Join(recordArray[:nextGroup], "")
	thisGroup = strings.ReplaceAll(thisGroup, "?", "#")

	// if the next group can't fit all the damaged springs, abort
	if thisGroup != strings.Repeat("#", nextGroup) {
		return 0
	}

	// if the rest of the record is just the last group, then we're
	// done and there's only one possibility
	if len(record) == nextGroup {
		if len(groups) == 1 {
			return 1
		}
		return 0
	}

	// make sure the character that follows this group can be a separator
	if recordArray[nextGroup] == "." || recordArray[nextGroup] == "?" {
		record = strings.Join(recordArray[nextGroup+1:], "")
		rawGroups = strings.Trim(strings.Join(strings.Fields(fmt.Sprint(groups[1:])), ","), "[]")
		cachedTotal, _ := c.cache.Get(record+rawGroups, func() (interface{}, error) {
			return c.countPermutations(record, rawGroups), nil
		})
		return cachedTotal.(int)
	}

	return 0
}

// just move on to next entry as dot doesn't alter groups
func (c cache) dot(values, rawGroups string) int {
	values = strings.Join(strings.Split(values, "")[1:], "")
	cachedTotal, _ := c.cache.Get(values+rawGroups, func() (interface{}, error) {
		return c.countPermutations(values, rawGroups), nil
	})
	return cachedTotal.(int)
}

func parseGroups(rawGroups string) []int {
	var groupsPattern []int
	digitsRegexp := regexp.MustCompile(`\d+`)
	for _, d := range digitsRegexp.FindAllStringSubmatch(rawGroups, -1) {
		num, _ := strconv.Atoi(d[0])
		groupsPattern = append(groupsPattern, num)
	}
	return groupsPattern
}

func getRows(multiplier int) (rows []row) {
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
		rows = append(rows, row{
			values: multipliedRawValues,
			groups: multipliedRawDigits,
		})
	}
	return rows
}
