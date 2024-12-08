package day8

import (
	"adventofcode/m/v2/util"
	"fmt"
)

func Day8(inputFile string, part int) {
	if part == 0 {
		fmt.Printf("Antenodes: %d\n", solve(inputFile, false))
	} else {
		fmt.Printf("Antenodes: %d\n", solve(inputFile, true))
	}
}

func solve(inputFile string, repeat bool) int {
	ls := util.LineScanner(inputFile)
	line, ok := util.Read(ls)

	grid := &Grid{
		make(map[int]map[int]rune),
		make(map[int]map[int]rune),
		0,
	}
	antennas := make(map[rune][]*Point)
	i, j := 0,0
	for ok {
		j = 0
		for _, c := range line {
			if grid.g[i] == nil {
				grid.g[i] = make(map[int]rune)
			}
			grid.g[i][j] = c
			if c != '.' {
				antennas[c] = append(antennas[c], &Point{i,j})
			}
			j++
		}
		i++
		line, ok = util.Read(ls)
	}

	grid.placeAntenodes(antennas, repeat)
	return grid.antenodeCount
}

type Point struct {
	x, y int
}

type Grid struct {
	g map[int]map[int]rune
	antenodes map[int]map[int]rune
	antenodeCount int
}

func (g *Grid) placeAntenodes(antennas map[rune][]*Point, repeat bool) {
	for _, signal := range(antennas) {
		// Connect a line through every 2 antennas and add the antenodes:
		for i := 0; i < len(signal); i++ {
			for j := i+1; j < len(signal); j++ {
				g.placeAlongLine(signal[i], signal[j], repeat)	
			}
		}
	}
}

// Places antinodes along the line {from; to}. If repeat, it
// repeats the antenna placements further along.
func (g *Grid) placeAlongLine(from, to *Point, repeat bool) {
	dx := to.x-from.x
	dy := to.y-from.y

	if repeat {
		g.addAntenode(from)
	}
	fromX, fromY := from.x, from.y
	for fromX-dx >= 0 && fromX-dx < len(g.g) && fromY-dy >= 0 && fromY-dy < len(g.g[fromX]) {
		g.addAntenode(&Point{fromX-dx, fromY-dy})

		if !repeat {
			break
		}
		fromX = fromX-dx
		fromY = fromY-dy
	}

	if repeat {
		g.addAntenode(to)
	}
	toX, toY := to.x, to.y
	for toX+dx >= 0 && toX+dx < len(g.g) && toY+dy >= 0 && toY+dy < len(g.g[toX]) {
		g.addAntenode(&Point{toX+dx, toY+dy})
		if !repeat {
			break
		}
		toX = toX+dx
		toY = toY+dy
	}
}

func (g *Grid) addAntenode(p *Point) {
	if g.antenodes[p.x] == nil {
		g.antenodes[p.x] = make(map[int]rune)
	}
	// We add an antenode only if it wasn't added before 
	if g.antenodes[p.x][p.y] != '#' {
		g.antenodes[p.x][p.y] = '#'
		g.antenodeCount++
	}
}