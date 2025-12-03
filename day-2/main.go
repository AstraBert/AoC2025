package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Id struct {
	Value string
}

func (i *Id) Validate() bool {
	if strings.HasPrefix(i.Value, "0") {
		return false
	}
	if len(i.Value)%2 != 0 {
		return true
	}
	middle := len(i.Value) / 2
	firstHalf := i.Value[:middle]
	seconHalf := i.Value[middle:]
	return firstHalf != seconHalf
}

func (i *Id) ValidateComplex() bool {
	if strings.HasPrefix(i.Value, "0") {
		return false
	}
	n := len(i.Value)
	for splitIdx := 1; splitIdx <= n/2; splitIdx++ {
		subStr := i.Value[:splitIdx]
		if n%splitIdx == 0 {
			repeated := strings.Repeat(subStr, n/splitIdx)
			if repeated == i.Value {
				return false
			}
		}
	}
	return true
}

func obtainInvalidIds(line string, complex bool) (int, error) {
	idRanges := strings.Split(line, ",")
	invalidIds := 0
	for _, idRange := range idRanges {
		parts := strings.Split(idRange, "-")
		idRangeStart, err := strconv.Atoi(parts[0])
		if err != nil {
			return 0, err
		}
		idRangeStop, err := strconv.Atoi(parts[1])
		if err != nil {
			return 0, err
		}
		for i := idRangeStart; i <= idRangeStop; i++ {
			idValue := &Id{Value: strconv.Itoa(i)}
			if !complex {
				if valid := idValue.Validate(); !valid {
					invalidIds += i
				}
			} else {
				if valid := idValue.ValidateComplex(); !valid {
					invalidIds += i
				}
			}
		}
	}
	return invalidIds, nil
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
	ids, err := obtainInvalidIds(strings.Trim(string(content), "\n"), complex)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ids)
}
