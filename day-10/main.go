package main

import (
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
			ls[i] = i
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
		button := strings.TrimSpace(strings.ReplaceAll(l, ")", ""))
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
				intLs = append(intLs, j)
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

func main() {
	lines, _ := getLines("test.txt")
	getGroupsRegex(lines)
}
