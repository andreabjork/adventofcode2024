package day13

import (
	"adventofcode/m/v2/util"
	"fmt"
	"strings"
	"strconv"
)

func Day13(inputFile string, part int) {
	if part == 0 {
		fmt.Printf("Tokens: %d\n", solve(inputFile, false))
	} else {
		fmt.Printf("Tokens: %d\n", solve(inputFile, true))
	}
}

func solve(inputFile string, unlimited bool) int {
	ls := util.LineScanner(inputFile)

	var first, second, prize string 
	var x, y int 
	var A, B, C *Vec
	var pts []string
	ok := true
	tokens := 0
	for ok {
		first, _ = util.Read(ls)
		second, _ = util.Read(ls)
		prize, _ = util.Read(ls)

		pts = strings.Split(strings.TrimPrefix(first, "Button A: "), ", ")
		x, _ = strconv.Atoi(strings.TrimPrefix(pts[0], "X+"))
		y, _ = strconv.Atoi(strings.TrimPrefix(pts[1], "Y+"))
		A = &Vec{x, y}

		pts = strings.Split(strings.TrimPrefix(second, "Button B: "), ", ")
		x, _ = strconv.Atoi(strings.TrimPrefix(pts[0], "X+"))
		y, _ = strconv.Atoi(strings.TrimPrefix(pts[1], "Y+"))
		B = &Vec{x, y}

		pts = strings.Split(strings.TrimPrefix(prize, "Prize: "), ", ")
		x, _ = strconv.Atoi(strings.TrimPrefix(pts[0], "X="))
		y, _ = strconv.Atoi(strings.TrimPrefix(pts[1], "Y="))
		if unlimited {
			C = &Vec{10000000000000+x, 10000000000000+y}
		} else {
			C = &Vec{x, y}
		}

		n, m := eqSolver(A, B, C)
		fmt.Printf("n, m = %d, %d ", n, m)
		if n >= 0 && m >= 0 && (unlimited || n <= 100 && m <= 100) {
			fmt.Printf("\t\tTokens %d\n", 3*n+m)
			tokens += 3*n + m	
		} else {
			fmt.Printf("\t\tNO WIN\n")
		}

		_, ok = util.Read(ls)
	}
	fmt.Println()
	return tokens
}

type Vec struct {
	x, y int
}

// Solves the linear algebra for
//
// n*A + m*B = C 
//
// where n is the number of button presses for button A,
//       m is the number of button presses for button B
//       C is the Vector that contains the end locations of x, y
//       A, B are the vectors that describe the impact of pressing buttons A, B respectively.
func eqSolver(A, B, C *Vec) (int, int) {
	// n*A_1 + m*B_1 = C_1 => n = (C_1 - m*B_1)/A_1
	// n*A_2 + m*B_2 = C_2 => n = (C_2 - m*B_2)/A_2
	//
	// thus
	// A_2*(C_1 - m*B_1) = A_1*(C_2 - m*B_2)
	// =>
	// m = (A_2*C_1 - A_1*C_2) / (A_2*B_1 - A_1*B_2)
	// n = (C_1 - m*B_1)/A_1

	var n, m int
	div := A.y * B.x - A.x*B.y 
	if div == 0 || A.x == 0 || A.y == 0 {
		// no solution
		return -1, -1
	} else {
		m = (A.y*C.x - A.x*C.y)/div
		n = (C.x - m*B.x) / A.x
		
		// Note that n, m are integers as they must be. However this means
		// the solution must be checked:
		if n*A.x + m*B.x != C.x || n*A.y + m*B.y != C.y {
			return -1, -1
		}
	}

	return n, m
}