package day5

import (
	"adventofcode/m/v2/util"
	"fmt"
	"strconv"
	"strings"
)

func Day5(inputFile string, part int) {
	if part == 0 {
		fmt.Printf("Lowest location: %d\n", plantSeeds(inputFile, false))
	} else {
		fmt.Printf("Lowest location: %d\n", plantSeeds(inputFile, true))
	}
}

func plantSeeds(inputFile string, useRange bool) int {
	seeds, maps := parse(inputFile)
	if !useRange {
		// Overwrite the seed range to 1
		newSeeds := []int{}
		for _, s := range seeds {
			newSeeds = append(newSeeds, s, 1)
		}
		seeds = newSeeds 
	}

	// Create chain seed-to-soil,-to-fertilizer,-to-water,-to... 
	var chain [](func(s []*Segment) []*Segment)
	for _, m := range maps {
		chain = append(chain, fs(m))
	}

	// Map every segment instead of mapping individual seeds
	segs := []*Segment{}
	for i := 0; i < len(seeds)-1; i += 2 {
		r := seeds[i+1]
		s := &Segment{seeds[i], seeds[i] + r}
		newSegs := []*Segment{s}
		for _, c := range chain {
			newSegs = c(newSegs)
		}

		segs = append(segs, newSegs...)
	}

	// Find the global minima
	min := util.MaxInt
	for _, s := range segs {
		if s.a < min {
			min = s.a
		} 
	}
		
	return min
}

type Segment struct {
	a int
	b int
}

func fs(def [][]int) func(segments []*Segment) []*Segment {
	return func(segs []*Segment) []*Segment {
		segments := []*Segment{}
		for _, s := range segs {
			aOutside := true 
			bOutside := true
			var a, b int = -1, -1
			for _, m := range def { // m[0] destination, m[1] source, m[2] range
				// We either map the entire segment
				if s.a >= m[1] && s.a <= m[1]+m[2] && s.b >= m[1] && s.b <= m[1]+m[2] {
					a = m[0] + (s.a - m[1])
					b = m[0] + (s.b - m[1])
					aOutside = false
					bOutside = false
				// or map part of it and consider the remaining segment 
				} else if s.a >= m[1] && s.a <= m[1]+m[2] {
					a = m[0] + (s.a - m[1])
					b = m[0] + m[2]
					s.a = m[1]+m[2] 
					aOutside = false
				} else if s.b >= m[1] && s.b <= m[1]+m[2] {
					a = m[0]
					b = m[0] + (s.b - m[1])
					s.b = a 
					bOutside = false
				} 

				if a > 0 && b > 0 {
					segments = append(segments, &Segment{a, b})
					a = -1 
					b = -1 
				}
			}

			// Identity mapping for any part of segment that falls outside 
			if aOutside || bOutside {
				segments = append(segments, &Segment{s.a, s.b})
			}

		}
		return segments
	}
}

func parse(inputFile string) ([]int, [][][]int) {
	ls := util.LineScanner(inputFile)
	seeds, ok := util.Read(ls)
	util.Read(ls)

	var maps [][][]int = make([][][]int, 0)
	var rules [][]int
	var vals []int
	var line string
	i := 0
	for ok {
		util.Read(ls)
		line, ok = util.Read(ls)
		rules = make([][]int, 0)
		for line != "" && ok {
			vals = asInts(strings.Split(line, " "))
			rules = append(rules, vals)
			line, ok = util.Read(ls)
		}
		i++
		maps = append(maps, rules)

	}

	return asInts(strings.Split(seeds, " ")), maps
}

func asInts(strs []string) []int {
	var ints []int = []int{}
	for _, s := range strs {
		i, err := strconv.Atoi(s)
		if err == nil {
			ints = append(ints, i)
		}
	}
	return ints
}