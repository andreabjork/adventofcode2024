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
				g = NewGuard(i, j)
			}
			j++
		}
		i++

		line, ok = util.Read(ls)
	}

	if checkPotentials {
		g.walkAndDetect(m)
		return m.potentials
	} else {
		g.walk(m)
		return g.visited
	}
}

// ====
// PATH
// ==== 
type Direction int
const (
	UP Direction = iota
  DOWN 
	LEFT
	RIGHT
)

// ====
// GRID
// ====
type Spot int
const (
	OBSTRUCTION Spot = iota
  FREE 
	GUARD
)

type Grid struct {
	m map[int]map[int]Spot 
	potential	map[int]map[int]bool
	potentials int
}

func NewGrid() *Grid {
	return &Grid{
		make(map[int]map[int]Spot),
		make(map[int]map[int]bool),
		0, 
	}
}

func (g *Grid) set(i int, j int, s Spot) {
	if g.m[i] == nil {
		g.m[i] = make(map[int]Spot)
	}
	g.m[i][j] = s
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
	path map[int]map[int]map[Direction]bool
	x, y int 
	dirX, dirY int // direction vector w length 1
	visited int
}

func NewGuard(x, y int) *Guard {
	g := &Guard{
		make(map[int]map[int]map[Direction]bool),
		x, y,
		-1, 0,
		0,
	}
	g.add(x, y)
	return g
}

func (g *Guard) turn() {
	x := g.dirX
	g.dirX = g.dirY
	g.dirY = -x
}

func (g *Guard) dir() Direction {
	if g.dirX == 1 && g.dirY == 0 {
		return DOWN
	} else if g.dirX == 0 && g.dirY == 1 {
		return RIGHT
	} else if g.dirX == 0 && g.dirY == -1 {
		return LEFT 
	} else {
		return UP
	}
}

func (g *Guard) crossed(i, j int) bool {
  return g.path[i][j][UP] || g.path[i][j][DOWN] || g.path[i][j][LEFT] || g.path[i][j][RIGHT]
}

func (g *Guard) add(i,j int) bool {
	if g.path[i] == nil {
		g.path[i] = make(map[int]map[Direction]bool)
	}

	if g.path[i][j] == nil {
		g.path[i][j] = make(map[Direction]bool)
	}

	if g.path[i][j][g.dir()] {
		// Path already exists, we're in a cycle
		return true
	} else {
		g.x = i
		g.y = j
		// if this segment has ever been visited before we don't count it 
		if !g.crossed(i,j) {
  		g.visited++
		}
		g.path[i][j][g.dir()] = true
		return false
	}
}

func (g *Guard) copy() *Guard {
	path := make(map[int]map[int]map[Direction]bool)
	for i := range g.path {
		path[i] = make(map[int]map[Direction]bool)
		for j := range g.path[i] {
			path[i][j] = make(map[Direction]bool)
			for k := range g.path[i][j] {
				path[i][j][k] = g.path[i][j][k]  
			}
		}
	}
	return &Guard{
		path,
		g.x, g.y,
		g.dirX, g.dirY,
		g.visited,
	}
}

// Walks until off the grid or a cycle is detected
func (g *Guard) walk(m *Grid) bool {
	// i, j, direction DOWN 1, 0 -> i+1, 0
	// i, j, direction RIGHT 0, 1 -> i, j+1
	// i, j, direction LEFT 0, -1 -> i, j-1
	// i, j, direction UP -1, 0 -> i-1, j
	nextX, nextY := g.x+g.dirX, g.y+g.dirY

	for m.onGrid(nextX, nextY) {
		if m.obstructed(nextX, nextY) {
			g.turn()
		} else if g.add(nextX, nextY) {
			// we hit a cycle and must break
			return true
		}

		nextX, nextY = g.x+g.dirX, g.y+g.dirY
	}
	
	return false
}

func (g *Guard) walkAndDetect(m *Grid) {
	// i, j, direction DOWN 1, 0 -> i+1, 0
	// i, j, direction RIGHT 0, 1 -> i, j+1
	// i, j, direction LEFT 0, -1 -> i, j-1
	// i, j, direction UP -1, 0 -> i-1, j
	nextX, nextY := g.x+g.dirX, g.y+g.dirY

	for m.onGrid(nextX, nextY) {
		if m.obstructed(nextX, nextY) {
			// Mark this hit
			g.turn()
			nextX, nextY = g.x+g.dirX, g.y+g.dirY
		} else  {
			// Otherwise, imagine there's an obstruction directly in front
			// of us at nextX, nextY and we have to turn. Check if that path cycles.
			if !g.crossed(nextX, nextY) {
				h := g.copy()
				h.turn()
				m.set(nextX, nextY, OBSTRUCTION)
				if h.walk(m) {
					m.addPotential(nextX, nextY)
				}
				m.set(nextX, nextY, FREE)
				
			}

			// then keep walking
			g.add(nextX, nextY)
			nextX, nextY = g.x+g.dirX, g.y+g.dirY
		}
	}
}
