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

type lens struct {
	label string
	value int
}

type hashMap [][]lens

func (h hashMap) add(s string, v int) {
	hash := calculateHashSum(s)
	if h[hash] == nil {
		h[hash] = []lens{}
	}
	if id := slices.IndexFunc(h[hash], func(c lens) bool { return c.label == s }); id != -1 {
		h[hash][id].value = v
		return
	}
	h[hash] = append(h[hash], lens{label: s, value: v})
}

func (h hashMap) delete(s string) {
	hash := calculateHashSum(s)
	if h[hash] == nil {
		return
	}
	if id := slices.IndexFunc(h[hash], func(c lens) bool { return c.label == s }); id != -1 {
		h[hash] = append(h[hash][:id], h[hash][id+1:]...)
	}
}

func (h hashMap) getFocusingPower() int {
	var result int
	for i, _ := range h {
		for j, lens := range h[i] {
			result += (i + 1) * (j + 1) * lens.value
		}
	}
	return result
}

func main() {
	var result int
	for _, row := range strings.Split(input, ",") {
		result += calculateHashSum(row)
	}
	log.Printf("Sum of hashes for give input: %d", result)
	lensStorage := make(hashMap, 256)
	lineRegex := regexp.MustCompile(`^(\w+)([-=])(\d+)?`)
	for _, row := range strings.Split(input, ",") {
		parsedRow := lineRegex.FindStringSubmatch(row)
		switch parsedRow[2] {
		case "-":
			lensStorage.delete(parsedRow[1])
		case "=":
			num, _ := strconv.Atoi(parsedRow[3])
			lensStorage.add(parsedRow[1], num)
		}
	}
	log.Printf("Focal length sum: %d", lensStorage.getFocusingPower())
}

func calculateHashSum(s string) int {
	var result int
	for _, c := range s {
		result = ((result + int(c)) * 17) % 256
	}
	return result
}
