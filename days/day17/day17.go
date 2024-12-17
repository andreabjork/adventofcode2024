package day17

import (
	"adventofcode/m/v2/util"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func Day17(inputFile string, part int) {
	if part == 0 {
		fmt.Printf("Output: %s\n", runProgram(inputFile))
	} else {
		fmt.Println("Not implmenented.")
	}
}

const DEBUG = false

func runProgram(inputFile string) string {
	ls := util.LineScanner(inputFile)
	regA, _ := util.Read(ls)
	regB, _ := util.Read(ls)
	regC, _ := util.Read(ls)
	_, _ = util.Read(ls)
	prog, _ := util.Read(ls)

	A, _ := strconv.Atoi(strings.TrimPrefix(regA, "Register A: "))
	B, _ := strconv.Atoi(strings.TrimPrefix(regB, "Register B: "))
	C, _ := strconv.Atoi(strings.TrimPrefix(regC, "Register C: "))
	program := []int{}
	for _, o := range strings.Split(strings.TrimPrefix(prog, "Program: "), ",") {
		val, _ := strconv.Atoi(o)
		program = append(program, val)
	}

	c := &Computer{A, B, C, program, 0, "", nil}
	c.initialize()
	c.run()
	
	return c.strip
}

type Computer struct {
	A, B, C int 
	program []int
	pointer int
	strip string
	instruction map[int]func(o int)
}

func (c *Computer) status() {
	fmt.Printf("Register A: %d\n", c.A)
	fmt.Printf("Register B: %d\n", c.B)
	fmt.Printf("Register C: %d\n", c.C)

	for _, v := range c.program {
		fmt.Printf("%d ", v)
	}
	fmt.Printf("\n")
	for i, _ := range c.program {
		if i == c.pointer {
			fmt.Printf("^ ")
		} else if c.pointer >= len(c.program) {
			fmt.Printf("HALT")
			break
		} else {
			fmt.Printf("  ")
		}
	}

	fmt.Println()
	fmt.Printf("Strip: %s\n", c.strip)
}

func (c *Computer) run() {
	if DEBUG {
		c.status()
		time.Sleep(3*time.Second)
	}

	if c.pointer >= len(c.program) {
		return
	}
	c.instruction[c.program[c.pointer]](c.program[c.pointer+1])
	c.run()
}

func (c *Computer) combo(o int) int {
	switch(o) {
	case 0, 1, 2, 3: { return o }
	case 4: { return c.A }
	case 5: { return c.B }
	case 6: { return c.C }
	}

	if true {
		panic("invalid operand")
	}
	
	return -1
}

func (c *Computer) initialize() {
	c.instruction = map[int]func(o int) {
		0: func(o int) { c.A = c.A/util.Pow(2,c.combo(o)); c.pointer+=2},
		1: func(o int) { c.B = c.B ^ o; c.pointer+=2},
		2: func(o int) { c.B = c.combo(o) % 8; c.pointer+=2},
		3: func(o int) { 
			if c.A != 0 {
				c.pointer = o 
			} else {
				c.pointer += 2
			}
		},
		4: func(o int) { c.B = c.B ^ c.C; c.pointer+=2 },
		5: func(o int) { c.strip += fmt.Sprintf("%d,", c.combo(o)%8); c.pointer+=2 },
		6: func(o int) { c.B = c.A/util.Pow(2,c.combo(o)); c.pointer += 2},
		7: func(o int) { c.C = c.A/util.Pow(2,c.combo(o)); c.pointer += 2},
	} 
}