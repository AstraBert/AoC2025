package main

import (
	"fmt"
	"log"
	"os"
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

type FuncCache struct {
	Server   string
	FoundDac bool
	FoundFft bool
}

func NewFuncCache(server string, foundDac, foundFft bool) FuncCache {
	return FuncCache{
		Server:   server,
		FoundDac: foundDac,
		FoundFft: foundFft,
	}
}

func getToOutComplex(servers map[string][]string, server string, foundDac bool, foundFft bool) int {
	cache := make(map[FuncCache]int)
	return getToOutComplexFuncCache(servers, server, foundDac, foundFft, cache)
}

func getToOutComplexFuncCache(servers map[string][]string, server string, foundDac bool, foundFft bool, cache map[FuncCache]int) int {
	key := NewFuncCache(server, foundDac, foundFft)
	if val, ok := cache[key]; ok {
		return val
	}
	allPaths := 0
	if servers[server][0] == "out" && foundDac && foundFft {
		return 1
	} else if servers[server][0] == "out" && (!foundDac || !foundFft) {
		return 0
	}
	switch server {
	case "fft":
		foundFft = true
	case "dac":
		foundDac = false
	}
	for _, s := range servers[server] {
		allPaths += getToOutComplex(servers, s, foundDac, foundFft)
	}
	cache[key] = allPaths
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
	paths := 0
	for k := range servers {
		if k == "svr" {
			paths += getToOutComplex(servers, k, false, false)
		}
	}
	return paths
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
