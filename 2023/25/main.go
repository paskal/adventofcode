package main

import (
	_ "embed"
	"log"
	"math/rand"
	"strings"
)

//go:embed input.txt
var input string

type VectorWithNulls struct {
	keys   []uint32
	values []int64
}

func (vec VectorWithNulls) Add(k uint32, v int64) {
	if len(vec.keys) == 0 {
		vec.keys = []uint32{k}
		vec.values = []int64{v}
		return
	}
	var lastGoodIndex int
	for _, value := range vec.keys {
		if value > k {
			vec.keys = insert(vec.keys, lastGoodIndex, k)
			vec.values = insert(vec.values, lastGoodIndex, v)
			break
		}
		lastGoodIndex++
	}
}

// multiply two non-zero length VectorWithNulls and return the result
func (vec VectorWithNulls) Multiply(another VectorWithNulls) int {
	var result int64
	var secondIndex int
	for firstIndex := 0; firstIndex < len(vec.keys); firstIndex++ {
		parentKey := vec.keys[firstIndex]
		for {
			if parentKey < another.keys[secondIndex] {
				result += another.values[secondIndex]
			}
			if parentKey == another.keys[secondIndex] {
				result += vec.values[firstIndex] * another.values[secondIndex]
			}
			if parentKey > another.keys[secondIndex] {
				break
			}
			secondIndex++
		}
	}
	for ; secondIndex < len(another.keys); secondIndex++ {
		result += another.values[secondIndex]
	}
	return int(result)
}

func insert[T any](a []T, index int, value T) []T {
	if len(a) == index {
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...)
	a[index] = value
	return a
}

func main() {
	for i, s := range strings.Split(input, "\n") {
		log.Printf("line %d: %q", i+1, s)
	}
	//in, err := os.ReadFile("./input.txt")
	//if err != nil {
	//	log.Panicf("Can't open file, %v", err)
	//}
	//for _, row := range strings.Split(in, "\n") {
	//	entries := (strings.Split(row, ","))
	//}

	randInt := rand.Int()
}
