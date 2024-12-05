package day4

import (
	"adventofcode/m/v2/util"
	"fmt"
)

func Day4(inputFile string, part int) {
	if part == 0 {
		fmt.Printf("XMAS: %d\n", xmas(inputFile, "xmas"))
	} else {
		fmt.Printf("XMAS: %d\n", xmas(inputFile, "x-mas"))
	}
}

type Point struct {
	i int 
	j int
}

func xmas(inputFile string, toFind string) int {
	ls := util.LineScanner(inputFile)
	i := 0
	line, ok := util.Read(ls) 

	// Create word chart
	words := make(map[int]map[int]rune)
	X := []*Point{}
	A := []*Point{}
	for ok {
		words[i] = make(map[int]rune)

		for j, r := range line {
			words[i][j] = r
			if r == 'X' {
				X = append(X, &Point{i,j})
			}
			if r == 'A' {
				A = append(A, &Point{i,j})
			}
		}
		i++
		line, ok = util.Read(ls)
	}
	
	sum := 0
	if toFind == "xmas" {
		for _, p := range X {
			sum += countXmas(words, p)
		}
	} else if toFind == "x-mas" {
		for _, p := range A {
			sum += countMasMas(words, p)
		}
	}
	
	return sum
}

func countMasMas(words map[int]map[int]rune, p *Point) int {
	if p.i - 1 < 0 || p.j - 1 < 0 || p.i + 1 >= len(words) || p.j + 1 >= len(words[p.i]) {
		return 0
	}

	c := 0
	// M . M
	// . A .
	// S . S
	if words[p.i-1][p.j-1] == 'M' &&
	words[p.i+1][p.j-1] == 'S' &&
	words[p.i-1][p.j+1] == 'M' &&
	words[p.i+1][p.j+1] == 'S' {
		c++
	}

  // S . M
	// . A .
	// S . M 
	if words[p.i-1][p.j-1] == 'S' &&
	words[p.i+1][p.j-1] == 'S' &&
	words[p.i-1][p.j+1] == 'M' &&
	words[p.i+1][p.j+1] == 'M' {
		c++
	}

  // S . S 
	// . A .
	// M . M 
	if words[p.i-1][p.j-1] == 'S' &&
	words[p.i+1][p.j-1] == 'M' &&
	words[p.i-1][p.j+1] == 'S' &&
	words[p.i+1][p.j+1] == 'M' {
		c++
	}

  // M . S 
	// . A .
	// M . S
	if words[p.i-1][p.j-1] == 'M' &&
	words[p.i+1][p.j-1] == 'M' &&
	words[p.i-1][p.j+1] == 'S' &&
	words[p.i+1][p.j+1] == 'S' {
		c++
	}
	return c
}

// Counts occurrences of word XMAS start at point p in all directions
func countXmas(words map[int]map[int]rune, p *Point) int {
	c := 0
	// S
	// A
	// M
	// X
	if p.i - 3 >= 0 && 
	words[p.i-3][p.j] == 'S' &&
	words[p.i-2][p.j] == 'A' &&
	words[p.i-1][p.j] == 'M' {
		c++
	}

	// X M A S
	if p.j + 3 < len(words[p.i]) &&
	words[p.i][p.j+1] == 'M' && 
	words[p.i][p.j+2] == 'A' &&
	words[p.i][p.j+3] == 'S' {
		c++
	}

	// S A M X
	if p.j - 3 >= 0 &&
	words[p.i][p.j-1] == 'M' && 
	words[p.i][p.j-2] == 'A' &&
	words[p.i][p.j-3] == 'S' {
		c++
	}

	// X
	// M
	// A
	// S
	if p.i + 3 < len(words) &&
	words[p.i+1][p.j] == 'M' && 
	words[p.i+2][p.j] == 'A' &&
	words[p.i+3][p.j] == 'S' {
		c++
	}

	// X
	// . M
	// . . A
	// . . . S
	if p.j + 3 < len(words[p.i]) && p.i + 3 < len(words) &&
	words[p.i+1][p.j+1] == 'M' && 
	words[p.i+2][p.j+2] == 'A' &&
	words[p.i+3][p.j+3] == 'S' {
		c++
	}

	// . . . X
	// . . M
	// . A
	// S
	if p.j - 3 >= 0 && p.i + 3 < len(words) &&
	words[p.i+1][p.j-1] == 'M' && 
	words[p.i+2][p.j-2] == 'A' &&
	words[p.i+3][p.j-3] == 'S' {
		c++
	}

	// . . . S
	// . . A
	// . M
	// X
	if p.j + 3 < len(words[p.i]) && p.i - 3 >= 0 &&
	words[p.i-1][p.j+1] == 'M' && 
	words[p.i-2][p.j+2] == 'A' &&
	words[p.i-3][p.j+3] == 'S' {
		c++
	}

	// S
	// . A
	// . . M
	// . . . X
	if p.j - 3 >= 0 && p.i - 3 >= 0 &&
	words[p.i-1][p.j-1] == 'M' && 
	words[p.i-2][p.j-2] == 'A' &&
	words[p.i-3][p.j-3] == 'S' {
		c++
	}

	return c
}
