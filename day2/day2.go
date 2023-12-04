package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var availableCubes = map[string]int{"red": 12, "green": 13, "blue": 14}

func getGameId(line string) int {
	text := strings.Replace(line, "Game ", "", 1)
	gameId, err := strconv.Atoi(text)
	if err != nil {
		panic(err)
	}
	return gameId
}
func getGameColours(game string) (r, g, b int) {
	r = 0
	g = 0
	b = 0
	cubes := strings.Split(game, ", ")
	for _, cube := range cubes {
		trimmedCube := strings.TrimSpace(cube)
		if strings.HasSuffix(trimmedCube, "red") {
			r = parse(cube, "red")
		}
		if strings.HasSuffix(trimmedCube, "green") {
			g = parse(cube, "green")
		}
		if strings.HasSuffix(trimmedCube, "blue") {
			b = parse(cube, "blue")
		}

	}
	return r, g, b
}

func parse(cube string, colour string) int {
	strValue := strings.TrimSpace(strings.Replace(cube, colour, "", 1))
	value, err := strconv.Atoi(strValue)
	if err != nil {
		panic(err)
	}
	return value
}

func isGamePossible(gamesString string) (isPossible bool, setId, setPower int) {
	// get gamesString Id
	isPossible = true
	split := strings.Split(gamesString, ":")
	setId = getGameId(split[0])
	// get games string
	games := strings.Split(split[1], ";")
	// get number of cubes
	var minR, minG, minB int
	for _, game := range games {
		r, g, b := getGameColours(game)
		// determine is possible
		if r > availableCubes["red"] || g > availableCubes["green"] || b > availableCubes["blue"] {
			isPossible = false
		}
		minR = max(minR, r)
		minG = max(minG, g)
		minB = max(minB, b)

	}
	setPower = minR * minG * minB
	return isPossible, setId, setPower
}

func Day2() {
	file, err := os.Open("day2/data.txt")
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}
	var possibleIds []int
	var setPowers []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		possible, gameId, setPower := isGamePossible(scanner.Text())
		setPowers = append(setPowers, setPower)
		if possible {
			possibleIds = append(possibleIds, gameId)
		}
	}

	availableSetsSum := 0
	for _, value := range possibleIds {
		availableSetsSum += value
	}
	setPowersSum := 0
	for _, value := range setPowers {
		setPowersSum += value
	}
	fmt.Printf("availableSetsSum: %d \n", availableSetsSum)
	fmt.Printf("setPowersSum: %d \n", setPowersSum)
}
