package day3

import (
  "adventofcode/m/v2/util"
	"fmt"
	"io"
	"strconv"
)

func Day3(inputFile string, part int) {
	if part == 0 {
		fmt.Printf("Sum of parts: %d\n", sumParts(inputFile))
	} else {
		fmt.Printf("Sum of priorities of all groups: %d\n", sumParts(inputFile))
	}
}

func sumParts(inputFile string) int { 
  ls := util.Reader(inputFile)

  var marked map[int]map[int]bool = make(map[int]map[int]bool) // marked[i][j] = true if (i,j) is adjacent to a symbol
  var numbers map[int]map[int][]int = make(map[int]map[int][]int) // numbers[i][j] = x,d where x starts at position (i,j) and x is d digits
  r, sz, err := ls.ReadRune()
  
  var i, j int = 0, 0
  var num []rune
  for err != io.EOF && sz > 0 {
    if r == '\n' {
      addNumber(numbers, num, i, j)
      num = []rune{}
      i++
      j = -1 
    } else if r == '.' {
      addNumber(numbers, num, i, j)
      num = []rune{}
    } else if r >= '0' && r <= '9' {
      num = append(num, r)
    } else { // Symbols
      addNumber(numbers, num, i, j)
      num = []rune{}

      for ii := util.Max(0, i-1); ii <= i+1; ii++ {
        for jj := util.Max(0, j-1); jj <= j+1; jj++ {
          if marked[ii] == nil {
            marked[ii] = make(map[int]bool)
          }
          marked[ii][jj] = true
        }
      }
    }

    r, sz, err = ls.ReadRune()
    j++
  }

  var parts = []int{}
  var sum = 0
  var found bool = false
  for i, m := range numbers {
    for j, mm := range m {
      found = false
      for d := 0; d < mm[1]; d++ {
        if marked[i][j+d] {
          parts = append(parts, mm[0])
          sum += mm[0]
          found = true
          break
        }
      }
    }
  }
  return sum 
}

func addNumber(numbers map[int]map[int][]int, num []rune, i int, j int) {
  if len(num) > 0 {
    x, _ := strconv.Atoi(string(num))
    if numbers[i] == nil { 
      numbers[i] = make(map[int][]int)
    }
    numbers[i][j-len(num)] = []int{x,len(num)}
  }
}