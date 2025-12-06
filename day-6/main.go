package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var operations map[string]func(...uint64) uint64 = map[string]func(...uint64) uint64{
	"+": func(toAdd ...uint64) uint64 {
		var sum uint64 = 0
		for _, num := range toAdd {
			sum += num
		}
		return sum
	},
	"*": func(toMultiply ...uint64) uint64 {
		var product uint64 = 1
		for _, num := range toMultiply {
			product *= num
		}
		return product
	},
}

func getNumbersFromLines(lines []string) uint64 {
	var total uint64 = 0
	nums := []string{}
	for i := len(lines[0]) - 1; i >= 0; i-- {
		num := ""
		for j, line := range lines {
			hasOperation := false
			if unicode.IsDigit(rune(line[i])) {
				num += string(line[i])
			} else if rune(line[i]) == '*' || rune(line[i]) == '+' {
				nums = append(nums, num)
				operator := string(line[i])
				operands := []uint64{}
				for _, num := range nums {
					operand, _ := strconv.Atoi(num)
					operands = append(operands, uint64(operand))
				}
				total += operations[operator](operands...)
				nums = []string{}
				hasOperation = true
			}
			if j == len(lines)-1 && num != "" && !hasOperation {
				nums = append(nums, num)
			}
		}
	}
	return total
}

func getProblemsInputs(file string) ([][]string, error) {
	content, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(content), "\n")
	lineContents := [][]string{}
	for _, line := range lines {
		lineContent := strings.Fields(strings.Trim(line, "\n"))
		lineContents = append(lineContents, lineContent)
	}
	columns := make([][]string, len(lineContents[0]))
	for i := range len(lineContents[0]) {
		for j := range len(lineContents) {
			columns[i] = append(columns[i], lineContents[j][i])
		}
	}
	return columns, nil
}

func convertNum(s string) (uint64, error) {
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return uint64(n), nil
}

func getProblemOutputs(inputs [][]string) (uint64, error) {
	var totalCount uint64 = 0
	for _, input := range inputs {
		operator := ""
		operands := []uint64{}
		for _, s := range input {
			if s == "+" || s == "*" {
				operator = s
			} else {
				operand, err := convertNum(s)
				if err != nil {
					return 0, err
				}
				operands = append(operands, operand)
			}
		}
		if operator == "" {
			return 0, errors.New("expected an operator between '+' and '*', got none")
		}
		totalCount += operations[operator](operands...)
	}
	return totalCount, nil
}

func SolveAllProblems(file string, complex bool) (uint64, error) {
	if !complex {
		inputs, err := getProblemsInputs(file)
		if err != nil {
			return 0, err
		}
		return getProblemOutputs(inputs)
	} else {
		content, err := os.ReadFile(file)
		if err != nil {
			return 0, err
		}
		lines := strings.Split(string(content), "\n")
		return getNumbersFromLines(lines), nil
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("You should provide one positional argument (choices: 'simple' and 'complex')")
	}
	complex := false
	if os.Args[1] == "complex" {
		complex = true
	}
	result, err := SolveAllProblems("input.txt", complex)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
