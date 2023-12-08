package day5

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type mapper struct {
	min    int
	max    int
	offset int
}

func (m *mapper) IsValueInRange(value int) bool {
	return value >= m.min && value <= m.max
}
func (m *mapper) Map(value int) (success bool, converted int) {
	if value >= m.min && value <= m.max {
		return true, value + m.offset
	}
	return false, 0
}

type Converter struct {
	Ranges      []mapper
	Destination string
}

func (c *Converter) Convert(value int) int {
	converted := false
	retVal := value
	for i := 0; i < len(c.Ranges); i++ {
		rg := c.Ranges[i]
		if rg.IsValueInRange(value) {
			_, conv := rg.Map(value)
			retVal = conv
			converted = true
			break
		}
	}
	if converted {
		return retVal
	}
	return value

}
func parseRange(line string) mapper {
	split := strings.Split(line, " ")
	beginning, err := strconv.Atoi(strings.TrimSpace(split[1]))
	if err != nil {
		panic(err)
	}
	end, err := strconv.Atoi(strings.TrimSpace(split[0]))
	if err != nil {
		panic(err)
	}
	count, err := strconv.Atoi(strings.TrimSpace(split[2]))
	if err != nil {
		panic(err)
	}

	return mapper{
		min:    beginning,
		max:    beginning + count - 1, //-1 because it must include the beginning itself
		offset: end - beginning,
	}
}

func getSeeds(line string) []int {
	split := strings.Split(line, ":")
	numStrings := strings.Split(split[1], " ")

	var numbers []int
	for _, numString := range numStrings {
		if len(numString) == 0 {
			continue
		}
		num, err := strconv.Atoi(strings.TrimSpace(numString))
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, num)
	}
	return numbers
}
func getDestination(line string) string {
	cleared := strings.Replace(line, " map:", "", 1)
	split := strings.Split(cleared, "-to-")
	return split[1]
}
func convert(value int, node *ConverterNode) int {
	converter := node.Value
	retVal := converter.Convert(value)
	if node.Next != nil {
		return convert(retVal, node.Next)
	}
	return retVal
}
func Day5() {

	fileName := "day5/data.txt"
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	list := LinkedList{}
	var currConverter *Converter
	var seeds []int
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 && currConverter != nil {
			list.Append(currConverter)
			continue
		}
		if strings.HasPrefix(line, "seeds:") {
			seeds = getSeeds(line)
			continue
		}
		if strings.HasSuffix(line, "map:") {
			destination := getDestination(line)
			currConverter = &Converter{
				Destination: destination,
			}
			continue
		}
		if len(line) > 0 {
			rg := parseRange(line)
			currConverter.Ranges = append(currConverter.Ranges, rg)
			continue
		}
	}
	list.Append(currConverter)
	//todo: convert seeds
	var converted []int
	minim := math.MaxInt32
	for _, seed := range seeds {
		head := list.Head
		val := convert(seed, head)
		converted = append(converted, val)
		minim = min(val, minim)

	}
	fmt.Printf("minimal: %d", minim)
}
