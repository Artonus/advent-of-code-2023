package day08

import (
	"advent-of-code-2023/pkg/math"
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
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
func getStartingSteps(directions map[string]direction) []string {
	var steps []string
	for key, _ := range directions {
		if strings.HasSuffix(key, "A") {
			steps = append(steps, key)
		}
	}
	return steps
}
func areAllDestinationSteps(steps []string) bool {
	for _, step := range steps {
		if !strings.HasSuffix(step, "Z") {
			return false
		}
	}
	return true
}

func getSteps2(maps map[string]direction, startingPoints []string, directions string) int {
	var (
		results []int
		wg      sync.WaitGroup
		lcm     int
	)
	results = make([]int, len(startingPoints))
	for i := 0; i < len(startingPoints); i++ {
		wg.Add(1)
		go func(start string, res *int) {
			defer wg.Done()

			*res = 0
			dirLen := len(directions)
			current := start
			for !strings.HasSuffix(current, "Z") {
				next := directions[*res%dirLen]
				if string(next) == "L" {
					current = maps[current].l

				}
				if string(next) == "R" {
					current = maps[current].r
				}
				*res++
			}
		}(startingPoints[i], &results[i])

		// use gorutines to calculate number of steps for each of the starting points,
		// then use the lowest common multiplication of them all
	}
	wg.Wait()
	lcm = math.LCM(results[0], results[1:])

	return lcm
}
func Day8() {
	fileName := "day08/data.txt"

	directions, maps := getDirections(fileName)

	stepsList := getStartingSteps(maps)
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
	fmt.Printf("It took %d steps to reach ZZZ\n", steps)

	steps2 := getSteps2(maps, stepsList, directions)
	fmt.Printf("It took %d steps to reach all Z's", steps2)

}
