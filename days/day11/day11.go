package day11

import (
	"adventofcode/m/v2/util"
	"fmt"
	"strings"
	"strconv"
)

func Day11(inputFile string, part int) {
	if part == 0 {
		fmt.Printf("Stones after 25 blinks: %d\n", solve(inputFile, 25))
	} else {
		fmt.Printf("Stones after 75 blinks: %d\n", solve(inputFile, 75))
	}
}

func solve(inputFile string, blinks int) int {
	ls := util.LineScanner(inputFile)
	line, _ := util.Read(ls)

	// Initialize stones, which only ever contains the input and no additional stones
	stones := strings.Split(line, " ")

	// dp.mem[b][x] = n if stone with value 'x' becomes 'n' stones in 'b' blinks
	count := 0
	dp := &DP{make(map[int]map[string]int)}
	for _, stone := range stones {
		count += dp.blink(stone, blinks)
	}

	return count
}

func (dp *DP) blink(stone string, blinks int) int {
	if dp.mem[blinks] == nil {
		dp.mem[blinks] = make(map[string]int)
	}
	
	if blinks == 0 {
		return 1
	}

	if n, ok := dp.mem[blinks][stone]; ok {
		return n
	}

	if stone == "0" {
		dp.mem[blinks][stone] = dp.blink("1", blinks-1)
	} else if len(stone) % 2 == 0 {
		first, _ := strconv.Atoi(stone[:len(stone)/2])
		second, _ := strconv.Atoi(stone[len(stone)/2:])
		dp.mem[blinks][stone] = dp.blink(strconv.Itoa(first), blinks-1) + dp.blink(strconv.Itoa(second), blinks-1)
	} else {
		x, _ := strconv.Atoi(stone)
		dp.mem[blinks][stone] = dp.blink(strconv.Itoa(x*2024), blinks-1)
	}

	return dp.mem[blinks][stone]
}

type DP struct {
	mem map[int]map[string]int
}