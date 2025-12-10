package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Machine struct {
	Lights   []int
	Switches [][]int
	Jolts    []int
}

func getLines(file string) ([]string, error) {
	content, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(content), "\n"), nil
}

func convertToLights(lights string) []int {
	lights = lights[1 : len(lights)-1]
	ls := make([]int, len(lights))
	for i, l := range lights {
		if l == '#' {
			ls[i] = 1
		} else {
			ls[i] = 0
		}
	}
	return ls
}

func convertToSwitches(switches string, machineLength int) [][]int {
	ls := strings.Split(switches, ") ")
	finalLs := [][]int{}
	for _, l := range ls {
		button := strings.TrimSpace(strings.ReplaceAll(strings.ReplaceAll(l, ")", ""), "(", ""))
		buttonLs := []string{}
		if strings.ContainsRune(button, ',') {
			buttonLs = strings.Split(button, ",")
		} else {
			buttonLs = append(buttonLs, button)
		}
		intLs := []int{}
		for j := range machineLength {
			sJ := strconv.Itoa(j)
			if slices.Contains(buttonLs, sJ) {
				intLs = append(intLs, 1)
			} else {
				intLs = append(intLs, 0)
			}
		}
		finalLs = append(finalLs, intLs)
	}
	return finalLs
}

func convertToJolts(jolts string) []int {
	jolts = jolts[1 : len(jolts)-1]
	nums := strings.Split(jolts, ",")
	ints := make([]int, len(nums))
	for i, n := range nums {
		j, _ := strconv.Atoi(n)
		ints[i] = j
	}
	return ints
}

func getGroupsRegex(lines []string) []*Machine {
	r, _ := regexp.Compile(`(\[[\.\#]*\])([\(\d\,)\s]*)(\{[\d,]*\})`)
	machines := []*Machine{}
	for _, line := range lines {
		matches := r.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			m := &Machine{}
			m.Lights = convertToLights(match[1])
			m.Switches = convertToSwitches(match[2], len(m.Lights))
			m.Jolts = convertToJolts(match[3])
			machines = append(machines, m)
		}
	}
	return machines
}

func lightMachine(m *Machine) int {
	n := len(m.Lights)
	numButtons := len(m.Switches)
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, numButtons+1)
		for j := range numButtons {
			matrix[i][j] = m.Switches[j][i]
		}
		matrix[i][numButtons] = m.Lights[i]
	}

	pivotCol := 0
	for row := 0; row < n && pivotCol < numButtons; row++ {
		pivotRow := -1
		for r := row; r < n; r++ {
			if matrix[r][pivotCol] == 1 {
				pivotRow = r
				break
			}
		}

		if pivotRow == -1 {
			pivotCol++
			row--
			continue
		}

		matrix[row], matrix[pivotRow] = matrix[pivotRow], matrix[row]

		for r := range n {
			if r != row && matrix[r][pivotCol] == 1 {
				for c := 0; c <= numButtons; c++ {
					matrix[r][c] ^= matrix[row][c]
				}
			}
		}
		pivotCol++
	}

	freeVars := []int{}

	for col := range numButtons {
		isPivot := false
		for row := range n {
			if matrix[row][col] == 1 {
				isLeading := true
				for c := 0; c < col; c++ {
					if matrix[row][c] == 1 {
						isLeading = false
						break
					}
				}
				if isLeading {
					isPivot = true
					break
				}
			}
		}
		if !isPivot {
			freeVars = append(freeVars, col)
		}
	}

	minPresses := numButtons + 1
	numFreeVars := len(freeVars)

	for combo := 0; combo < (1 << numFreeVars); combo++ {
		testSol := make([]int, numButtons)
		for i, fv := range freeVars {
			if (combo>>i)&1 == 1 {
				testSol[fv] = 1
			}
		}

		for row := n - 1; row >= 0; row-- {
			pivotCol := -1
			for col := range numButtons {
				if matrix[row][col] == 1 {
					pivotCol = col
					break
				}
			}

			if pivotCol == -1 {
				continue
			}

			val := matrix[row][numButtons]
			for col := pivotCol + 1; col < numButtons; col++ {
				if matrix[row][col] == 1 {
					val ^= testSol[col]
				}
			}
			testSol[pivotCol] = val
		}

		presses := 0
		for _, v := range testSol {
			presses += v
		}

		if presses < minPresses {
			minPresses = presses
		}
	}

	return minPresses
}

func LightAllMachines(file string) (int, error) {
	lines, err := getLines(file)
	if err != nil {
		return 0, err
	}
	machines := getGroupsRegex(lines)
	presses := 0
	for _, m := range machines {
		presses += lightMachine(m)
	}
	return presses, nil
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
		s, err := LightAllMachines("input.txt")
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Println(s)
	} else {
		panic(errors.New("solution not yet implemented"))
	}
}
