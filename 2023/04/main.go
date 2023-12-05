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
	for i, card := range strings.Split(input, "\n") {
		card, _ = strings.CutPrefix(card, "Card "+strconv.Itoa(i+1)+": ")
		cardContent := strings.Split(card, " | ")
		if len(cardContent) != 2 {
			log.Panicf("card content is not separated by \"|\" in two parts as expected: %d", len(cardContent))
		}
		var chosenNumbers = map[int]struct{}{}
		var currentNumber string
		isDigit := regexp.MustCompile(`\d`)
		for _, c := range strings.Split(cardContent[0], "") {
			if c != " " && isDigit.MatchString(c) {
				currentNumber += c
			}
			if c == " " && currentNumber != "" {
				num, _ := strconv.Atoi(currentNumber)
				chosenNumbers[num] = struct{}{}
				currentNumber = ""
			}
		}
		if currentNumber != "" {
			num, _ := strconv.Atoi(currentNumber)
			chosenNumbers[num] = struct{}{}
			currentNumber = ""
		}
		var sumCard int
		for _, c := range strings.Split(cardContent[1], "") {
			if c != " " && isDigit.MatchString(c) {
				currentNumber += c
			}
			if c == " " && currentNumber != "" {
				num, _ := strconv.Atoi(currentNumber)
				if _, ok := chosenNumbers[num]; ok {
					if sumCard == 0 {
						sumCard++
					} else {
						sumCard *= 2
					}
				}
				currentNumber = ""
			}
		}
		if currentNumber != "" {
			num, _ := strconv.Atoi(currentNumber)
			if _, ok := chosenNumbers[num]; ok {
				if sumCard == 0 {
					sumCard++
				} else {
					sumCard *= 2
				}
			}
			currentNumber = ""
		}
		sum += sumCard
	}
	log.Printf("Sum of winning cards: %d", sum)
}
