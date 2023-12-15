package day15

import (
	"adventofcode/m/v2/util"
	"fmt"
	"strings"
	"strconv"
)

func Day15(inputFile string, part int) {
	if part == 0 {
		fmt.Printf("Hash sum: %d\n", checkHash(inputFile))
	} else {
		fmt.Printf("Focusing power: %d\n", focusPower(inputFile))
	}
}

func checkHash(path string) int {
	ls := util.LineScanner(path)
	line, _ := util.Read(ls)
	parts := strings.Split(line, ",")
	sum := 0
	for _, s := range parts {
		sum += hash(s)
	}
	return sum
}

func focusPower(path string) int {
	ls := util.LineScanner(path)
	line, _ := util.Read(ls)
	parts := strings.Split(line, ",")

	boxes := make([]*LLenses, 256) // boxes[i] = l where l is the last lens in box i
	lenses := make(map[string]*Lens, 0)
	for _, s := range parts {
		var label string 
		var focal int
		var op rune
		// Extract label, focal and operation
		if s[len(s)-1] == '-' {
			label = s[:len(s)-1]
			op = '-'
		} else {
			pp := strings.Split(s, "=")
			label = pp[0]
			focal, _ = strconv.Atoi(pp[1])
			op = '='
		}
		// Initialize lens if needed
		if _, ok := lenses[label]; !ok {
			lenses[label] = &Lens{label, focal, nil, nil, false}
		}
		// as well as linked lens list
		h := hash(label);
		if boxes[h] == nil {
			boxes[h] = &LLenses{nil, 0}
		}

		if op == '-' {
			// Remove the lens 
			var prev *Lens
			if prev = lenses[label].prev; prev != nil {
				lenses[label].prev.next = lenses[label].next
				lenses[label].prev = nil
			}
			if lenses[label].next != nil {
				lenses[label].next.prev = prev
				lenses[label].next = nil
			}
			if lenses[label].inBox {
				lenses[label].inBox = false
				boxes[h].n--
			}

			if boxes[h].last == lenses[label] {
				boxes[h].last = prev
			}
		}

		if op == '=' {
			// Update existing
			if lenses[label].inBox {
				lenses[label].focalLength = focal
			} else { // or add at the end
				lenses[label].inBox = true
				lenses[label].focalLength = focal 
				if boxes[h].last != nil {
					boxes[h].last.next = lenses[label]
					lenses[label].prev = boxes[h].last
				}
				boxes[h].last = lenses[label]
				boxes[h].n++
			}
		}
	}
	return power(boxes)	
}

func power(boxes []*LLenses) int {
	pow := 0
	for i := 0; i < 256; i++ {
		n := 0
		if boxes[i] != nil && boxes[i].last != nil {
			l := boxes[i].last 
			pow += (i+1)*(boxes[i].n-n)*l.focalLength
			for l.prev != nil {
				l = l.prev
				n++
				pow += (i+1)*(boxes[i].n-n)*l.focalLength
			}
			n = 0
		}
	}
	return pow 
}

// Linked list to manage the lenses
type LLenses struct {
	last  	*Lens 
	n 			int
}

type Lens struct {
	label 			  string
	focalLength 	int
	prev 					*Lens
	next 					*Lens
	inBox 				bool
}

func hash(s string) int {
	h := 0
	for _, r := range s {
		h += int(r)
		h *= 17
		h = h % 256
	}
	return h 
}
