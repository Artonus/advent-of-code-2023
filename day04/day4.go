package day04

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type sketchcard struct {
	winningNumbers []int
	numbers        []int
	count          int
}

func getNumbers(line string) []int {
	var numbers []int
	for _, s := range strings.Split(line, " ") {
		trimmed := strings.TrimSpace(s)
		if len(trimmed) == 0 {
			continue
		}
		number, err := strconv.Atoi(trimmed)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, number)
	}
	return numbers
}

func getCardNumbers(line string) []int {
	numbers := strings.Split(line, "|")
	return getNumbers(numbers[1])
}

func getWinningNumbers(line string) []int {
	numbers := strings.Split(line, "|")
	return getNumbers(numbers[0])
}

func getPoints(numbers, winningNumbers []int) (matches, points int) {
	matches = 0

	for _, number := range numbers {
		if slices.Contains(winningNumbers, number) {
			matches++
		}
	}
	if matches == 1 {
		return 1, 1
	}
	return matches, int(math.Pow(2, float64(matches-1)))
}

func getCardId(line string) int {
	split := strings.Split(line, ":")
	numStr := strings.TrimSpace(strings.Replace(split[0], "Card", "", 1))
	num, err := strconv.Atoi(numStr)
	if err != nil {
		panic(err)
	}
	return num
}
func readSketchcards(fileName string) map[int]*sketchcard {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	sketchcards := make(map[int]*sketchcard)
	for scanner.Scan() {
		line := scanner.Text()
		cardId := getCardId(line)
		split := strings.Split(line, ":")
		winningPoints := getWinningNumbers(split[1])
		cardPoints := getCardNumbers(split[1])
		sketchcards[cardId] = &sketchcard{
			numbers:        cardPoints,
			winningNumbers: winningPoints,
			count:          1,
		}
	}
	return sketchcards
}
func Day4() {
	fileName := "day04/data.txt"

	cardPoints := make(map[int]int)
	sketchCards := readSketchcards(fileName)
	for i := 1; i <= len(sketchCards); i++ {
		card := sketchCards[i]
		for j := 0; j < card.count; j++ {
			matches, points := getPoints(card.numbers, card.winningNumbers)
			cardPoints[i] = points

			for k := 1; k <= matches; k++ {
				sc := sketchCards[i+k]
				sc.count++
			}
		}
	}

	pointsSum := 0
	for _, point := range cardPoints {
		pointsSum += point
	}
	sketchCardsSum := 0
	for _, point := range sketchCards {
		sketchCardsSum += point.count
	}
	fmt.Printf("Sum points: %d \n", pointsSum)
	fmt.Printf("Sketchcard points: %d", sketchCardsSum)
}
