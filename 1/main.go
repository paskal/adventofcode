package main

import (
	_ "embed"
	"log"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

var digits = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func main() {
	var sum int64
	for _, s := range strings.Split(input, "\n") {
		var first, last int64
		for i, r := range strings.Split(s, "") {
			if n, err := strconv.ParseInt(r, 10, 8); err == nil {
				if first == 0 {
					first = n
				}
				last = n
				continue
			}
			for j, digit := range digits {
				if len(s[i:]) >= len(digit) && s[i:i+len(digit)] == digit {
					//log.Printf("found %s in %q", digit, s[i:])
					if first == 0 {
						first = int64(j) + 1
					}
					last = int64(j) + 1
					break
				}
			}
		}
		sum += first*10 + last
		if sum < 100 {
			log.Printf("sum for %q is %d", s, first*10+last)
		}
	}
	log.Printf("Sum: %d", sum)
}
