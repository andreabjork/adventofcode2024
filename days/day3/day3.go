package day3

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func Day3(inputFile string, part int) {
	if part == 0 {
		fmt.Printf("All muls: %d\n", mulSimple(inputFile))
	} else {
		fmt.Printf("Adjusted muls: %d\n", mul(inputFile))
	}
}

func mulSimple(inputFile string) int {
	txt, _ := os.ReadFile(inputFile)
  instruction := string(txt)

	sum := 0
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := re.FindAllStringSubmatch(instruction, 20000)
	for _, m := range matches {
		a, _ := strconv.Atoi(m[1])
		b, _ := strconv.Atoi(m[2])
		sum += a * b
	}
	return sum
}

func mul(inputFile string) int {
	txt, _ := os.ReadFile(inputFile)
  instruction := string(txt)
  dos := []string{}
  s := ""
  ignore := false
  for i := 0; i < len(instruction); i++ {
    if !ignore {
      s += string(instruction[i])
    }
  
    // We ignore any dos() until the first don't() has been found
    if ignore && i+4 < len(instruction) && instruction[i:i+4] == "do()" {
      ignore = false
    } else if !ignore && i+7 < len(instruction) && instruction[i:i+7] == "don't()" {
      ignore = true
      dos = append(dos, s)
      s = ""
    }
  }

	dos = append(dos, s)


	sum := 0
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	for _, d := range dos {
		matches := re.FindAllStringSubmatch(d, 20000)
		for _, m := range matches {
			a, _ := strconv.Atoi(m[1])
			b, _ := strconv.Atoi(m[2])
			sum += a * b
		}
	}
	return sum
}
