package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

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
		}
		for i := length - 1; i >= 0; i-- {
			value := rune(line[i])
			if unicode.IsDigit(value) {
				end = value
				break
			}
		}
		//fmt.Printf("start: %s end: %s \n", string(start), string(end))
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
