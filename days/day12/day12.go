package day12

import (
	"adventofcode/m/v2/util"
	"fmt"
	"strings"
	"strconv"
	//"bufio"
	//"os"
)

func Day12(inputFile string, part int) {
	if part == -1 {
		solve(inputFile, true)
	} else if part == 0 {
		price, _ := solve(inputFile, false)
		fmt.Printf("Price for plots: %d\n", price)
	} else {
		_, discountedPrice := solve(inputFile, false)
		fmt.Printf("Price for plots: %d\n", discountedPrice)
	}
}

func solve(inputFile string, runAssertion bool) (int, int) {
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

		if runAssertion && strings.HasPrefix(line, "assert ") {
			parts := strings.Split(line, " ")
			expect, _ := strconv.Atoi(parts[1])

			fmt.Printf("Running test:\n%s", lines)
			_, discounted := g.calculatePrice()
			if expect == discounted {
				fmt.Printf("\tOK\n\n")
			} else {
				fmt.Printf("\tFAIL: Expected %d got %d\n\n", expect, discounted)
			}

			i = 0
			g = &Garden{make(map[int]map[int]*SinglePlot)}
			lines = ""

			line, ok = util.Read(ls)
			line, ok = util.Read(ls)
		}
	}

	return g.calculatePrice()
}

func (g *Garden) print(r *Region) {
	var Green = "\033[32m" 
	var Reset = "\033[0m" 
	for i := 0; i < len(g.plot); i++ {
		for j := 0; j < len(g.plot[i]); j++ {
			if g.plot[i][j].region == r {
				fmt.Printf(Green+string(g.plot[i][j].plant)+Reset)
			} else {
				fmt.Printf(string(g.plot[i][j].plant))
			}
		}
		fmt.Println()
	}
}

func (g *Garden) calculatePrice() (int, int) {
	price := 0
	discountedPrice := 0
	for i := 0; i < len(g.plot); i++ {
		for j := 0; j < len(g.plot[i]); j++ {
			if !g.plot[i][j].added {
				//fmt.Printf("   Creating region {%s} starting at (%d,%d)\n", string(g.plot[i][j].plant), i, j)
				g.plot[i][j].added = true
				region := &Region{[]*SinglePlot{g.plot[i][j]}, 1, g.perimeter(i,j), 4}
				g.plot[i][j].region = region
				g.discover(i, j, region)
				fmt.Printf("\t{%s}: Area: %d, Perimeter: %d, Sides: %d\n", string(g.plot[i][j].plant), region.area, region.perimeter, region.sides)
				//g.print(region)
				price += region.area*region.perimeter
				discountedPrice += region.area*region.sides
				//fmt.Println()
				//bufio.NewReader(os.Stdin).ReadBytes('\n') 
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
	// discover entire region
	for _, q := range [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}} {
		if i+q[0] >= 0 && 
		i+q[0] < len(g.plot) && 
		j+q[1] >= 0 && 
		j+q[1] < len(g.plot[i+q[0]]) && // (i+r, j+s) is within bounds, and 
		g.plot[i+q[0]][j+q[1]].plant == g.plot[i][j].plant && // the plant type is the same, and
		!g.plot[i+q[0]][j+q[1]].added { // the plant doesn't belong to a region yet
			//fmt.Printf("\tAdding {%s}:(%d,%d)\n", string(g.plot[i+q[0]][j+q[1]].plant), i+q[0], j+q[1])
		  g.plot[i+q[0]][j+q[1]].added = true
			g.plot[i+q[0]][j+q[1]].region = reg
			reg.plants = append(reg.plants, g.plot[i+q[0]][j+q[1]])
			reg.area++
			reg.perimeter += g.perimeter(i+q[0], j+q[1])
			g.sides(i+q[0], j+q[1], reg)

			//g.print(reg)
			//fmt.Printf("\tArea: %d, Perimeter: %d, Sides: %d\n", reg.area, reg.perimeter, reg.sides)
			//bufio.NewReader(os.Stdin).ReadBytes('\n') 

			g.discover(i+q[0], j+q[1], reg)
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

func (g *Garden) adjacency(i, j int) int {
	adjacency := 0
	for _, q := range [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}} {
		if i+q[0] >= 0 && 
		i+q[0] < len(g.plot) && 
		j+q[1] >= 0 && 
		j+q[1] < len(g.plot[i+q[0]]) && // (i+r, j+s) is within bounds, and 
		g.plot[i+q[0]][j+q[1]].plant == g.plot[i][j].plant && // the plant type is the same
		g.plot[i+q[0]][j+q[1]].region != nil && g.plot[i+q[0]][j+q[1]].region == g.plot[i][j].region { // the adjacent plant already belongs to the plot
			adjacency++
		}
	}

	return adjacency
}

func (g *Garden) sides(i, j int, r *Region) int {
	isolated := false
	cornerAdjacency := 0
	adjacency := g.adjacency(i,j)
	xSum := 0
	ySum := 0
	corners := 0
	for _, q := range [][]int{{-1, -1}, {1, -1}, {1, 1}, {-1, 1}} {
		if i+q[0] >= 0 && 
		i+q[0] < len(g.plot) && 
		j+q[1] >= 0 && 
		j+q[1] < len(g.plot[i+q[0]]) && // (i+r, j+s) is within bounds, and 
		g.plot[i+q[0]][j+q[1]].plant == g.plot[i][j].plant && // the plant type is the same
		g.plot[i+q[0]][j+q[1]].region != nil && g.plot[i+q[0]][j+q[1]].region == g.plot[i][j].region { // the adjacent plant already belongs to the plot
			corners++

			// corner adjacency 
			cornerAdjacency = g.adjacency(i+q[0], j+q[1])

			if i+q[0]-1 >= 0 && j+q[1]-1 >= 0 && 
			g.plot[i+q[0]-1][j+q[+1]].region != g.plot[i][j].region &&
			g.plot[i+q[0]][j+q[1]-1].region != g.plot[i][j].region {
				isolated = true
			}
			xSum += q[0]
			ySum += q[1]
		}
	}

	sides := -100
	if adjacency == 0 {
		sides = 0
	} else if adjacency  == 1 {
		if corners == 0 {
			sides = 0
		} else if corners == 1 {
			sides = 2
		} else if corners == 2 && !(xSum == 0 && ySum == 0) {
			sides = 4
		} else if corners == 2 {
			sides = 2
		} else if corners == 3 || corners == 4 {
			sides = 4
		} 
	} else if adjacency == 2 {
		if corners == 1 {
			sides = -2
		} else if corners == 2 {
			sides = 0
		} else if corners == 3 && !isolated {
			// Todo this is not correct as 2, at least not in cases where 1 corner touches none of the adjacents to this one
			sides = 2
		} else if corners == 3 && isolated {
			sides = 0
		} else if corners == 4 {
			sides = 4
		}
	} else if adjacency == 3 {
		if corners == 1 && cornerAdjacency == 1 { 
			sides = -2
		} else if corners == 1 && cornerAdjacency == 2 {
			sides = -4
		} else if corners == 2 {
			sides = -4
		} else if corners == 3 {
			sides = -2
		} else if corners == 4 {
			sides = 0
		}
	} else if adjacency == 4 {
		sides = -4
	}
	if sides == -100 {
		//fmt.Printf("Got adjacency %d but corners %d\n", adjacency, corners)
		r.sides = 0
		//panic("should never reach here")
	}
	r.sides += sides
	//fmt.Printf("\tAdding {%s}:(%d,%d) with adjacency %d, corners %d = sides %d; TOTAL sides = %d\n", string(g.plot[i][j].plant), i, j, adjacency, corners, sides, r.sides)
	// ---------------------> CASE 1: i,j has adjacency 1
	// and no corners -> sides stay the same
	// A A A {A} 
	//
	// and one corner -> sides increase by 2
	// A A 
	//  {A}
	// 
	// and two corners, top or bottom -> sides increase by 4 
	// A  A  A
	//   {A} 
	// and two corners, top or bottom -> sides increase by 2
	// A  A
	//   {A}
	//       A
	// and three corners -> sides increase by 
	// A  A  A
	//   {A} 
	// A
	// ---------------------> CASE 2: i,j has adjacency 2
	// and one corner -> sides decrease by 2
	// A  A
	// A {A}
	// and two corners -> sides stay the same
	//  A  A    A  A    A A      A A A
	//  A {A}     {A}    {A}      {A}
	//  A       A  A      A A      A
	// and three corners, none isolated -> sides increase by 2 
	// A A     
	//  {A}     
	// A A A     
	// or three corners, one isolated -> sides stay the same...
	// A  A  A
	// A {A}    
	//       A     
	// and four corners -> sides increase by 4 (from 8 to 12)
	// A A A
	//  {A}
	// A A A
	// ---------------------> CASE 3: i,j has adjacency 3
	// and one corner, which has adjacency 2 -> sides decrease by 4
	// A  A
	// A {A}
	//    A  
	// and one corner, which has adjacency 1 -> sides decrease by 2
	//    A  A
	// A {A}
	//    A  
	// and two corners -> sides decrease by 4
	// A  A      
	// A {A}    
	// A  A    
	// and three corners -> sides decrease by 2 
	// A  A  A       A  A
	// A {A}      A {A}
	// A  A          A
	// and four corners -> sides stays the same
	// A  A  A
	// A {A}
	// A  A  A
	// ---------------------> CASE 3: i,j has adjacency 3
	//                  -> sides reduces by 4
	// A  A  A
	// A {A} A
	// A  A  A

	

	return sides
}