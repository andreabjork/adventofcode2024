package day5

import (
	"adventofcode/m/v2/util"
	"fmt"
	"strconv"
	"strings"
)

func Day5(inputFile string, part int) {
	if part == 0 {
		fmt.Printf("Sum of middles: %d\n", middles(inputFile, false))
	} else {
		fmt.Printf("Sum of middles: %d\n", middles(inputFile, true))
	}
}

type Link struct {
	val 	int
	prev  *Link 
	next  *Link
	max   int
}

func middles(inputFile string, fix bool) int {
	ls := util.LineScanner(inputFile)
	line, ok := util.Read(ls)

	// a -> b if a must come before b (rule a|b)
	rules := make(map[int][]int)	

	for line != "" {
		r := strings.Split(line, "|")
		a,_ := strconv.Atoi(r[0])
		b,_ := strconv.Atoi(r[1])
		
		rules[a] = append(rules[a], b)
		line, ok = util.Read(ls)
	}

	// Create linked list
	line, ok = util.Read(ls)
	var middle int
	for ok {
		pages := strings.Split(line, ",")
		var first, curr *Link
		for i := 0; i < len(pages); i++ {
			pp, _ := strconv.Atoi(pages[i])
			new := &Link{pp, curr, nil, len(pages)}
			if i == 0 {
				first = new
			}
			if curr != nil {
				curr.next = new 
			}
			curr = new
		}

		change := checkOrder(first, rules, fix)
		for first.prev != nil {
			first = first.prev
		}

		if (fix && change) || (!fix && !change) {
			for i := 0; i < first.max/2; i++ {
				first = first.next
			}
			middle += first.val
		}
		line, ok = util.Read(ls)
	}
	return middle
}

func (l *Link) get() int {
	if l != nil {
		return l.val 
	} else {
		return -1
	}
}

// Fixes the order only if fix = true
func checkOrder(a *Link, rules map[int][]int, fix bool) bool {
	orderChanged := false
	seen := make(map[int]*Link)
	for a != nil {
		for _, B := range rules[a.val] {
			if b := seen[B]; b != nil {
				orderChanged = true
				if !fix {
					return orderChanged
				}
				if a.next != nil {
					a.next.prev = b
				}
				if b.prev != nil {
					b.prev.next = b.next
				}
				if b.next != nil {
					b.next.prev = b.prev 
				}

				b.next = a.next	
				a.next = b 
				b.prev = a 
				seen[B] = nil
			}
		}
		seen[a.val] = a
		a = a.next
	}
	return orderChanged
} 

// For debugging only
func print(a *Link) {
	saveFirst := a
	// find the first link
	for a.prev != nil {
		a = a.prev
	}
	for a != nil {
		fmt.Printf("%d ", a.val)
		a = a.next
		
	}
	fmt.Println()
	a = saveFirst
}