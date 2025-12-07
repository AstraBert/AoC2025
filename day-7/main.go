package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

type Coordinates struct {
	X int
	Y int
}

func NewCoordinates(x int, y int) *Coordinates {
	return &Coordinates{
		X: x,
		Y: y,
	}
}

func getMatrix(file string) ([][]rune, error) {
	content, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(content), "\n")
	matrix := make([][]rune, len(lines))
	for i, line := range lines {
		matrix[i] = make([]rune, len(line))
		for j, c := range line {
			matrix[i][j] = c
		}
	}
	return matrix, nil
}

func CountSplits(file string) (int, error) {
	matrix, err := getMatrix(file)
	if err != nil {
		return 0, err
	}
	beams := []int{slices.Index(matrix[0], 'S')}
	if beams[0] == -1 {
		return 0, fmt.Errorf("could not find 'S', index returned: %d", beams[0])
	}
	splits := 0
	for _, line := range matrix[1:] {
		nextBeams := map[int]bool{}
		for _, b := range beams {
			if line[b] == '.' {
				nextBeams[b] = true
			} else {
				nextBeams[b+1] = true
				nextBeams[b-1] = true
				splits++
			}
		}
		newBeams := []int{}
		for b := range nextBeams {
			newBeams = append(newBeams, b)
		}
		beams = newBeams
	}
	return splits, nil
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("You should provide one positional argument (choices: 'simple' and 'complex')")
	}
	complex := false
	if os.Args[1] == "complex" {
		complex = true
	}
	if !complex {
		result, err := CountSplits("input.txt")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result)
	} else {
		panic(errors.New("solution not implemented yet"))
	}
}
