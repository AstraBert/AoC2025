package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func getLines(file string) ([]string, error) {
	content, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(content), "\n"), nil
}

func getServerDistribution(lines []string) map[string][]string {
	srvs := make(map[string][]string)
	for _, line := range lines {
		parts := strings.Split(line, ":")
		srvs[parts[0]] = []string{}
		for child := range strings.FieldsSeq(parts[1]) {
			srvs[parts[0]] = append(srvs[parts[0]], strings.TrimSpace(child))
		}
	}
	return srvs
}

func getToOut(servers map[string][]string, server string) int {
	paths := 0
	if servers[server][0] == "out" {
		return 1
	}
	for _, s := range servers[server] {
		numPaths := getToOut(servers, s)
		paths += numPaths
	}
	return paths
}

func getToOutComplex(servers map[string][]string, server string) [][]string {
	allPaths := [][]string{}
	if servers[server][0] == "out" {
		return [][]string{{server}}
	}
	for _, s := range servers[server] {
		childPaths := getToOutComplex(servers, s)
		for _, childPath := range childPaths {
			newPath := append([]string{server}, childPath...)
			allPaths = append(allPaths, newPath)
		}
	}
	return allPaths
}

func getAllPaths(servers map[string][]string) int {
	paths := 0
	for k := range servers {
		if k == "you" {
			paths += getToOut(servers, k)
		}
	}
	return paths
}

func getAllPathsComplex(servers map[string][]string) int {
	paths := [][]string{}
	for k := range servers {
		if k == "svr" {
			paths = getToOutComplex(servers, k)
		}
	}
	count := 0
	for _, p := range paths {
		if slices.Contains(p, "fft") && slices.Contains(p, "dac") {
			count += 1
		}
	}
	return count
}

func FindAllWaysOut(file string) (int, error) {
	lines, err := getLines(file)
	if err != nil {
		return 0, err
	}
	servers := getServerDistribution(lines)
	return getAllPaths(servers), nil
}

func FindAllWaysOutComplex(file string) (int, error) {
	lines, err := getLines(file)
	if err != nil {
		return 0, err
	}
	servers := getServerDistribution(lines)
	return getAllPathsComplex(servers), nil
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
		s, err := FindAllWaysOut("input.txt")
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Println(s)
	} else {
		s, err := FindAllWaysOutComplex("input.txt")
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Println(s)
	}
}
