package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func scanLine(line string) (int, error) {
	max := 0
	for startIdx := range len(line) {
		for secondIdx := startIdx + 1; secondIdx < len(line); secondIdx++ {
			number, err := strconv.Atoi(string(line[startIdx]) + string(line[secondIdx]))
			if err != nil {
				return 0, err
			}
			if number > max {
				max = number
			}
		}
	}
	return max, nil
}

func scanLineComplex(line string) (int, error) {
	toKeep := 12
	toRemove := len(line) - toKeep

	stack := []byte{}

	for i := range len(line) {
		// While we can still remove digits and current digit is larger than stack top
		for len(stack) > 0 && toRemove > 0 && stack[len(stack)-1] < line[i] {
			stack = stack[:len(stack)-1]
			toRemove--
		}
		stack = append(stack, line[i])
	}

	stack = stack[:toKeep]

	maxNum, err := strconv.Atoi(string(stack))
	return maxNum, err
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("You should provide one positional argument (choices: 'simple' and 'complex')")
	}
	complex := false
	if os.Args[1] == "complex" {
		complex = true
	}
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(content), "\n")
	jolts := 0
	for _, line := range lines {
		maxJolt := 0
		if !complex {
			maxJolt, err = scanLine(strings.Trim(line, "\n"))
			if err != nil {
				log.Fatal(err)
			}
		} else {
			maxJolt, err = scanLineComplex(strings.Trim(line, "\n"))
			if err != nil {
				log.Fatal(err)
			}
		}
		jolts += maxJolt
	}
	fmt.Println(jolts)
}
