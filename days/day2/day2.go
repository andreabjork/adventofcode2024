package day2

import (
	"adventofcode/m/v2/util"
	"fmt"
	"strings"
	"strconv"
)

func Day2(inputFile string, part int) {
	if part == 0 {
		fmt.Printf("Safe reports: %d\n", safeReports(inputFile, 0))
	} else {
		fmt.Printf("Safe reports, dampened: %d\n", safeReports(inputFile, 1))
	}
}

func safeReports(inputFile string, threshold int) int {
	ls := util.LineScanner(inputFile)
	report, ok := util.Read(ls)
	var levels []string
	var numSafe int = 0
	for ok {
		levels = strings.Split(report, " ")
		if isSafe(levels, threshold) {
			numSafe++
		}
		report, ok = util.Read(ls)
	}

	return numSafe
} 

// 0 <= idx <= len(slice)
// Returns true if level at idx does not break safety requirements
func safeAtIndex(slice []string, idx int) bool {
	safe := true
	x, _ := strconv.Atoi(slice[idx])	
	var a, b int

	if idx > 0 {
		a, _ = strconv.Atoi(slice[idx-1])
		safe = util.Abs(a-x) >= 1 && util.Abs(a-x) <= 3 
	}
	
	if idx+1 < len(slice) {
		b, _ = strconv.Atoi(slice[idx+1])
		safe = safe && util.Abs(x-b) >= 1 && util.Abs(x-b) <= 3
	}

	if idx > 0 && idx+1 < len(slice) {
		safe = safe && util.Sign(a-x) == util.Sign(x-b) && util.Sign(a-x) != 0
	}

	return safe
}

func isSafe(slice []string, threshold int) bool {
	//fmt.Printf("Checking %s // ", slice)
	safe := true
	for i := 0; i < len(slice); i++ {
		safe = safe && safeAtIndex(slice, i)
		if !safe && threshold > 0 {
			// we can try to remove the unsafe level and check again	
			threshold--
			return isSafe(remove(slice, i), threshold) || 
			isSafe(remove(slice, i+1), threshold) ||
			(i > 0 && isSafe(remove(slice, i-1), threshold))
		} else if !safe {
			break
		}
	}
	return safe
}

func remove(slice []string, idx int) []string {
	if idx == 0 {
		return slice[idx+1:]
	} else if idx == len(slice)-1 {
		return slice[:idx]
	} else {
		newSlice := util.CopySlice(slice[:idx])
		return append(newSlice, slice[idx+1:]...)
	}
}