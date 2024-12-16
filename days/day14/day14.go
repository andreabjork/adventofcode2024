package day14

import (
	"adventofcode/m/v2/util"
	"fmt"
	"strings"
	"strconv"
	"time"
)

func Day14(inputFile string, part int) {
	if part == 0 {
		fmt.Printf("Safety factor: %d\n", solve(inputFile, 103, 101, false))
	} else {
		fmt.Printf("Safety factor: %d\n", solve(inputFile, 103, 101, true))
	}
}

func solve(inputFile string, height, width int, xmasTree bool) int {
	ls := util.LineScanner(inputFile)
	line, ok := util.Read(ls)

	g := NewGrid(height, width)
	robots := []*Robot{}
	var parts, pos, vel []string
	var x, y, vX, vY int
	for ok {
		parts = strings.Split(line, " ")
		pos = strings.Split(strings.TrimPrefix(parts[0], "p="), ",")
		vel = strings.Split(strings.TrimPrefix(parts[1], "v="), ",")
		x, _ = strconv.Atoi(pos[0])
		y, _ = strconv.Atoi(pos[1])
		vX, _ = strconv.Atoi(vel[0])
		vY, _ = strconv.Atoi(vel[1])
	
		robots = append(robots, &Robot{x, y, vX, vY})

		line, ok = util.Read(ls)
	}

	if !xmasTree {
		t := 100
		for _, robot := range robots {
			g.add(robot.atTime(t, height, width))	
		}
		return g.safetyFactor()
	} else {
		allGrids := []*Grid{}
		t := 11000 // max 103*101 iterations
		for i := 0; i < t; i++ {
			gg := NewGrid(height, width)
			for _, robot := range robots {
				gg.add(robot.atTime(i, height, width))
			}

			allGrids = append(allGrids, gg)
		}

		minSafety := allGrids[0].safetyFactor() 
		for i, gg := range allGrids {
			if s := gg.safetyFactor(); s < minSafety {
				minSafety = s
				fmt.Print("\033[H\033[2J")
				gg.print()                      		
				fmt.Println(i)                  		
				time.Sleep(400*time.Millisecond)		
			}
		}
	}
	

	// Found by watching the output above print out small-safety factor configurations
	return 7132
}

type Robot struct {
	x, y int
	vX, vY int
}

func (r *Robot) atTime(t, height, width int) *Robot {
	rx := (r.x+t*r.vX) % width  
	ry := (r.y+t*r.vY) % height

	if rx < 0 {
		rx = (rx + (util.Abs(rx)%width)*width)%width
	} 
	if ry < 0 {
		ry = (ry + (util.Abs(ry)%height)*height)%height
	}
	return &Robot{rx, ry, r.vX, r.vY}
}

type Grid struct {
	g map[int]map[int][]*Robot
}

func NewGrid(height, width int) *Grid {
	g := make(map[int]map[int][]*Robot, height)
	for i := 0; i < height; i++ {
		g[i] = make(map[int][]*Robot, width)
		for j := 0; j < width; j++ {
			g[i][j] = []*Robot{}
		}
	}

	return &Grid{g}
}

func (g *Grid) add(r *Robot) {
	//fmt.Printf("Adding at y,x = %d, %d\n", r.y, r.x)
	//fmt.Printf("-90 mod 103", -90%103)
	//fmt.Printf("len of y: %d\n", len(g.g))
	//fmt.Printf("len of x: %d\n", len(g.g[0]))
	g.g[r.y][r.x] = append(g.g[r.y][r.x], r)
}

func (g *Grid) safetyFactor() int {
	i_center := len(g.g)/2
	j_center := len(g.g[0])/2

	var tl, bl, br, tr int = 0, 0, 0, 0
	// Top left quadrant
	for i := 0; i < i_center; i++ {
		for j := 0; j < j_center; j++ {
			tl += len(g.g[i][j])
		}
	} 


	// Bottom left quadrant 
	for i := i_center+1; i < len(g.g); i++ {
		for j := 0; j < j_center; j++ {
			bl += len(g.g[i][j])
		}
	} 

	// Bottom right quadrant
	for i := i_center+1; i < len(g.g); i++ {
		for j := j_center+1; j < len(g.g[0]); j++ {
			br += len(g.g[i][j])
		}
	}

	// Top right quadrant
	for i := 0; i < i_center; i++ {
		for j := j_center+1; j < len(g.g[0]); j++ {
			tr += len(g.g[i][j])
		}
	}

	return tl*bl*br*tr
}

func (g *Grid) print() {
	i_center := len(g.g)/2
	j_center := len(g.g[0])/2
	//fmt.Printf("Height is %d, middle is %d\n", len(g.g), i_center)
	//fmt.Printf("Width is %d, middle is %d\n", len(g.g[0]), j_center)
	for i := 0; i < len(g.g)/2; i++ {
		if i == i_center {
			fmt.Println()
			continue
		}
		for j := 0; j < len(g.g[i]); j++ {
			if j == j_center {
				fmt.Printf(" ")
			} else {
				if len(g.g[i][j])==0 {
					fmt.Printf(".")
				} else {
					fmt.Printf("#")
					//fmt.Printf("%d", len(g.g[i][j]))
				}
			}
		}
		fmt.Printf("\n")
	}
	fmt.Println()
}