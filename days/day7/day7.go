package day7

import (
	"adventofcode/m/v2/util"
	"fmt"
	"strconv"
	"strings"
)

func Day7(inputFile string, part int) {
	if part == 0 {
		fmt.Printf("Total winnings: %d\n", camelCards(inputFile, false))
	} else {
		fmt.Printf("Total winnings: %d\n", camelCards(inputFile, true))
	}
}

func camelCards(inputFile string, allowJokers bool) int {
	ls := util.LineScanner(inputFile)
	line, ok := util.Read(ls)

	// Set the value of the Joker to lowest if we're using them as a free-for-all
	if allowJokers {
		cardValue['J'] = 1
	}

	hands := []*Hand{}
	for ok {
		parts := strings.Split(line, " ")
		cards := []rune(parts[0])
		cardTypes := make(map[rune]bool)
		bid, _ := strconv.Atoi(parts[1])

		// Count card types in order to consider all J-variations
		for _, c := range cards {
			cardTypes[c] = true
		}

		variations := [][]rune{} 
		if allowJokers && cardTypes['J'] {
			for t := range cardTypes {
				if t != 'J' {
					variations = append(variations, replaceJWith(cards, t))
				}
			}
			// If we have ONLY Js then we treat them as aces:
			if len(cardTypes) == 1 {
				variations = append(variations, replaceJWith(cards, 'A'))
			}
		} else {
			variations = [][]rune{cards} 
		}

		var best int = 0
		for _, v := range variations {
			count := make(map[rune]int)
			max := 0
			for _, card := range v {
				count[card]++
				if count[card] > max {
					max = count[card]
				}
			}
			unique := len(count)
			max_unique, _ := strconv.Atoi(strconv.Itoa(max) + strconv.Itoa(unique))

			if max_unique > best {
				best = max_unique
			}
		}
		hand := &Hand{cards, 0, bid}
		switch best {
			case 51: // five of a kind 
			hand.typ = 7 
			case 42: // four of a kind
			hand.typ = 6
			case 32: // full house
			hand.typ = 5 
			case 33: // 3 of a kind
			hand.typ = 4
			case 23: // 2 pairs
			hand.typ = 3
			case 24: // 2 of a kind
			hand.typ = 2
			case 15: // high card
			hand.typ = 1
			default:
				panic("invalid hand")
		}
		
		hands = insertSorted(hands, hand)
		line, ok = util.Read(ls)
	}
	
	// Sum up the rank multiples
	sum := 0
	for i, h := range hands {
		rank := i+1
		sum += h.bid*rank
	}

	return sum
}

type Hand struct {
	cards []rune
	typ 	int 
	bid   int
}

var cardValue = map[rune]int {
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'J': 11,
	'Q': 12,
	'K': 13,
	'A': 14,
}

// Returns the hand of the higher rank
func (h *Hand) rank(k *Hand) *Hand {
	if h.typ == k.typ {
		for i := range h.cards {
			if cardValue[h.cards[i]] > cardValue[k.cards[i]] {
				return h
			} else if cardValue[k.cards[i]] > cardValue[h.cards[i]] {
				return k
			}
		}
	} else if h.typ > k.typ {
		return h
	}

	return k
}

func insertSorted(hands []*Hand, k *Hand) []*Hand {
	var cutoff int = len(hands)
	for i, h := range hands {
		if h.rank(k) == h {
			cutoff = i
			break
		} 
	}

	head := []*Hand{}
	tail := []*Hand{}
	if cutoff > 0 {
		head = hands[:cutoff]
	}

	if cutoff < len(hands) {
		tail = hands[cutoff:]
	}
	return append(head, append([]*Hand{k}, tail...)...)
}

func replaceJWith(cards []rune, r rune) []rune {
	cc := make([]rune, len(cards))
	for i := range cards {
		if cards[i] == 'J' {
			cc[i] = r
		} else {
			cc[i] = cards[i]
		}
	}

	return cc
}

