package day6

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// distance = tHold * (tMax-tHold)
type race struct {
	time     int
	distance int
}

func getNumWaysToWinRace(race race) int {
	numWaysToWin := 0
	for tHold := 0; tHold <= race.time; tHold++ {
		dist := tHold * (race.time - tHold)
		if dist > race.distance {
			numWaysToWin++
		}
	}
	return numWaysToWin
}
func getValues(line, prefix string) []int {
	valuesLine := strings.Replace(line, prefix, "", 1)
	split := strings.Split(valuesLine, "  ")
	var values []int
	for _, i2 := range split {
		trimmed := strings.TrimSpace(i2)
		if len(trimmed) == 0 {
			continue
		}
		val, err := strconv.Atoi(trimmed)
		if err != nil {
			panic(err)
		}
		values = append(values, val)
	}
	return values
}
func getSingleRace(fileName string) race {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		closeFileErr := file.Close()
		if closeFileErr != nil {
			panic(err)
		}
	}(file)
	scanner := bufio.NewScanner(file)
	var time int
	var distance int
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Time:") {
			trimmed := strings.Replace(line, "Time:", "", 1)
			strValue := strings.ReplaceAll(trimmed, " ", "")
			value, convErr := strconv.Atoi(strValue)
			if err != nil {
				panic(convErr)
			}
			time = value
		}
		if strings.HasPrefix(line, "Distance:") {
			trimmed := strings.Replace(line, "Distance:", "", 1)
			strValue := strings.ReplaceAll(trimmed, " ", "")
			value, convErr := strconv.Atoi(strValue)
			if err != nil {
				panic(convErr)
			}
			distance = value
		}
	}
	return race{
		time:     time,
		distance: distance,
	}
}
func getRaces(fileName string) []race {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		closeFileErr := file.Close()
		if closeFileErr != nil {
			panic(err)
		}
	}(file)
	scanner := bufio.NewScanner(file)
	var times []int
	var distances []int
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Time:") {
			times = getValues(line, "Time:")
		}
		if strings.HasPrefix(line, "Distance:") {
			distances = getValues(line, "Distance:")
		}
	}

	if len(times) != len(distances) {
		panic("Wrong number of times and distances")
	}
	races := make([]race, len(times))
	for i := 0; i < len(times); i++ {
		races[i] = race{
			time:     times[i],
			distance: distances[i],
		}
	}
	return races
}
func Day6() {
	fileName := "day6/data.txt"
	races := getRaces(fileName)

	waysToWin := make([]int, len(races))

	for i := 0; i < len(races); i++ {
		waysToWin[i] = getNumWaysToWinRace(races[i])
	}

	product := 1
	for _, value := range waysToWin {
		product *= value
	}
	fmt.Printf("You can win a race in %d many ways\n", product)

	raceScenario2 := getSingleRace(fileName)
	waysToWinSingleRace := getNumWaysToWinRace(raceScenario2)
	fmt.Printf("You can win a single race in %d many ways", waysToWinSingleRace)
}
