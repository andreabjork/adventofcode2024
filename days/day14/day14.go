package day14

import (
	"adventofcode/m/v2/util"
	"fmt"
	"strings"
	"strconv"
	"time"
	"os"
	"os/exec"
)

func Day14(inputFile string, part int) {
	if part == 0 {
		fmt.Printf("Safety factor: %d\n", solve(inputFile, 103, 101))
	} else {
		fmt.Printf("Safety factor: %d\n", solve(inputFile, 103, 101))
	}
}

func solve(inputFile string, height, width int) int {
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
	
	//testRobot := &Robot{2,4,2,-3}
	//fmt.Printf("%+v\n", testRobot.atTime(5, height, width))
	//if true {
	//	return 0
	//}
	t := 100
	for _, robot := range robots {
		g.add(robot.atTime(t, height, width))	
	}

	g.print()

	t = 10000
	for i := 0; i < t; i++ {
		g = NewGrid(height, width)
		for _, robot := range robots {
			g.add(robot.atTime(t, height, width))
		}
		g.print()
    cmd := exec.Command(`printf '\33c\e[3J'`) // clears the scrollback buffer as well as the screen.
    cmd.Stdout = os.Stdout
    cmd.Run()
		time.Sleep(300*time.Microsecond)
	}


	return g.safetyFactor()
}

type Robot struct {
	x, y int
	vX, vY int
}

func (r *Robot) atTime(t, height, width int) *Robot {
	r.x = (r.x+t*r.vX) % width  
	r.y = (r.y+t*r.vY) % height

	if r.x < 0 {
		r.x = (r.x + (util.Abs(r.x)%width)*width)%width
	} 
	if r.y < 0 {
		r.y = (r.y + (util.Abs(r.y)%height)*height)%height
	}
	return r
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
	for i := 0; i < len(g.g); i++ {
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