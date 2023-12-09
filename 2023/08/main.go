package main

import (
	_ "embed"
	"log"
	"regexp"
	"strings"
)

//go:embed input.txt
var input string

type node struct {
	value       string
	left, right *node
}

func main() {
	nodes, instructions := getNodesAndInstructions()
	steps := calculateStepLength(nodes, instructions)
	log.Printf("Steps from **A to **Z: %d", steps)
}

func getNodesAndInstructions() (map[string]*node, []string) {
	nodes := map[string]*node{}
	var instructions []string
	nodeRegexp := regexp.MustCompile(`\((\w{3}), (\w{3})\)`)
	for _, s := range strings.Split(input, "\n") {
		rawLine := strings.Split(s, " = ")
		if len(rawLine) < 2 {
			if len(s) > 0 {
				instructions = strings.Split(s, "")
			}
			continue
		}
		if _, ok := nodes[rawLine[0]]; !ok {
			nodes[rawLine[0]] = &node{value: rawLine[0]}
		}
		left := nodeRegexp.FindStringSubmatch(rawLine[1])[1]
		if _, ok := nodes[left]; !ok {
			nodes[left] = &node{value: left}
		}
		nodes[rawLine[0]].left = nodes[left]
		right := nodeRegexp.FindStringSubmatch(rawLine[1])[2]
		if _, ok := nodes[right]; !ok {
			nodes[right] = &node{value: right}
		}
		nodes[rawLine[0]].right = nodes[right]
	}
	return nodes, instructions
}

func calculateStepLength(nodes map[string]*node, instructions []string) int {
	var startingNodes = []*node{}
	for _, n := range nodes {
		if strings.HasSuffix(n.value, "A") {
			startingNodes = append(startingNodes, n)
		}
	}

	var steps []int
	for _, n := range startingNodes {
		var currentNodeSteps int
		currentNode := n
		for i := 0; !strings.HasSuffix(currentNode.value, "Z"); i++ {
			switch instructions[i%len(instructions)] {
			case "L":
				currentNode = currentNode.left
			case "R":
				currentNode = currentNode.right
			}
			currentNodeSteps++
		}
		steps = append(steps, currentNodeSteps)
	}
	if len(steps) == 1 {
		return steps[0]
	}
	if len(steps) == 2 {
		return lcm(steps[0], steps[1])
	}
	return lcm(steps[0], steps[1], steps[2:]...)
}

// greatest common divisor (GCS) via Euclidean algorithm
func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find The Least Common Multiple (LCM) via GCS
func lcm(a, b int, integers ...int) int {
	result := a * b / gcd(a, b)

	for i := 0; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}

	return result
}
