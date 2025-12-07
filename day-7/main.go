package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

type Coordinates struct {
	X     int
	Y     int
	Score int
}

func (c *Coordinates) Sum(c1 *Coordinates) *Coordinates {
	return NewCoordinates(c.X, c.Y, c.Score+c1.Score)
}

func NewCoordinates(x int, y int, score int) *Coordinates {
	return &Coordinates{
		X:     x,
		Y:     y,
		Score: score,
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

func CountSplitsComplex(file string) (int, error) {
	matrix, err := getMatrix(file)
	if err != nil {
		return 0, err
	}
	paths := []*Coordinates{NewCoordinates(slices.Index(matrix[0], 'S'), 0, 1)}
	for y, line := range matrix[1:] {
		newPaths := []*Coordinates{}
		for _, path := range paths {
			if line[path.X] == '.' {
				path.Y = y
				newPaths = append(newPaths, path)
			} else {
				if !slices.ContainsFunc(newPaths, func(c *Coordinates) bool {
					return c.X == path.X-1 && c.Y == path.Y
				}) {
					newPaths = append(newPaths, NewCoordinates(path.X-1, path.Y, path.Score))
				} else {
					idx := slices.IndexFunc(newPaths, func(c *Coordinates) bool {
						return c.X == path.X-1 && c.Y == path.Y
					})
					newPaths[idx] = newPaths[idx].Sum(path)
				}
				if !slices.ContainsFunc(newPaths, func(c *Coordinates) bool {
					return c.X == path.X+1 && c.Y == path.Y
				}) {
					newPaths = append(newPaths, NewCoordinates(path.X+1, path.Y, path.Score))
				} else {
					idx := slices.IndexFunc(newPaths, func(c *Coordinates) bool {
						return c.X == path.X+1 && c.Y == path.Y
					})
					newPaths[idx] = newPaths[idx].Sum(path)
				}
			}
		}
		paths = newPaths
	}
	total := 0
	for _, p := range paths {
		total += p.Score
	}
	return total, nil
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
		result, err := CountSplitsComplex("input.txt")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result)
	}
}
