package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func collapseRanges(ranges map[int]int) map[int]int {
	newRanges := map[int]int{}
	keys := make([]int, 0)

	for beginning := range ranges {
		keys = append(keys, beginning)
	}
	slices.Sort(keys)

	for _, id := range keys {
		var found bool
		for beginning, end := range newRanges {
			if id >= beginning && id <= end {
				found = true
				if newRanges[beginning] < ranges[id] {
					newRanges[beginning] = ranges[id]
				}
				break
			}
		}
		if !found {
			newRanges[id] = ranges[id]
		}
	}

	return newRanges
}

func main() {
	ranges := map[int]int{} // beginning as key, ending as value
	var ids []int
	raw := strings.Split(input, "\n\n")
	for r := range strings.SplitSeq(raw[0], "\n") {
		interval := strings.Split(r, "-")
		beginning, _ := strconv.Atoi(interval[0])
		end, _ := strconv.Atoi(interval[1])
		if ranges[beginning] < end {
			ranges[beginning] = end
		}
	}
	for i := range strings.SplitSeq(raw[1], "\n") {
		id, _ := strconv.Atoi(i)
		ids = append(ids, id)
	}

	ranges = collapseRanges(ranges)

	var spoiledIngridients []int
	for _, id := range ids {
		var found bool
		for beginning, end := range ranges {
			if id >= beginning && id <= end {
				found = true
				break
			}
		}
		if !found {
			spoiledIngridients = append(spoiledIngridients, id)
		}
	}

	var goodIngridientsCount int
	for beginning, end := range ranges {
		goodIngridientsCount += end - beginning + 1
	}

	fmt.Printf("Fresh ingridients: %d out of %d with %d ranges\n", len(ids)-len(spoiledIngridients), len(ids), len(ranges))
	fmt.Printf("Good ingridients count: %d\n", goodIngridientsCount)
}
