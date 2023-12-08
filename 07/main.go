package main

import (
	_ "embed"
	"log"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

var cardValue = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"J": 11,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
}

var handValues = map[string]int{
	"five":  7,
	"four":  6,
	"house": 5,
	"three": 4,
	"two":   3,
	"one":   2,
	"high":  1,
}

type hand struct {
	cards    []string
	score    int
	handType string
}

func (h *hand) calculateType() {
	handCount := map[string]int{}
	for _, c := range h.cards {
		handCount[c]++
	}
	count := map[int][]string{}
	for c, num := range handCount {
		count[num] = append(count[num], c)
	}
	if _, ok := count[5]; ok {
		h.handType = "five"
		return
	}
	if _, ok := count[4]; ok {
		h.handType = "four"
		return
	}
	if _, threeOk := count[3]; threeOk {
		if _, twoOk := count[2]; twoOk {
			h.handType = "house"
			return
		}
		h.handType = "three"
		return
	}
	if _, twoOk := count[2]; twoOk {
		if len(count[2]) == 2 {
			h.handType = "two"
			return
		}
		h.handType = "one"
		return
	}
	h.handType = "high"
}

func compareHands(a, b *hand) int {
	if a.handType == b.handType {
		for i := 0; i < len(a.cards); i++ {
			if a.cards[i] != b.cards[i] {
				return cardValue[a.cards[i]] - cardValue[b.cards[i]]
			}
		}
	}
	return handValues[a.handType] - handValues[b.handType]
}

func main() {
	var hands []*hand
	for _, s := range strings.Split(input, "\n") {
		rawStrings := strings.Split(s, " ")
		num, _ := strconv.Atoi(rawStrings[1])
		hands = append(hands, &hand{
			cards: strings.Split(rawStrings[0], ""),
			score: num,
		})
	}
	for _, h := range hands {
		h.calculateType()
	}
	slices.SortFunc(hands, compareHands)
	var winnings int
	for i, h := range hands {
		winnings += h.score * (i + 1)
	}
	log.Printf("Resulting winnings %d", winnings)
}
