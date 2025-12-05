package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type MaxMinRange struct {
	Min int
	Max int
}

func (r *MaxMinRange) Contains(num int) bool {
	return num >= r.Min && num <= r.Max
}

func (r *MaxMinRange) Shares(other *MaxMinRange) []int {
	shared := []int{}
	for i := r.Min; i <= r.Max; i++ {
		if other.Contains(i) {
			shared = append(shared, i)
		}
	}
	return shared
}

func NewRangeFromStr(min string, max string) (*MaxMinRange, error) {
	minInt, err := strconv.Atoi(min)
	if err != nil {
		return nil, err
	}
	maxInt, err := strconv.Atoi(max)
	if err != nil {
		return nil, err
	}
	return &MaxMinRange{
		Max: maxInt,
		Min: minInt,
	}, nil
}

func getRangesAndIngredients(file string) ([]*MaxMinRange, []string, error) {
	content, err := os.ReadFile(file)
	if err != nil {
		return nil, nil, err
	}
	lines := strings.Split(string(content), "\n")
	emptyLineIdx := slices.Index(lines, "")
	if emptyLineIdx < 0 {
		return nil, nil, fmt.Errorf("no empty lines, Index returned %d", emptyLineIdx)
	}
	ranges := lines[:emptyLineIdx]
	foodIds := lines[emptyLineIdx+1:]
	rangesMaxMin := []*MaxMinRange{}
	for i, r := range ranges {
		parts := strings.Split(r, "-")
		if len(parts) < 2 {
			return nil, nil, fmt.Errorf("line %d is not a range: %s", i, r)
		}
		rng, err := NewRangeFromStr(parts[0], parts[1])
		if err != nil {
			return nil, nil, err
		}
		rangesMaxMin = append(rangesMaxMin, rng)
	}

	return rangesMaxMin, foodIds, nil
}

func getFreshFoodIds(foodIds []string, ranges []*MaxMinRange) (int, error) {
	totalCount := 0
	for _, idx := range foodIds {
		idxInt, err := strconv.Atoi(idx)
		if err != nil {
			return 0, nil
		}
		for _, r := range ranges {
			if r.Contains(idxInt) {
				totalCount++
				break
			}
		}
	}
	return totalCount, nil
}

func CountFreshFoods(file string) (int, error) {
	ranges, foodIds, err := getRangesAndIngredients(file)
	if err != nil {
		return 0, err
	}
	return getFreshFoodIds(foodIds, ranges)
}

func CountFreshFoodsComplex(file string) (int, error) {
	ranges, _, err := getRangesAndIngredients(file)
	if err != nil {
		return 0, err
	}
	shared := []int{}
	for i, r := range ranges {
		for j := i; j < len(ranges); j++ {
			sharedInternal := r.Shares(ranges[j])
			for _, el := range sharedInternal {
				if !slices.Contains(shared, el) {
					shared = append(shared, el)
				}
			}
		}
	}
	return len(shared), nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("You should provide one positional argument (choices: 'simple' and 'complex')")
	}
	complex := false
	if os.Args[1] == "complex" {
		complex = true
	}
	if !complex {
		result, err := CountFreshFoods("input.txt")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result)
	} else {
		result, err := CountFreshFoodsComplex("input.txt")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result)
	}
}
