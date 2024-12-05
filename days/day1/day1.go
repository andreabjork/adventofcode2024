package day1

import (
	"adventofcode/m/v2/util"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Day1(inputFile string, part int) {
	if part == 0 {
		fmt.Printf("List distance: %d\n", distance(inputFile))
	} else {
		fmt.Printf("List similarity: %d\n", similarity(inputFile))
	}
}

func distance(inputFile string) int {
	A, B, _, _ := parse(inputFile)

	sort.Slice(A, func(i, j int) bool {
		return A[i] <= A[j]	
	})

	sort.Slice(B, func(i, j int) bool {
		return B[i] <= B[j]	
	})

	sum := 0
	for i := 0; i < len(A); i++ {
		sum += util.Abs(A[i]-B[i])
	}
	return sum
}

func similarity(inputFile string) int {
	_, _, xA, xB := parse(inputFile)

	sum := 0
	for i, _ := range xA {
		sum += i*xA[i]*xB[i]
	}
	return sum
}

func parse(inputFile string) ([]int, []int, map[int]int, map[int]int) {
	ls := util.LineScanner(inputFile)
	line, ok := util.Read(ls)
	var parts []string

	var A, B []int = []int{}, []int{}
	var mapA, mapB map[int]int = make(map[int]int), make(map[int]int)
	for ok {
		parts = strings.Split(line, "   ")
		a, _ := strconv.Atoi(parts[0])
		b, _ := strconv.Atoi(parts[1])
		A = append(A, a)
		B = append(B, b)
		mapA[a]++ 
		mapB[b]++
		line, ok = util.Read(ls)
	}
	return A, B, mapA, mapB
}


