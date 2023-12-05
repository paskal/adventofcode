package main

import (
	_ "embed"
	"log"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	var sum int
	var sumWonScratchCards int
	var multipliers = map[int]int{}
	for i, card := range strings.Split(input, "\n") {
		if multipliers[i] == 0 {
			multipliers[i] = 1
		}
		sumWonScratchCards += multipliers[i]
		rawCard := strings.Split(card, ": ")
		if len(rawCard) != 2 {
			log.Panicf("card content is not separated by \":\" in two parts as expected: %d", len(rawCard))
		}
		card = rawCard[1]
		cardContent := strings.Split(card, " | ")
		if len(cardContent) != 2 {
			log.Panicf("card content is not separated by \"|\" in two parts as expected: %d", len(cardContent))
		}
		var winningNumbers = map[int]struct{}{}
		var currentNumber string
		isDigit := regexp.MustCompile(`\d`)
		for _, c := range strings.Split(cardContent[0], "") {
			if c != " " && isDigit.MatchString(c) {
				currentNumber += c
			}
			if c == " " && currentNumber != "" {
				num, _ := strconv.Atoi(currentNumber)
				winningNumbers[num] = struct{}{}
				currentNumber = ""
			}
		}
		if currentNumber != "" {
			num, _ := strconv.Atoi(currentNumber)
			winningNumbers[num] = struct{}{}
			currentNumber = ""
		}
		var sumCard int
		var wonCards int
		for _, c := range strings.Split(cardContent[1], "") {
			if c != " " && isDigit.MatchString(c) {
				currentNumber += c
			}
			if c == " " && currentNumber != "" {
				num, _ := strconv.Atoi(currentNumber)
				if _, ok := winningNumbers[num]; ok {
					if sumCard == 0 {
						sumCard++
					} else {
						sumCard *= 2
					}
					wonCards++
					wonCard := i + wonCards
					if multipliers[wonCard] == 0 {
						multipliers[wonCard] = 1
					}
					multipliers[wonCard] += multipliers[i]
				}
				currentNumber = ""
			}
		}
		if currentNumber != "" {
			num, _ := strconv.Atoi(currentNumber)
			if _, ok := winningNumbers[num]; ok {
				if sumCard == 0 {
					sumCard++
				} else {
					sumCard *= 2
				}
				wonCards++
				wonCard := i + wonCards
				if multipliers[wonCard] == 0 {
					multipliers[wonCard] = 1
				}
				multipliers[wonCard] += multipliers[i]
			}
			currentNumber = ""
		}
		sum += sumCard
	}
	log.Printf("Sum of winning cards: %d, won scratch cards: %d", sum, sumWonScratchCards)
}
