package day18

import (
	"adventofcode/m/v2/util"
	"fmt"
	"strconv"
	"strings"
)

func Day18(inputFile string, part int) {
	if part == 0 {
		fmt.Printf("Surface area: %d\n", shoelace(inputFile, false))
	} else {
		fmt.Printf("Surface area: %d\n", shoelace(inputFile, true))
	}
}

// Not my original idea, but easy to implement.
// https://www.reddit.com/r/adventofcode/comments/18l2tap/2023_day_18_the_elves_and_the_shoemaker/
// https://en.wikipedia.org/wiki/Shoelace_formula
//
// We sum up all triangles formed by adding a new edge,
// with Right, Up being positive areas and Down, Left being negative
//
// The triangles are formed via vectors
//
//           ^ (x2,y2) = w
//          /
//         /
// -----> / (x1,y1) = v
//
// Where the area of the parallelogram formed by these vectors is
// | v x w | = x1*y2 - x2*y1
//
// We consider 1/2 | v x w | since we want only half of the parallelogram.
func shoelace(inputFile string, useHex bool) int {
	ls := util.LineScanner(inputFile)
	line, ok := util.Read(ls)

	var x2, y2 float64
	var x1, y1 float64 = 0, 0
	var sum float64
	var tr int = 0
	for ok {
		parts := strings.Split(line, " ")
		dir := []rune(parts[0])
		stepss, _ := strconv.Atoi(parts[1])
		steps := float64(stepss)

		if useHex {
			color := parts[2]
			st := color[2:len(color)-2]
			d := rune(color[len(color)-2])
			steps = float64(hexToDec(st))
			switch d {
			case '0':
				dir = []rune{'R'}
			case '1':
				dir = []rune{'D'}
			case '2':
				dir = []rune{'L'}
			case '3':
				dir = []rune{'U'}
			}
		}

		switch dir[0] {
		case 'U':
			x2 = x1 - steps
			y2 = y1
		case 'D':
			x2 = x1 + steps 
			y2 = y1
		case 'L':
			x2 = x1
			y2 = y1 - steps
		case 'R':
			x2 = x1
			y2 = y1 + steps
		}
		tr += int(steps)
		sum += 0.5*(x1*y2 - x2*y1)
		x1, y1 = x2, y2

		line, ok = util.Read(ls)
	}

	// Pick's theorem
	return util.Abs(int(sum))+tr/2+1
}

func hexToDec(hex string) int {
   n, err := strconv.ParseInt(hex, 16, 64)

	 if err != nil {
		 panic(err)
	 }
	 return int(n)
}
