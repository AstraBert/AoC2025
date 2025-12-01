package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Safe struct {
	State    int
	Password int
}

func (s *Safe) Validate() {
	s.State = ((s.State % 100) + 100) % 100
	if s.State == 0 {
		s.Password++
	}
}

func rotationToInt(rotation string) (int, error) {
	if strings.HasPrefix(rotation, "L") {
		num, err := strconv.Atoi(strings.ReplaceAll(rotation, "L", ""))
		if err != nil {
			return 0, err
		}
		return -num, nil
	} else {
		num, err := strconv.Atoi(strings.ReplaceAll(rotation, "R", ""))
		if err != nil {
			return 0, err
		}
		return num, nil
	}
}

func (s *Safe) Rotate(rotation string) error {
	num, err := rotationToInt(rotation)
	if err != nil {
		return err
	}
	s.State += num
	s.Validate()
	return nil
}

func (s *Safe) RotateComplex(rotation string) error {
	num, err := rotationToInt(rotation)
	if err != nil {
		return err
	}
	isNeg := num < 0
	steps := int(math.Abs(float64(num)))
	for range steps {
		if isNeg {
			s.State -= 1
		} else {
			s.State += 1
		}
		s.Validate()
	}
	return nil
}

func SimpleMethod() {
	s := Safe{State: 50}
	bts, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.SplitSeq(string(bts), "\n")
	for line := range lines {
		line = strings.Trim(line, "\n")
		if line != "" {
			err = s.Rotate(line)
			if err != nil {
				panic(err)
			}
		}
	}
	fmt.Printf("The password is: %d\n", s.Password)
}

func ComplexMethod() {
	s := Safe{State: 50}
	bts, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.SplitSeq(string(bts), "\n")
	for line := range lines {
		line = strings.Trim(line, "\n")
		if line != "" {
			err = s.RotateComplex(line)
			if err != nil {
				panic(err)
			}
		}
	}
	fmt.Printf("The password is: %d\n", s.Password)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("You should provide one positional argument (choices: 'simple' and 'complex')")
	}
	if os.Args[1] == "simple" {
		SimpleMethod()
	} else {
		ComplexMethod()
	}
}
