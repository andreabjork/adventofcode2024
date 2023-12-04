package day4

import (
	"adventofcode/m/v2/util"
	"fmt"
	"regexp"
	"strings"
	"strconv"
)

func Day4(inputFile string, part int) {
	if part == 0 {
		pts, _ := points(inputFile)
		fmt.Printf("Points: %d\n", pts)
	} else {
		_, cards := points(inputFile)
		fmt.Printf("Number of cards: %d\n", cards)
	}
}

func points(inputFile string) (int, int) {
	ls := util.LineScanner(inputFile)
	line, ok := util.Read(ls) 

	sum := 0
	numCards := 0
	parseCard := regexp.MustCompile(`Card\s+(\d+):\s+([^|]*)\s+\|\s+([^|]*)`)
	replace := regexp.MustCompile(`\s\s+`)
	var cards map[int]int = make(map[int]int)
	for ok {
		matches := parseCard.FindAllStringSubmatch(line, 3)
		card, _ := strconv.Atoi(matches[0][1])
		cards[card]++
		wins := strings.Split(replace.ReplaceAllString(matches[0][2], " "), " ")
		nums := strings.Split(replace.ReplaceAllString(matches[0][3], " "), " ")
		// Mark winning numbers
		isWinningNumber := make(map[int]bool)
		for _, w := range wins {
			ww, _ := strconv.Atoi(w)
			isWinningNumber[ww] = true
		}

		// Count number of wins and score this card
		total := 0
		pts := 0 
		for _, n := range nums {
			nn, _ := strconv.Atoi(n)
			if isWinningNumber[nn] {
				if pts == 0 {
					pts++
				} else {
					pts*=2
				}
				total++
			}
		}

		// For every copy of this card, we score more cards
		for c := card+1; c <= card+total; c++ {
			cards[c]+=cards[card]
		}
	
		line, ok = util.Read(ls)
		sum += pts
		numCards += cards[card]
	}

	return sum, numCards
}
