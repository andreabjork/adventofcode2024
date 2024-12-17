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
		fmt.Printf("Simulating")
		simulate()
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
	c.run("")
	
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
	for i  := range c.program {
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

func (c *Computer) run(expect string) bool {
	if DEBUG {
		c.status()
		time.Sleep(3*time.Second)
	}

	if c.pointer >= len(c.program) {
		return true
	}
	c.instruction[c.program[c.pointer]](c.program[c.pointer+1])
	
	if expect != "" {
		if !strings.HasPrefix(expect, c.strip) {
			return false
		}
	}
	return c.run(expect)
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

// 																										A VALID RANGE FOR A
// ----------------------------------------------------------------------
// A changes value only when (0,3) is called, and it always runs A = A / 8
// We need to print value for A 16 times, changing it at least 14 times.
// So we search from pow(8, 14)
// Each pass through of the program updates the value in A exactly once,
// prints exactly one output and moves the pointer to the beginning once.
// 
// The program halts when the value in A reaches 0. We need to print exactly 
// 16 digits and halt thereafter. 
//
// The value in A after one pass through of the program is A/8. 
// The value in A after two pass throughs is A/8*8
// .
// .
// .
// The value in A after 15 pass throughs is A/pow(8, 15)
// The value in A after 16 pass throughs is A/pow(8, 16)
// For this to become 0 we need A < pow(8,16).
// Importantly, A cannot be 0 after only 15 pass throughs so A >= pow(8,15) 
//
// 																											RESTRICTIONS ON A
// -----------------------------------------------------------------------
// Some conditions can further restrict our available A values.
// 
// 2,4: assigns B = A%8
// 1,5: assigns B = B xor 101
// 7,5: C = A/pow(2,B)
// 1,6: assigns B = B xor 6 
// 4,2: assigns B = B xor C   
// 5,5: prints B%8
// 
// In other words, we must have ((((A%8) xor 5) xor 6) xor A/2^B) %8 = 2
// (A % 8) xor X % 8 = 2
// this means the last 3 bytes are 010, which means we can have
// 
// xor 110 x 101 = 0
// 
// xor 5 xor 6 = 101 xor 110 = 010
// 
// This means the last 3 bits of ((((A%8) xor 5) xor 6) xor A/2^(A%8 xor 5 xor 6)) must be 010.

// Then
// 010 xor 110 = 100
// then
// 100 xor 101 = 001
// 
// So we must have A % 8 == 1 
// Then, we must have 
// 
// lets see
// 000 xor 101 = 101 = 5 mod 8
// 001 xor 101 = 100 = 4 mod 8
// 010 xor 101 = 111 = 7 mod 8
// 100 xor 101 = 001 = 1 mod 8
// 101 xor 101 = 000 = 0 mod 8
// 110 xor 101 = 011 = 3 mod 8
// 111 xor 101 = 010 = 2 mod 8   <- A%8 must therefore be 7
func simulate() {
	found := false
	fmt.Printf("Entering loop\n")
	i := 0
	a := util.Pow(8,15)+7+8*i
	for ; a < util.Pow(8,16); i++ {
		a = util.Pow(8,15)+6+8*i
		c := &Computer{a, 0, 0, []int{2,4,1,5,7,5,1,6,4,2,5,5,0,3,3,0}, 0, "", nil}
		c.initialize()
		expect := "2,4,1,5,7,5,1,6,4,2,5,5,0,3,3,0,"

		halted := c.run(expect)
		if halted {
			fmt.Printf("Found one such a")
			fmt.Printf("A = %d\n", a)
			found = true
			break
		}
		fmt.Printf("%d/%d - %s\n", a-util.Pow(8,15), util.Pow(8,16)-util.Pow(8,15), c.strip)
	}

	if found {
		fmt.Printf("We finished")
	}
}
