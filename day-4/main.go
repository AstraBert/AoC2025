package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// [i][j] -> [i][j-1], [i][j+1], [i-1][j], [i-1][j-1], [i-1][j+1], [i+1][j], [i+1][j-1], [i+1][j+1]

func getMatrix(lines string) [][]rune {
	rows := strings.Split(lines, "\n")
	matrix := [][]rune{}
	for _, row := range rows {
		columns := []rune{}
		for _, c := range row {
			columns = append(columns, c)
		}
		matrix = append(matrix, columns)
	}
	return matrix
}

func scanMatrix(matrix [][]rune) int {
	totalCount := 0
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] != '@' {
				continue
			}
			previousRow := i - 1
			nextRow := i + 1
			previousColumn := j - 1
			nextColumn := j + 1
			values := []rune{}
			if nextColumn < len(matrix[i]) {
				values = append(values, matrix[i][nextColumn])
			}
			if previousColumn >= 0 {
				values = append(values, matrix[i][previousColumn])
			}
			if previousRow >= 0 {
				values = append(values, matrix[previousRow][j])
				if previousColumn >= 0 {
					values = append(values, matrix[previousRow][previousColumn])
				}
				if nextColumn < len(matrix[previousRow]) {
					values = append(values, matrix[previousRow][nextColumn])
				}
			}
			if nextRow < len(matrix) {
				values = append(values, matrix[nextRow][j])
				if previousColumn >= 0 {
					values = append(values, matrix[nextRow][previousColumn])
				}
				if nextColumn < len(matrix[nextRow]) {
					values = append(values, matrix[nextRow][nextColumn])
				}
			}
			neighborsCount := 0
			for _, val := range values {
				if val == '@' {
					neighborsCount++
				}
			}
			if neighborsCount < 4 {
				totalCount++
			}
		}
	}
	return totalCount
}

func scanMatrixRecursive(matrix [][]rune) int {
	totalCounts := 0
	for {
		totalCount := 0
		for i := range matrix {
			for j := range matrix[i] {
				if matrix[i][j] != '@' {
					continue
				}
				previousRow := i - 1
				nextRow := i + 1
				previousColumn := j - 1
				nextColumn := j + 1
				values := []rune{}
				if nextColumn < len(matrix[i]) {
					values = append(values, matrix[i][nextColumn])
				}
				if previousColumn >= 0 {
					values = append(values, matrix[i][previousColumn])
				}
				if previousRow >= 0 {
					values = append(values, matrix[previousRow][j])
					if previousColumn >= 0 {
						values = append(values, matrix[previousRow][previousColumn])
					}
					if nextColumn < len(matrix[previousRow]) {
						values = append(values, matrix[previousRow][nextColumn])
					}
				}
				if nextRow < len(matrix) {
					values = append(values, matrix[nextRow][j])
					if previousColumn >= 0 {
						values = append(values, matrix[nextRow][previousColumn])
					}
					if nextColumn < len(matrix[nextRow]) {
						values = append(values, matrix[nextRow][nextColumn])
					}
				}
				neighborsCount := 0
				for _, val := range values {
					if val == '@' {
						neighborsCount++
					}
				}
				if neighborsCount < 4 {
					matrix[i][j] = '.'
					totalCount++
				}
			}
		}
		totalCounts += totalCount
		if totalCount == 0 {
			break
		}
	}
	return totalCounts
}

func GetAccessibleRolls(file string) (int, error) {
	content, err := os.ReadFile(file)
	if err != nil {
		return 0, err
	}
	matrix := getMatrix(string(content))
	totalCount := scanMatrix(matrix)
	return totalCount, nil
}

func GetAccessibleRollsComplex(file string) (int, error) {
	content, err := os.ReadFile(file)
	if err != nil {
		return 0, err
	}
	matrix := getMatrix(string(content))
	totalCount := scanMatrixRecursive(matrix)
	return totalCount, nil
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
		result, err := GetAccessibleRolls("input.txt")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result)
	} else {
		result, err := GetAccessibleRollsComplex("input.txt")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result)
	}
}
