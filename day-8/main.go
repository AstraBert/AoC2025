package main

import (
	"fmt"
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
	return uint64((p1.X - p2.X) ^ 2 + (p1.Y - p2.Y) ^ 2 + (p1.Z - p2.Z) ^ 2)
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
	points := make([]*Point, len(lines))
	for i, line := range lines {
		points[i] = NewPointFromStr(strings.Trim(line, "\n"))
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
		if l < 1000 {
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
	circuits := [][]string{{pairs[0].From, pairs[0].To}}
	for _, pair := range pairs[1:] {
		appended := false
		for i, circuit := range circuits {
			lastEl := circuit[len(circuit)-1]
			if pair.From == lastEl {
				circuit = append(circuit, pair.To)
				circuits[i] = circuit
				appended = true
				break
			}
		}
		if !appended {
			circuits = append(circuits, []string{pair.From, pair.To})
		}
	}
	slices.SortFunc(circuits, func(a []string, b []string) int {
		return len(b) - len(a)
	})
	total := 1
	for _, circuit := range circuits[:3] {
		fmt.Println(circuit)
		total *= len(circuit)
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
	s, _ := GetCircuits("input.txt", 1000)
	fmt.Println(s)
}
