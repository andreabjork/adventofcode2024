package day15

import (
	"adventofcode/m/v2/util"
	"bufio"
	"fmt"
)

func Day15(inputFile string, part int) {
	if part == 0 {
		fmt.Printf("GPS: %d\n", solve(inputFile, false))
	} else {
		fmt.Printf("GPS: %d\n", solve(inputFile, true))
	}
}

func solve(inputFile string, double bool) int {
	ls := util.LineScanner(inputFile)
	var w *Warehouse
	var robot *Robot
	var boxes []IUnit
	if !double {
		w, robot, boxes = singleWarehouse(ls)
	} else {
		w, robot, boxes = doubleWarehouse(ls)
	}

	line, ok := util.Read(ls)
	for ok {
		for _, dir := range line {
			switch dir {
			case '^':
				w.moveIfAble(robot, -1, 0, false)
			case '<':
				w.moveIfAble(robot, 0, -1, false)
			case '>':
				w.moveIfAble(robot, 0, 1, false)
			case 'v':
				w.moveIfAble(robot, 1, 0, false)
			}
		}

		line, ok = util.Read(ls)
	}
	
	sum := 0
	for _, b := range boxes {
		sum += b.gps()
	}
	return sum
}

func singleWarehouse(ls *bufio.Scanner) (*Warehouse, *Robot, []IUnit) {
	line, ok := util.Read(ls)
	w := &Warehouse{make(map[int]map[int]IUnit)}
	var robot *Robot
	boxes := []IUnit{}
	i, j := 0, 0
	for ok {
		j = 0
		for _, r := range line {
			if w.grid[i] == nil {
				w.grid[i] = make(map[int]IUnit)
			}

			switch r {
			case '#':
				w.grid[i][j] = &Wall{&Unit{i,j}}
			case 'O':
				w.grid[i][j] = &Box{&Unit{i,j}}
				boxes = append(boxes, w.grid[i][j])
			case '@':
				robot = &Robot{&Unit{i,j}}
				w.grid[i][j] = robot
			case '.':
				w.grid[i][j] = &Empty{&Unit{i,j}}
			default:
				panic("unexpected input")
			}
			j++
		}

		i++
		line, ok = util.Read(ls)
		if line == "" {
			break
		}
	}

	return w, robot, boxes
}

func doubleWarehouse(ls *bufio.Scanner) (*Warehouse, *Robot, []IUnit) {
	line, ok := util.Read(ls)
	w := &Warehouse{make(map[int]map[int]IUnit)}
	var robot *Robot
	boxes := []IUnit{}
	i, j := 0, 0
	for ok {
		j = 0
		for _, r := range line {
			if w.grid[i] == nil {
				w.grid[i] = make(map[int]IUnit)
			}

			switch r {
			case '#':
				w.grid[i][j] = &Wall{&Unit{i,j}}
				w.grid[i][j+1] = &Wall{&Unit{i,j}}
			case 'O':
				w.grid[i][j] = &LBox{&Unit{i,j}}
				w.grid[i][j+1] = &RBox{&Unit{i,j+1}}
				boxes = append(boxes, w.grid[i][j])
			case '@':
				robot = &Robot{&Unit{i,j}}
				w.grid[i][j] = robot
				w.grid[i][j+1] = &Empty{&Unit{i,j}}
			case '.':
				w.grid[i][j] = &Empty{&Unit{i,j}}
				w.grid[i][j+1] = &Empty{&Unit{i,j}}
			default:
				panic("unexpected input")
			}
			j+=2
		}

		i++
		line, ok = util.Read(ls)
		if line == "" {
			break
		}
	}

	return w, robot, boxes
}
func (w *Warehouse) print() {
	for i := 0; i < len(w.grid); i++ {
		for j := 0; j < len(w.grid[i]); j++ {
			if w.grid[i][j].symbol() == '@' {
				var Boldred = "\033[31;1m"
				var Reset = "\033[0m"  
				fmt.Printf(Boldred+"@"+Reset)	
			} else if w.grid[i][j].symbol() == '[' || w.grid[i][j].symbol() == ']' {
				var BoldBlue = "\033[34;1m"
				var Reset = "\033[0m"  
				fmt.Printf(BoldBlue+string(w.grid[i][j].symbol())+Reset)	
			} else {
				fmt.Printf(string(w.grid[i][j].symbol()))
			}
		}
		fmt.Printf("\n")
	}
	fmt.Println()
}

func (w *Warehouse) inWarehouse(i, j int) bool {
	return i >= 0 && j >= 0 && i < len(w.grid) && j < len(w.grid[i]) 
}

func (w *Warehouse) isBlocked(u IUnit, dX, dY int, pushed bool) bool {
	if w.inWarehouse(u.x()+dX, u.y()+dY) {
		next := w.grid[u.x()+dX][u.y()+dY]
		if _, ok := next.(*LBox); ok && pushed && dX != 0 {
			right := w.grid[u.x()+dX][u.y()+dY+1]
			return w.isBlocked(next, dX, dY, true) || w.isBlocked(right, dX, dY, true)
		} else if _, ok := next.(*RBox); ok && pushed && dX != 0 {
			left := w.grid[u.x()+dX][u.y()+dY-1]
			return w.isBlocked(next, dX, dY, true) || w.isBlocked(left, dX, dY, true)
		} else if _, ok := next.(*Empty); ok {
			return false
		} else if _, ok := next.(*Wall); ok {
			return true
		} else { 
			return w.isBlocked(next, dX, dY, true)
		}
	} 

	return true
}


func (w *Warehouse) moveIfAble(u IUnit, dX, dY int, pushed bool) {
	if w.isBlocked(u, dX, dY, true) {
		return
	} else {
		w.move(u, dX, dY, pushed)
	}

}

// Precondition: the unit u is not blocked
func (w *Warehouse) move(u IUnit, dX, dY int, pushed bool) {
	if !pushed {
		w.grid[u.x()][u.y()] = &Empty{&Unit{u.x(), u.y()}}
	} 
	var left, right IUnit
	next := w.grid[u.x()+dX][u.y() + dY]
	_, isLeft := u.(*LBox)
	_, isRight := u.(*RBox)
	if dX != 0 && isLeft && pushed {
		right = w.grid[u.x()][u.y()+1]
	} else if dX != 0 && isRight && pushed {
		left = w.grid[u.x()][u.y()-1]
	}
	// move this unit
	u.setX(u.x()+dX)
	u.setY(u.y()+dY)
	w.grid[u.x()][u.y()] = u
	
	if _, ok := next.(*Empty); !ok {
		w.move(next, dX, dY, true)
	}
	if left != nil {
		w.move(left, dX, dY, false)
	} else if right != nil {
		w.move(right, dX, dY, false)
	}
}

func (b Unit) gps() int {
	return 100*b.i + b.j
}

type Warehouse struct {
	grid map[int]map[int]IUnit
}

/**
* This solution plays with embedded structs (just to understand how they work)
* so it's overly complex and would be more neatly done with enums.
**/ 
type IUnit interface {
	x() int 
	y() int
	setX(int)
	setY(int)
	symbol() rune
	gps() int
}

type Unit struct {
	i, j int
}

func (u *Unit) x() int {
	return u.i
}

func (u *Unit) y() int {
	return u.j
}

func (u *Unit) setX(x int) {
	u.i = x  
}

func (u *Unit) setY(y int) {
	u.j = y
}

type Wall struct {
	*Unit
}

func (w Wall) symbol() rune {
	return '#'
}
func (e Empty) symbol() rune {
	return '.'
}
func (b Box) symbol() rune {
	return 'O'
}

func (r Robot) symbol() rune {
	return '@'
}

func (lb LBox) symbol() rune {
	return '['
}

func (rb RBox) symbol() rune {
	return ']'
}

type Empty struct {
	*Unit
}

type Robot struct {
	*Unit
}

type Box struct {
	*Unit
}

type LBox struct {
	*Unit
}

type RBox struct {
	*Unit
}
