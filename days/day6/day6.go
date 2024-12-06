package day6

import (
	"adventofcode/m/v2/util"
	"fmt"
)

func Day6(inputFile string, part int) {
	if part == 0 {
		fmt.Printf("Visited: %d\n", guardWalk(inputFile, false))
	} else {
		fmt.Printf("Potential Obstructions: %d\n", guardWalk(inputFile, true))
	}
}

func guardWalk(inputFile string, checkPotentials bool) int {
	ls := util.LineScanner(inputFile)
	line, ok := util.Read(ls)
	
	i, j := 0, 0
	var m *Grid
	var g *Guard
	m = NewGrid()
	for ok {
		j = 0
		for _,  c := range(line) {
			if c == '.' {
				m.set(i, j, FREE)
			} else if c == '#' {
				m.set(i, j, OBSTRUCTION)
			} else if c == '^' {
				m.set(i, j, FREE)
				m.visit(i, j)
				g = &Guard{i, j, -1, 0}
			}
			j++
		}
		i++

		line, ok = util.Read(ls)
	}

	if checkPotentials {
		g.walkAndDetect(m)
		fmt.Printf("AFTER %+v\n", m.potential)
		for i := range m.potential {
			for j, b := range m.potential[i] {
				if b {
					fmt.Printf("Potential: %d, %d\n", i,j)
				}
			}
		}
		//fmt.Printf("OBS: %+v\n", m.hits)
		return m.potentials
	} else {
		g.walk(m)
		return m.visited
	}
}

// ====
// GRID
// ====
type Spot int
const (
	OBSTRUCTION Spot = iota
  FREE 
	GUARD
)

type Direction int
const (
	UP Direction = iota
  DOWN 
	LEFT
	RIGHT
)

type Grid struct {
	m map[int]map[int]Spot 
	v map[int]map[int]bool
	hits map[int]map[int]map[Direction]bool // Marks when an obstruction is hit from a given direction
	potential	map[int]map[int]bool
	potentials int
	visited int
}

func NewGrid() *Grid {
	return &Grid{
		make(map[int]map[int]Spot),
		make(map[int]map[int]bool),
		make(map[int]map[int]map[Direction]bool),
		make(map[int]map[int]bool),
		0,
		0,
	}
}

func (g *Grid) set(i int, j int, s Spot) {
	if g.m[i] == nil {
		g.m[i] = make(map[int]Spot)
	}
	g.m[i][j] = s
}

func (g *Grid) visit(i int, j int) {
	if g.v[i] == nil {
		g.v[i] = make(map[int]bool)
	}
	if !g.v[i][j] {
		g.v[i][j] = true 
		g.visited++
	}
}

func (g *Grid) addPotential(i int, j int) {
	if g.potential[i] == nil {
		g.potential[i] = make(map[int]bool)
	}
	if !g.potential[i][j] {
		g.potential[i][j] = true 
		g.potentials++
	}
}

func (g *Grid) hit(i int, j int, dir Direction) {
	if g.hits[i] == nil {
		g.hits[i] = make(map[int]map[Direction]bool)
	}
	if g.hits[i][j] == nil {
		g.hits[i][j] = make(map[Direction]bool)
	}
	g.hits[i][j][dir] = true
}

func (g *Grid) print() {
	for i := 0; i < len(g.m); i++ {
		for j := 0; j < len(g.m[i]); j++ {
			if g.v[i][j] {
				fmt.Printf("X")
			} else if g.m[i][j] == OBSTRUCTION {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
	fmt.Println("---")
}

func (g *Grid) onGrid(i int, j int) bool {
	return i >= 0 && i < len(g.m) && j >= 0 && j < len(g.m[0])
}

func (g *Grid) obstructed(i int, j int) bool {
	return g.m[i][j] == OBSTRUCTION
}

// =====
// GUARD
// =====
type Guard struct {
	x, y int 
	dirX, dirY int // direction vector w length 1
}

func (g *Guard) turn() {
	x := g.dirX
	g.dirX = g.dirY
	g.dirY = -x
}

func (g *Guard) turnBack() {
	x := g.dirX
	g.dirX = -g.dirY
	g.dirY = x
}

func (g *Guard) dir() Direction {
	if g.x == 1 && g.y == 0 {
		return DOWN
	} else if g.x == 0 && g.y == 1 {
		return RIGHT
	} else if g.x == 0 && g.y == -1 {
		return LEFT 
	} else {
		return UP
	}
}

func (g *Guard) walk(m *Grid) {
	// i, j, direction DOWN 1, 0 -> i+1, 0
	// i, j, direction RIGHT 0, 1 -> i, j+1
	// i, j, direction LEFT 0, -1 -> i, j-1
	// i, j, direction UP -1, 0 -> i-1, j
	nextX, nextY := g.x+g.dirX, g.y+g.dirY

	for m.onGrid(nextX, nextY) {
		if m.obstructed(nextX, nextY) {
			//m.print()
			g.turn()
		} else {
			g.x = nextX
			g.y = nextY
		}

		m.visit(g.x, g.y)
		nextX, nextY = g.x+g.dirX, g.y+g.dirY

	}
}

// Hitting an obstruction 2x while facing the same way implies a cycle.
//
// While walking, we place an imaginary obstruction in front of us at each step.
// If this obstruction would lead us to hit an obstruction we've already hit before
// while facing the same way, we mark it as a potential cyclical obstruction before continuing.
func (g *Guard) walkAndDetect(m *Grid) {
	// i, j, direction DOWN 1, 0 -> i+1, 0
	// i, j, direction RIGHT 0, 1 -> i, j+1
	// i, j, direction LEFT 0, -1 -> i, j-1
	// i, j, direction UP -1, 0 -> i-1, j
	nextX, nextY := g.x+g.dirX, g.y+g.dirY

	for m.onGrid(nextX, nextY) {
		if m.obstructed(nextX, nextY) {
			// Mark this hit
			m.hit(nextX, nextY, g.dir())
			g.turn()
		} else {
			// Otherwise, imagine there's an obstruction directly in front
			// of us at nextX, nextY and we have to turn. Check if we hit a previously hit
			// obstruction this way
			fmt.Printf("At %d, %d: imagining block at %d, %d\n", g.x, g.y, nextX, nextY)
			g.turn()
			stopped, x, y := g.nextStop(m)
			if stopped && m.hits[x][y][g.dir()] {
				m.addPotential(nextX, nextY)
				fmt.Printf("Cycle! The guard would stop at %d, %d again\n\n", x, y)
			} else if stopped {
				fmt.Printf("The guard would stop at %d, %d, no cycle.\n", x, y)
			} else {
				fmt.Printf("If we put the block there, the guard would go off the grid\n")
			}

			// Otherwise, put g back to its original spot with its original rotation
			g.turnBack()
			g.x = nextX
			g.y = nextY
		}

		m.visit(g.x, g.y)
		nextX, nextY = g.x+g.dirX, g.y+g.dirY
	}
}

func (g *Guard) nextStop(m *Grid) (bool, int, int) {
	nextX, nextY := g.x+g.dirX, g.y+g.dirY

	for m.onGrid(nextX, nextY) {
		if m.obstructed(nextX, nextY) {
			return true, nextX, nextY	
		} else {
			g.x = nextX
			g.y = nextY
		}

		nextX, nextY = g.x+g.dirX, g.y+g.dirY
	}

	// Went off the grid
	return false, -1, -1
}