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
	"*": 1,
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
	var jokers int
	for _, c := range h.cards {
		if c == "*" {
			jokers++
			continue
		}
		handCount[c]++
	}
	count := map[int][]string{}
	for c, num := range handCount {
		count[num] = append(count[num], c)
	}
	if possibleWithJoker(5, jokers, count) >= 0 || jokers == 5 {
		h.handType = "five"
		return
	}
	if possibleWithJoker(4, jokers, count) >= 0 {
		h.handType = "four"
		return
	}
	if reducedJokers := possibleWithJoker(3, jokers, count); reducedJokers >= 0 {
		if possibleWithJoker(2, reducedJokers, count) >= 0 {
			h.handType = "house"
			return
		}
		h.handType = "three"
		return
	}
	if reducedJokers := possibleWithJoker(2, jokers, count); reducedJokers >= 0 {
		if len(count[2]) == 1 {
			h.handType = "two"
			return
		}
		h.handType = "one"
		return
	}
	h.handType = "high"
}

func possibleWithJoker(desired int, jokers int, count map[int][]string) int {
	for i := 0; i <= desired; i++ {
		if len(count[desired-i]) != 0 && jokers-i >= 0 {
			// as we have nested conditions, it's crucial to delete previously used letters
			// so that nested check won't count them again
			if len(count[desired-i]) == 1 {
				delete(count, desired-i)
			} else {
				count[desired-i] = count[desired-i][:len(count[desired-i])-1]
			}
			return jokers - i
		}
	}
	return -1
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
	var partTwoHands []*hand
	for _, s := range strings.Split(input, "\n") {
		rawStrings := strings.Split(s, " ")
		num, _ := strconv.Atoi(rawStrings[1])
		hands = append(hands, &hand{
			cards: strings.Split(rawStrings[0], ""),
			score: num,
		})
		rawStrings = strings.Split(strings.ReplaceAll(s, "J", "*"), " ")
		partTwoHands = append(partTwoHands, &hand{
			cards: strings.Split(rawStrings[0], ""),
			score: num,
		})
	}
	for _, h := range hands {
		h.calculateType()
	}
	for _, h := range partTwoHands {
		h.calculateType()
	}
	slices.SortFunc(hands, compareHands)
	slices.SortFunc(partTwoHands, compareHands)
	var winnings int
	for i, h := range hands {
		winnings += h.score * (i + 1)
	}
	var partTwoWinnings int
	for i, h := range partTwoHands {
		partTwoWinnings += h.score * (i + 1)
	}

	log.Printf("Resulting winnings %d, winnings with joker: %d", winnings, partTwoWinnings)
}
