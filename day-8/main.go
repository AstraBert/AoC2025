package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
	Z int
}

func NewPointFromStr(s string) *Point {
	parts := strings.Split(s, ",")
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])
	z, _ := strconv.Atoi(parts[2])
	return &Point{
		X: x,
		Y: y,
		Z: z,
	}
}

func (p1 *Point) Distance(p2 *Point) uint64 {
	dx := p1.X - p2.X
	dy := p1.Y - p2.Y
	dz := p1.Z - p2.Z
	return uint64(dx*dx + dy*dy + dz*dz)
}

func (p *Point) ToString() string {
	return fmt.Sprintf("%d,%d,%d", p.X, p.Y, p.Z)
}

type PointPair struct {
	From     string
	To       string
	Distance uint64
}

func (pp1 *PointPair) ShouldSwap(pp2 *PointPair) bool {
	return pp1.Distance > pp2.Distance
}

func NewPointPair(from, to string, distance uint64) *PointPair {
	return &PointPair{
		From:     from,
		To:       to,
		Distance: distance,
	}
}

func linesToPoints(file string) ([]*Point, error) {
	content, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(content), "\n")
	points := make([]*Point, 0, len(lines))
	for _, line := range lines {
		points = append(points, NewPointFromStr(strings.Trim(line, "\n")))
	}
	return points, nil
}

func calculateDistances(points []*Point) []*PointPair {
	distances := []*PointPair{}
	for i, p1 := range points {
		for j := i + 1; j < len(points); j++ {
			p2 := points[j]
			distance := p1.Distance(p2)
			distances = append(distances, NewPointPair(p1.ToString(), p2.ToString(), distance))
		}
	}
	return distances
}

func getNMinDistances(n int, pairs []*PointPair) []*PointPair {
	if n == 0 {
		return nil
	}
	if n >= len(pairs) {
		return pairs
	}
	minDists := []*PointPair{}
	for _, p := range pairs {
		l := len(minDists)
		if l < n {
			minDists = append(minDists, p)
		} else {
			maxDistIdx := 0
			for i := range minDists {
				if minDists[i].Distance > minDists[maxDistIdx].Distance {
					maxDistIdx = i
				}
			}
			if minDists[maxDistIdx].ShouldSwap(p) {
				minDists[maxDistIdx] = p
			}
		}
	}
	return minDists
}

func chainPairs(pairs []*PointPair) int {
	pointToCircuit := make(map[string]int)
	circuits := make(map[int][]string)

	nextCircuitID := 0

	for _, pair := range pairs {
		fromCircuit, fromExists := pointToCircuit[pair.From]
		toCircuit, toExists := pointToCircuit[pair.To]

		if !fromExists && !toExists {
			circuits[nextCircuitID] = []string{pair.From, pair.To}
			pointToCircuit[pair.From] = nextCircuitID
			pointToCircuit[pair.To] = nextCircuitID
			nextCircuitID++

		} else if fromExists && !toExists {
			circuits[fromCircuit] = append(circuits[fromCircuit], pair.To)
			pointToCircuit[pair.To] = fromCircuit

		} else if !fromExists && toExists {
			circuits[toCircuit] = append(circuits[toCircuit], pair.From)
			pointToCircuit[pair.From] = toCircuit

		} else {
			if fromCircuit == toCircuit {
				continue
			} else {
				for _, point := range circuits[toCircuit] {
					circuits[fromCircuit] = append(circuits[fromCircuit], point)
					pointToCircuit[point] = fromCircuit
				}
				delete(circuits, toCircuit)
			}
		}
	}

	sizes := []int{}
	for _, circuit := range circuits {
		sizes = append(sizes, len(circuit))
	}

	slices.SortFunc(sizes, func(a int, b int) int {
		return b - a
	})

	// Multiply top 3
	total := 1
	for i := 0; i < 3 && i < len(sizes); i++ {
		total *= sizes[i]
	}

	return total
}

func GetCircuits(file string, numPairs int) (int, error) {
	points, err := linesToPoints(file)
	if err != nil {
		return 0, err
	}
	distances := calculateDistances(points)
	minDists := getNMinDistances(numPairs, distances)
	circuits := chainPairs(minDists)
	return circuits, nil
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
		s, err := GetCircuits("input.txt", 1000)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Println(s)
	} else {
		panic(errors.New("solution not implemented yet"))
	}
}
