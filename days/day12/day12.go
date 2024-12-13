package day12

import (
	"adventofcode/m/v2/util"
	"fmt"
)


func Day12(inputFile string, part int) {
	if part == 0 {
		price, _ := calculatePrice(inputFile, false)
		fmt.Printf("Price for plots: %d\n", price)
	} else {
		_, discountedPrice := calculatePrice(inputFile, false)
		fmt.Printf("Price for plots: %d\n", discountedPrice)
	}
}

func calculatePrice(inputFile string, runAssertion bool) (int, int) {
	ls := util.LineScanner(inputFile)
	line, ok := util.Read(ls)
	
	// Parse input
	lines := ""
	i := 0
	g := &Garden{make(map[int]map[int]*SinglePlot)}
	for ok {
		for j, r := range line {
			if g.plot[i] == nil {
				g.plot[i] = make(map[int]*SinglePlot)
			}
			g.plot[i][j] = &SinglePlot{r, false, i, j, nil}
		}
		i++
		lines += line+"\n"
		line, ok = util.Read(ls)
	}

	// For every plant not belonging to a plot, create a region
	// and discover the entire region
	price := 0
	discountedPrice := 0
	for i := 0; i < len(g.plot); i++ {
		for j := 0; j < len(g.plot[i]); j++ {
			if !g.plot[i][j].added {
				g.plot[i][j].added = true
				region := &Region{[]*SinglePlot{g.plot[i][j]}, 1, g.perimeter(i,j), 4}
				g.plot[i][j].region = region
				g.discover(i, j, region)
				price += region.area*region.perimeter
				discountedPrice += region.area*region.sides
			}
		}	
	}

	return price, discountedPrice
}

type Garden struct {
	plot map[int]map[int]*SinglePlot
}

type SinglePlot struct {
	plant rune 
	added bool
	i, j int
	region *Region
}

type Region struct {
	plants []*SinglePlot
	area int 
	perimeter int 
	sides int
}

// Discovers all plants matching g.plot[i][j] and adds it to region
func (g *Garden) discover(i, j int, reg *Region) {
	for _, q := range [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}} { // visits all adjacent plots
		if i+q[0] >= 0 && 
		i+q[0] < len(g.plot) && 
		j+q[1] >= 0 && 
		j+q[1] < len(g.plot[i+q[0]]) { // (i+r, j+s) is within bounds, and 
			// the plant type is the same and it doesn't belong to a region yet
			if next := g.plot[i+q[0]][j+q[1]]; next.plant == g.plot[i][j].plant && !next.added {
			  next.added = true
				next.region = reg
				reg.plants = append(reg.plants, next)
				reg.area++
				reg.perimeter += g.perimeter(i+q[0], j+q[1])
				reg.sides += g.cornerDelta(i+q[0], j+q[1], reg)
				g.discover(i+q[0], j+q[1], reg)
			} 
		}
	}
}

// x = perimeter(i,j) if adding (i,j) to the region adds x to the perimeter
func (g *Garden) perimeter(i, j int) int {
	perimeter := 4
	for _, q := range [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}} {
		if i+q[0] >= 0 && 
		i+q[0] < len(g.plot) && 
		j+q[1] >= 0 && 
		j+q[1] < len(g.plot[i+q[0]]) && // (i+r, j+s) is within bounds, and 
		g.plot[i+q[0]][j+q[1]].plant == g.plot[i][j].plant { // the plant type is the same
			perimeter--
		}
	}

	return perimeter
}

// computes the corner delta of adding plant at [i][j] to region r
func (g *Garden) cornerDelta(i, j int, r *Region) int {
	cornerDelta := 0
	p := g.plot[i][j].plant
	A, B, C := false, false, false
	for _, dir := range [][]int{{-1,-1}, {1,1}, {-1,1}, {1,-1}}  {
		A = g.has(i+dir[0], j, p, r)
		B = g.has(i, j+dir[1], p, r)
		C = g.has(i+dir[0], j+dir[1], p, r)
		cornerDelta += g.deltaForOneCorner(A, B, C)	
	}

	return cornerDelta
}

// true if g[i][j] has plant 'p' contained within region 'r'
// false if i, j out of bounds
func (g *Garden) has(i, j int, plant rune, r *Region) bool {
	if i >= 0 && j >= 0 && i < len(g.plot) && j < len(g.plot[i]) &&
	g.plot[i][j].plant == plant &&
	g.plot[i][j].region == r {
			return true 
	}
	return false 
}

func (g *Garden) deltaForOneCorner(A, B, C bool) int {
	// When adding a new plant [+] we will consider every corner area of 4 plants 
	// ([+] and the 3 possible neighbours at one corner) and determine whether any
	// corners get added or covered from this addition. 
	// The delta in number of corners is the same as the delta in number of sides 
	//
	// Imagine we add
	//
	// a  b  c
	// d [+] f 
	// g  h  i
	//
	// Where a, b, ... i may be a plant or it may be empty / a different region.
	//
	// We will then consider the regions
	// 
	// a  b    b  c   d [+]  [+] f
	// d [+]  [+] f   g  h    h  i
	// 
	// all separately and tally up the number of corners added/covered.
	//
	// For the BOTTOM RIGHT corner, we see that the following applies
	//
	// 
	// +1 corner added:
	// [+] -   [+] -    [+] -  [+] #   
	//  -  -    -  #     #  #   -  # 
	//
	// -1 corner added:
	// [+] #   [+] -  [+] #  [+] #
	//  -  -    #  -   #  #   #  -
	//	
	// These cases also cover adding [+] along either edge or in a corner
	// so long as A, B, C == false for any adjacent out of range.
	// 
	// By symmetry we can call the direct adjacents A and B, and the 
	// corner adjacent C. We then get:
	var corners int = 0
	if (!A && !B && !C) || 
	(!A && !B && C) || 
	(A && C && !B) || 
	(B && C && !A) {
		corners++
	}

	if (A && !B && !C) || 
	(B && !A && !C) || 
	(A && B && !C) || 
	(A && B && C) {
		corners--
	}
	return corners 
}
 


