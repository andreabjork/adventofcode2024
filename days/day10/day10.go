package day10

import (
	"fmt"
	"adventofcode/m/v2/util"
)

func Day10(inputFile string, part int) {
	if part == 0 {
		score, _ := solve(inputFile)
		fmt.Printf("Trailhead scores: %d\n", score)
	} else {
		_, rating := solve(inputFile)
		fmt.Printf("Trailhead ratings: %d\n", rating)
	}
}

func solve(inputFile string) (int, int) {
	ls := util.LineScanner(inputFile)
	line, ok := util.Read(ls)

	t := &Terrain{make(map[int]map[int]*Point)}
	trailHeads := []*Point{}
	i := 0
	for ok {
		for j, r := range line {
			p := P(i, j, int(r)-48)
			t.add(p)
			if int(r)-48 == 0 {
				trailHeads = append(trailHeads, p)
			}
		}	
		
		i++
		line, ok = util.Read(ls)
	}

	
	sumScores := 0
	sumRatings := 0
	for _, th := range trailHeads {
		reachable := make(map[*Point]int)
		score := 0 // +1 for each unique peak we can reach
		rating := 0 // +1 for each unique path to a peak
		var peak *Point
		for _, path := range th.getPaths(t) {
			peak = path[len(path)-1]
			reachable[peak]++
			if reachable[peak] == 1 {
				score++
			}
			rating++
		}
		
		sumScores += score
		sumRatings += rating
	}

	return sumScores, sumRatings
}

type Terrain struct {
	g map[int]map[int]*Point
}

func (t *Terrain) add(p *Point) {
	if t.g == nil {
		t.g = make(map[int]map[int]*Point)
	}
	if t.g[p.i] == nil {
		t.g[p.i] = make(map[int]*Point)
	}
	t.g[p.i][p.j] = p
}

type Point struct {
	i, j  int
	val   int
	paths [][]*Point // Every point contains all paths to 9
}

func P(i, j, val int) *Point {
	return &Point{
		i, j, val,
		nil,
	}
}

func (p *Point) addTo(i, j int, t *Terrain) [][]*Point {
	all := [][]*Point{}
	if i >= 0 && i < len(t.g) && j >= 0 && j < len(t.g[i]) && 
	t.g[i][j].val == p.val+1 {
		for _, q := range t.g[i][j].getPaths(t) {
			pp  := append([]*Point{p}, q...)
			all = append(all, pp)
		}
	}

	return all
}

// Returns all unique paths to all peaks in the terrain
func (p *Point) getPaths(t *Terrain) [][]*Point {
	if p.val == 9 {
		p.paths = [][]*Point{{p}}
		return p.paths
	}

	if p.paths == nil {
		all := [][]*Point{}
		all = append(all, p.addTo(p.i-1, p.j, t)...)
		all = append(all, p.addTo(p.i+1, p.j, t)...)
		all = append(all, p.addTo(p.i, p.j-1, t)...)
		all = append(all, p.addTo(p.i, p.j+1, t)...)
		return all
	} else {
		return p.paths
	}
}
