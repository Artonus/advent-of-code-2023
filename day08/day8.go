package day08

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var destination = "ZZZ"
var start = "AAA"

type direction struct {
	l string
	r string
}

func getDirection(line string) (node string, dir direction) {
	split := strings.Split(line, "=")

	node = strings.TrimSpace(split[0])

	dirs := strings.TrimFunc(strings.TrimSpace(split[1]), func(r rune) bool {
		if string(r) == "(" || string(r) == ")" {
			return true
		}
		return false
	})
	dirsSplit := strings.Split(dirs, ", ")
	dir = direction{
		l: strings.TrimSpace(dirsSplit[0]),
		r: strings.TrimSpace(dirsSplit[1]),
	}
	return node, dir
}
func getDirections(fileName string) (directions string, maps map[string]direction) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	maps = make(map[string]direction)
	scanner := bufio.NewScanner(file)
	readFirstLine := false
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		if readFirstLine == false {
			directions = line
			readFirstLine = true
			continue
		}
		node, dir := getDirection(line)
		maps[node] = dir
	}
	return directions, maps
}

func Day8() {
	fileName := "day08/data.txt"

	directions, maps := getDirections(fileName)
	
	current := start
	steps := 0
	dirLen := len(directions)
	for current != destination {
		next := directions[steps%dirLen]
		if string(next) == "L" {
			current = maps[current].l
		}
		if string(next) == "R" {
			current = maps[current].r
		}
		steps++
	}
	fmt.Printf("It took %d steps to reach ZZZ", steps)
}
