package day07

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	HighCard = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

var cards = []rune{'J', '2', '3', '4', '5', '6', '7', '8', '9', 'T', 'Q', 'K', 'A'}

type hand struct {
	hand     string
	bid      int
	handType int
}

func getHandScore(hand string) int {
	cardsCount := make(map[rune]int)
	jCount := 0
	for _, card := range hand {
		if rune(card) == 'J' {
			jCount++
			continue
		}
		cardsCount[card] += 1
	}
	maxVal := 0
	for _, i := range cardsCount {
		maxVal = max(maxVal, i)
	}
	maxVal += jCount

	if len(cardsCount) == 1 {
		return FiveOfAKind
	}
	if len(cardsCount) == 2 && maxVal == 4 {
		return FourOfAKind
	}
	if len(cardsCount) == 2 && maxVal == 3 {
		return FullHouse
	}
	if len(cardsCount) == 3 && maxVal == 3 {
		return ThreeOfAKind
	}
	if len(cardsCount) == 3 && maxVal == 2 {
		return TwoPair
	}
	if len(cardsCount) == 4 {
		return OnePair
	}
	if len(cardsCount) == 5 {
		return HighCard
	}
	return HighCard
}
func indexOf(element uint8, data []rune) int {
	for k, v := range data {
		if rune(element) == v {
			return k
		}
	}
	return -1 //not found.
}
func isHandBiggerThan(i, j hand) bool {
	for ii := 0; ii < len(i.hand); ii++ {
		if i.hand[ii] == j.hand[ii] {
			continue
		}
		iIdx := indexOf(i.hand[ii], cards)
		jIdx := indexOf(j.hand[ii], cards)
		return iIdx < jIdx
	}
	return false
}
func orderHands(hands []hand) []hand {
	sort.Slice(hands, func(i, j int) bool {
		return hands[i].handType < hands[j].handType || (hands[i].handType == hands[j].handType && isHandBiggerThan(hands[i], hands[j]))
	})
	return hands
}
func readHands(fileName string) []hand {
	file, err := os.Open(fileName)

	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	var hands []hand
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		haand := split[0]
		bidStr := split[1]

		bid, errParse := strconv.Atoi(bidStr)
		if errParse != nil {
			panic(err)
		}
		handScore := getHandScore(haand)
		hnd := hand{
			hand:     haand,
			bid:      bid,
			handType: handScore,
		}
		hands = append(hands, hnd)
	}

	return hands
}
func Day7() {
	fileName := "day07/data.txt"
	hands := readHands(fileName)
	for _, hnd := range hands {
		fmt.Printf("hand: %s, bid: %d, score: %d \n", hnd.hand, hnd.bid, hnd.handType)
	}
	ordered := orderHands(hands)

	totalWinnings := 0
	for i := 0; i < len(ordered); i++ {
		winning := ordered[i].bid * (i + 1)
		totalWinnings += winning
		fmt.Printf("rank: %d, hand: %s \n", i, ordered[i].hand)
	}

	fmt.Printf("total wins: %d", totalWinnings)
}
