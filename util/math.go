package util

import (
	"strconv"
)
const MaxUint = ^uint(0) 
const MinUint = 0 
const MaxInt = int(MaxUint >> 1) 
const MinInt = -MaxInt - 1

func Max(x int, y int) int {
	if x >= y {
	return x
	}
	return y
}

func Min(x int, y int) int {
	if x <= y {
	return x
	}
	return y
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Sum(arr []int) int {
	sum := 0
	for _, a := range arr {
		sum += a
	}
	return sum
}

func Sign(i int) int {
	if i > 0 {
		return 1
	} else if i < 0 {
		return -1
	} else {
		return 0
	}
}

func CopySlice(slice []string) []string {
	newSlice := []string{}
	for _, s := range slice {
		newSlice = append(newSlice, s)	
	}

	return newSlice
}

func Pow(x int, n int) int {
	val := 1
	for i := 0; i < n; i++ {
		val *= x
	}
	return val
}

func AsInts(strs []string) []int {
	var ints []int = []int{}
	for _, s := range strs {
		i, err := strconv.Atoi(s)
		if err == nil {
			ints = append(ints, i)
		}
	}
	return ints
}

func AsFloats(strs []string) []float64 {
	var floats []float64 = []float64{}
	for _, s := range strs {
		i, err := strconv.Atoi(s)
		if err == nil {
			floats = append(floats, float64(i))
		}
	}
	return floats
}