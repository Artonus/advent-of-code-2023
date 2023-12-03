package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var digits = map[string]rune{"one": '1', "two": '2', "three": '3', "four": '4', "five": '5', "six": '6', "seven": '7', "eight": '8', "nine": '9'}

func isDigit(text string, reverse bool) (bool, rune) {
	for key, digit := range digits {
		if reverse {
			if strings.HasSuffix(text, key) {
				return true, digit
			}
		} else {
			if strings.HasPrefix(text, key) {
				return true, digit
			}
		}

	}
	return false, 0
}

func Day1() {

	file, err := os.Open("day1/data.txt")
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}
	scanner := bufio.NewScanner(file)
	var values []int
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		line := scanner.Text()
		length := len(line)
		var start rune
		var end rune
		for i := 0; i < length; i++ {
			value := rune(line[i])
			if unicode.IsDigit(value) {
				start = value
				break
			}
			isDigit, digit := isDigit(line[i:], false)
			if isDigit {
				start = digit
				break
			}
		}
		for j := length - 1; j >= 0; j-- {
			value := rune(line[j])
			if unicode.IsDigit(value) {
				end = value
				break
			}
			isDigit, digit := isDigit(line[:j+1], true)
			if isDigit {
				end = digit
				break
			}
		}
		fmt.Printf("start: %s end: %s \n", string(start), string(end))
		//fmt.Sprintf("%s%s", start, end)
		intVal, err := strconv.Atoi(strings.Join([]string{string(start), string(end)}, ""))
		if err != nil {
			fmt.Printf("error: %s", err.Error())
			return
		}
		fmt.Printf("value = %d \n", intVal)
		values = append(values, intVal)
	}
	sum := 0
	for _, value := range values {
		sum += value
	}
	fmt.Printf("sum: %d", sum)
}
