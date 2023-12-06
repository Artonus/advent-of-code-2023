package day3

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type position struct {
	x    int
	y    int
	xMax int
	yMax int
}

func (p position) Offset(x, y int) position {
	return position{
		x:    p.x + x,
		y:    p.y + y,
		xMax: p.xMax,
		yMax: p.yMax,
	}
}
func (p position) IsValid() bool {
	return p.x >= 0 && p.x < p.xMax && p.y >= 0 && p.y < p.yMax
}

func isSymbol(char rune) bool {
	return char != '.' && unicode.IsDigit(char) == false
}

func isGear(char rune) bool {
	return char == '*'
}

// checks if the part of the array is a PartNumber
func isPartNumber(arr [][]rune, pos position, length int) (bool, position) {
	retVal := false
	var partNumPosition position
	for i := 0; i < length; i++ {
		currPos := pos.Offset(i, 0)
		offsets := []position{
			currPos.Offset(0, -1),  //top
			currPos.Offset(1, -1),  //top-right
			currPos.Offset(1, 0),   //right
			currPos.Offset(1, 1),   //right-bottom
			currPos.Offset(0, 1),   //bottom
			currPos.Offset(-1, 1),  //left-bottom
			currPos.Offset(-1, 0),  //left
			currPos.Offset(-1, -1), //left-top
		}

		for _, offset := range offsets {
			if offset.IsValid() == false {
				continue
			}
			symbol := isSymbol(arr[offset.y][offset.x])
			if symbol {
				fmt.Printf("detected symbol: %q \n", arr[offset.y][offset.x])
				retVal = true
				if isGear(arr[offset.y][offset.x]) {
					partNumPosition = offset
				}

				break
			}
		}
		if retVal {
			break
		}
	}

	return retVal, partNumPosition
}

func isBeginningOfNumber(arr [][]rune, pos position) (isNumber bool, number, length int) {
	var digits []string
	length = 0
	isNumber = false
	if unicode.IsDigit(arr[pos.y][pos.x]) {
		isNumber = true
		digits = append(digits, string(arr[pos.y][pos.x]))
		length++
		i := 1
		for pos.x+i < pos.xMax && unicode.IsDigit(arr[pos.y][pos.x+i]) {
			digits = append(digits, string(arr[pos.y][pos.x+i]))
			length++
			i++
		}
	}
	if isNumber == false {
		return isNumber, number, length
	}
	var err error
	number, err = strconv.Atoi(strings.Join(digits, ""))
	if err != nil {
		panic(err)
	}
	return isNumber, number, length
}

func getPartNumbers(arr [][]rune) (numbers []int, gears map[position][]int) {
	pos := position{
		0,
		0,
		len(arr), len(arr),
	}
	gears = make(map[position][]int)
	//var gears map[position][]int
	//var numbers []int
	for i := 0; i < pos.xMax; i++ {
		numberEnd := -1
		for j := 0; j < pos.yMax; j++ {
			//continue, existing number is already checked
			if j < numberEnd {
				continue
			}
			numberEnd = -1
			pos.x = j
			pos.y = i
			fmt.Printf("current value: %q", arr[pos.y][pos.x])
			isNumber, number, length := isBeginningOfNumber(arr, pos)
			if isNumber == false {
				continue
			}
			isPartNumber, gearPosition := isPartNumber(arr, pos, length)
			if isPartNumber {
				numbers = append(numbers, number)
				numberEnd = pos.x + length
			}
			if gearPosition != (position{}) {
				gears[gearPosition] = append(gears[gearPosition], number)
			}
		}
	}
	return numbers, gears
}
func getNumberOfLines(fileName string) int {
	linesCount := 0
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linesCount++
	}
	return linesCount
}

func Day3() {
	fileName := "day3/data.txt"
	linesCount := getNumberOfLines(fileName)

	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	arr := make([][]rune, linesCount)
	i := 0
	for scanner.Scan() {
		text := scanner.Text()
		//length := len(text)
		//arr[i] = make([]rune, length)
		for j := 0; j < len(text); j++ {
			char := text[j]
			arr[i] = append(arr[i], rune(char))
		}
		i++
	}
	numbers, gears := getPartNumbers(arr)
	fmt.Println(numbers)
	sum := 0
	for _, value := range numbers {
		sum += value
	}
	var gearRatios []int
	for _, ints := range gears {
		if len(ints) != 2 {
			continue
		}
		ratio := 1
		for _, val := range ints {
			ratio *= val
		}
		gearRatios = append(gearRatios, ratio)
	}
	gearRatioSum := 0
	for _, ratio := range gearRatios {
		gearRatioSum += ratio
	}
	fmt.Printf("sum: %d", sum)
	fmt.Printf("sum: %d", gearRatioSum)
}
