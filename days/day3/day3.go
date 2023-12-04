package day3

import (
	"adventofcode/m/v2/util"
	"fmt"
	"io"
	"strconv"
)

func Day3(inputFile string, part int) {
	if part == 0 {
    sum, _ := partsAndGears(inputFile)
		fmt.Printf("Sum of parts: %d\n", sum)
	} else {
    _, gears := partsAndGears(inputFile)
		fmt.Printf("Gear values: %d\n", multiply(gears))
	}
}

type Part struct {
  value   int
  length  int
  isValid  bool
  i       int
  j       int
}

type Gear struct {
  typ       rune 
  adjacents map[*Part]bool
}

func partsAndGears(inputFile string) (int, []*Gear) { 
  ls := util.Reader(inputFile)

  var marked map[int]map[int][]*Gear = make(map[int]map[int][]*Gear) // marked[i][j] = true if (i,j) is adjacent to a symbol
  var gears []*Gear = []*Gear{}
  var parts []*Part = []*Part{} 

  r, sz, err := ls.ReadRune()
  
  var i, j int = 0, 0
  var num []rune
  for err != io.EOF && sz > 0 {
    if r == '\n' {
      parts = addPart(parts, num, i, j)
      num = []rune{}
      i++
      j = -1 
    } else if r == '.' {
      parts = addPart(parts, num, i, j)
      num = []rune{}
    } else if r >= '0' && r <= '9' {
      num = append(num, r)
    } else { // Symbols
      parts = addPart(parts, num, i, j)
      num = []rune{}

      g := &Gear{r, make(map[*Part]bool)}
      for ii := util.Max(0, i-1); ii <= i+1; ii++ {
        for jj := util.Max(0, j-1); jj <= j+1; jj++ {
          if marked[ii] == nil {
            marked[ii] = make(map[int][]*Gear)
          }
          marked[ii][jj] = append(marked[ii][jj], g)
        }
      }
      // Track * gears
      if r == '*' {
        gears = append(gears, g)
      }
    }

    r, sz, err = ls.ReadRune()
    j++
  }

  // Find which parts are valid (adjacent to gears) and mark the part on the gear
  var sum = 0
  for _, n := range parts {
    for d := 0; d < n.length; d++ {
        if gs, ok := marked[n.i][n.j+d]; ok && len(gs) > 0 {
          for _, g := range gs {
            g.adjacents[n] = true
          }
          n.isValid = true
        }
      }
    if n.isValid {
        sum += n.value
    }
  }
  return sum, gears
}

func multiply(gears []*Gear) int {
  sum := 0
  for _, g := range gears {
    if g.typ == '*' && len(g.adjacents) > 1 {
      mult := 1
      for adj, _ := range g.adjacents {
        mult = mult*adj.value
      }
      sum += mult
    }
  }

  return sum
}

func addPart(parts []*Part, num []rune, i int, j int) []*Part {
  if len(num) > 0 {
    x, _ := strconv.Atoi(string(num))
    parts = append(parts, &Part{x,len(num),false, i, j-len(num)})
  }
  return parts 
}
