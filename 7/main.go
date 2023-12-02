package main

import (
	_ "embed"
	"log"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	for i, s := range strings.Split(input, "\n") {
		log.Printf("line %d: %q", i+1, s)
	}
}
