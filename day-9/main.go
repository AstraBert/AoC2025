package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type SimplePoint struct {
	X int
	Y int
}

type Edge struct {
	Start *SimplePoint
	End   *SimplePoint
}

func (p1 *SimplePoint) Area(p2 *SimplePoint) int {
	dx := math.Abs(float64(p1.X-p2.X)) + 1
	dy := math.Abs(float64(p1.Y-p2.Y)) + 1
	return int(dx * dy)
}

func (p1 *SimplePoint) AreaAndCorners(p2 *SimplePoint) (int, *SimplePoint, *SimplePoint) {
	dx := math.Abs(float64(p1.X-p2.X)) + 1
	dy := math.Abs(float64(p1.Y-p2.Y)) + 1
	corner3 := &SimplePoint{X: p1.X, Y: p2.Y}
	corner4 := &SimplePoint{X: p2.X, Y: p1.Y}

	return int(dx * dy), corner3, corner4
}

func (p1 *SimplePoint) Eq(p2 *SimplePoint) bool {
	return p1.X == p2.X && p1.Y == p2.Y
}

func NewPointFromStr(s string) *SimplePoint {
	parts := strings.Split(s, ",")
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])
	return &SimplePoint{
		X: x,
		Y: y,
	}
}

func NewEdge(p1 *SimplePoint, p2 *SimplePoint) *Edge {
	return &Edge{
		Start: p1,
		End:   p2,
	}
}

func getPoints(file string) ([]*SimplePoint, error) {
	content, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(content), "\n")
	points := make([]*SimplePoint, len(lines))
	for i, line := range lines {
		points[i] = NewPointFromStr(strings.Trim(line, "\n"))
	}
	return points, nil
}

func getMaxArea(points []*SimplePoint) int {
	maxArea := 0
	for i := range points {
		p1 := points[i]
		for j := i + 1; j < len(points); j++ {
			p2 := points[j]
			area := p1.Area(p2)
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}

func getMaxAreaComplex(file string) (int, error) {
	sort := func(a int, b int) (int, int) {
		if a < b {
			return a, b
		} else {
			return b, a
		}
	}
	abs := func(a int) int {
		return int(math.Abs(float64(a)))
	}
	manhattanDistance := func(a, b *SimplePoint) int {
		return abs(a.X-b.X) + abs(a.Y-b.Y)
	}
	content, err := os.ReadFile(file)
	if err != nil {
		return 0, err
	}
	lines := strings.Split(string(content), "\n")
	redTiles := []*SimplePoint{}
	edges := []*Edge{}
	firstTile := NewPointFromStr(strings.Trim(lines[0], "\n"))
	lastTile := NewPointFromStr(strings.Trim(lines[len(lines)-1], "\n"))

	for i := range len(lines) - 1 {
		pointFrom := NewPointFromStr(lines[i])
		pointTo := NewPointFromStr(lines[i+1])
		edges = append(edges, NewEdge(pointFrom, pointTo))
		redTiles = append(redTiles, pointFrom, pointTo)
	}
	edges = append(edges, NewEdge(firstTile, lastTile))
	intersections := func(minX, minY, maxX, maxY int) bool {
		for _, inter := range edges {
			iMinX, iMaxX := sort(inter.Start.X, inter.End.X)
			iMinY, iMaxY := sort(inter.Start.Y, inter.End.Y)
			if minX < iMaxX && maxX > iMinX && minY < iMaxY && maxY > iMinY {
				return true
			}
		}
		return false
	}
	result := 0
	for fromIndex := range len(redTiles) - 1 {
		for toIndex := fromIndex; toIndex < len(redTiles); toIndex++ {
			fromTile := redTiles[fromIndex]
			toTile := redTiles[toIndex]
			minX, maxX := sort(fromTile.X, toTile.X)
			minY, maxY := sort(fromTile.Y, toTile.Y)
			mD := manhattanDistance(fromTile, toTile)
			if mD*mD > result {
				if !intersections(minX, minY, maxX, maxY) {
					area := fromTile.Area(toTile)
					if area > result {
						result = area
					}
				}
			}
		}
	}
	return result, nil
}

func FindBiggestRectangle(file string) (int, error) {
	points, err := getPoints(file)
	if err != nil {
		return 0, err
	}
	maxArea := getMaxArea(points)
	return maxArea, nil
}

func FindBiggestRectangleComplex(file string) (int, error) {
	return getMaxAreaComplex(file)
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
		s, err := FindBiggestRectangle("input.txt")
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Println(s)
	} else {
		s, err := FindBiggestRectangleComplex("input.txt")
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Println(s)
	}
}
