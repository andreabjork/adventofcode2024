package day15

import (
	"adventofcode/m/v2/util"
	"fmt"
	//"reflect"
)

func Day15(inputFile string, part int) {
	if part == 0 {
		fmt.Printf("GPS: %d\n", solve(inputFile))
	} else {
		fmt.Println("Not implmenented.")
	}
}

func solve(inputFile string) int {
	ls := util.LineScanner(inputFile)
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

	w.print()

	line, ok = util.Read(ls)
	for ok {
		//fmt.Printf("Number of moves: %d\n", len(line))
		for _, dir := range line {
			//fmt.Printf("----------------------: Robot moves %s\n", string(dir))
			//util.Wait()
			switch dir {
			case '^':
				w.move(robot, -1, 0, false)
			case '<':
				w.move(robot, 0, -1, false)
			case '>':
				w.move(robot, 0, 1, false)
			case 'v':
				w.move(robot, 1, 0, false)
			}

		  //w.print()
		}

		line, ok = util.Read(ls)
	}
		
	sum := 0
	for _, b := range boxes {
		sum += b.gps()
	}
	return sum
}

func (w *Warehouse) print() {
	for i := 0; i < len(w.grid); i++ {
		for j := 0; j < len(w.grid[i]); j++ {
			fmt.Printf(string(w.grid[i][j].symbol()))
		}
		fmt.Printf("\n")
	}
	fmt.Println()
}

func (w *Warehouse) printDouble() {
	for i := 0; i < len(w.grid); i++ {
		for j := 0; j < len(w.grid[i]); j++ {
			fmt.Printf(string(w.grid[i][j].symbol()))
		}
		fmt.Printf("\n")
	}
	fmt.Println()
}

func (w *Warehouse) move(u IUnit, dX, dY int, pushed bool) {
	//fmt.Printf("Move %s (%d,%d) by (%d,%d)\n", reflect.TypeOf(u), u.x(), u.y(), dX, dY)
	if _, ok := u.(*Wall); ok {
		// Walls push back to ensure we don't move something into a wall
		//fmt.Printf("Hit wall %d,%d: wall pushes back on %d,%d in direction %d,%d\n", u.x(), u.y(), u.x()-dX, u.y()-dY, -dX, -dY)
		w.move(w.grid[u.x()][u.y()], -dX, -dY, true)
		w.grid[u.x()][u.y()] = u
	} else if u.x()+dX >= 0 && u.x()+dX < len(w.grid) &&
		u.y()+dY >= 0 && u.y()+dY < len(w.grid[u.x()]) {
		if _, ok := u.(*Robot); ok && !pushed {
			//fmt.Printf("Robot moved: setting original position (%d,%d) to empty\n", u.x(), u.y())
			w.grid[u.x()][u.y()] = &Empty{&Unit{u.x(), u.y()}}
		} 
		next := w.grid[u.x()+dX][u.y() + dY]
		// move this unit
		u.setX(u.x()+dX)
		u.setY(u.y()+dY)
		w.grid[u.x()][u.y()] = u
		
		// Continue moving the lot, including walls
		if _, ok := next.(*Empty); !ok {
			//fmt.Printf("Moving next non-empty element\n")
			w.move(next, dX, dY, true)
		} 
	} else {
		fmt.Printf("Oh no what happened here!?")
	}
}

func (b Unit) gps() int {
	return 100*b.i + b.j
}

type Warehouse struct {
	grid map[int]map[int]IUnit
}

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
type Empty struct {
	*Unit
}

type Robot struct {
	*Unit
}

type Box struct {
	*Unit
}
