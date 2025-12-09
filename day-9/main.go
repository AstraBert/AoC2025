package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	pin "github.com/serxoz/pinpol"
)

type SimplePoint struct {
	X int
	Y int
}

func (p *SimplePoint) ToPinPoint() pin.Point {
	return pin.Point{
		X: float64(p.X),
		Y: float64(p.Y),
	}
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

func isPointInList(point *SimplePoint, polygon []pin.Point) bool {
	return pin.IsInside(polygon, len(polygon), point.ToPinPoint())
}

func toPinPointsList(points []*SimplePoint) []pin.Point {
	pinpoints := make([]pin.Point, len(points))
	for i := range points {
		pinpoints[i] = points[i].ToPinPoint()
	}
	return pinpoints
}

func getMaxAreaComplex(points []*SimplePoint) int {
	maxArea := 0
	pinpoints := toPinPointsList(points)
	for i := range points {
		p1 := points[i]
		for j := i + 1; j < len(points); j++ {
			p2 := points[j]
			area, p3, p4 := p1.AreaAndCorners(p2)
			if isPointInList(p3, pinpoints) && isPointInList(p4, pinpoints) {
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}
	return maxArea
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
	points, err := getPoints(file)
	if err != nil {
		return 0, err
	}
	maxArea := getMaxAreaComplex(points)
	return maxArea, nil
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
