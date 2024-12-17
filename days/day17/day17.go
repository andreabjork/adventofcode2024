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

// Observations:
// 
// * The value of A only changes with instruction 0.
// * The input to instruction 0 is always 3. Combo(3) = 3 and thus
//   the value of A only changes via: A/2³ = A/8
//
// * Once the value of A becomes 0, it stays at 0. We can assume A is non-zero 
//   before, during and after the run of our program.
//
// * Before the value of A becomes 0, the behaviour of the
//   pointer is deterministic and independent of the exact value of A.
//
// * An output is printed if and only if we run instruction '5'. This means,
//   for an output string 0,3,5,4,3,0, we need to run instruction '5' exactly 6 times.
//
// * The instruction '5' always takes o = 4 as an argument. In other words, it always
//   uses the value of A % 8 for printing.
//
//
// 
// Let us represent A[i] := the value of A at level i in the call stack.
// We look for the value of A[0]. We know that
// 
//	A[6] = 0 mod 8
//  A[5] = 3 mod 8
//  A[4] = 4 mod 8
//  A[3] = 5 mod 8
//  A[2] = 3 mod 8
//  A[1] = 0 mod 8
//  A[0] = ?
//
// Finally, we have the equations for A where A is the original value:
//
// A / 8 = 0 mod 8
// A / 8*8 = 3 mod 8
// A / 8*8*8 = 5 mod 8
// A / 8*8*8*8  = 4 mod 8
// A / 8*8*8*8*8 = 3 mod 8 
// A / 8*8*8*8*8*8 = 0 mod 8 = A / 36*36*36 = 0 mod 8
//
//
// Now, if A < 8*36*36, the value of A would quickly become 0. 
// We can assume 8*36*36 < A < 36*36*36*36 (394149888 possibilities)
// 
// By multiplying through these equations we also get
// 
// A / 12 = 0 mod 8
// A / 144 = 3 mod 8
// A / 12*144 = 5 mod 8
// A / 144*144 = 4 mod 8 -> A / 12*144 -> 48 mod 8 = 0 mod 8
// A / 12*144*144 = 3 mod 8  -> A / 144*144 = 36 mod 8 = 4 mod 8 


// New try
// A / 144*144 = 4 mod 8 -> 144*4 = 576 
// A / 144*144*12 -> bleh
// A / 144*144*144*144 = 3 mod 8  -> A / 12*144*144 = 36 mod 8 = 4 mod 8 

// A changes value only when (0,3) is called, and it always runs A = A / 8
// We need to print value for A 16 times, changing it at least 14 times.
// So we search from pow(8, 14)
func simulate() {
	found := false
	fmt.Printf("Entering loop\n")
	for a := util.Pow(8,15); a < util.Pow(8,15)+36*36*36*36; a++ {
		//fmt.Printf("a=%d\n", a)
		// Only consider values of a that satisfy our criteria
		if true {
		//	(a / 8*36*36) % 8 == 3 &&
		//(a / 36*36) % 8 == 4 &&
		//(a / 36*8) % 8 == 5 &&
		//(a / 36) % 8 == 3 &&
		//(a / 8) % 8 == 0 {
				c := &Computer{a, 0, 0, []int{2,4,1,5,7,5,1,6,4,2,5,5,0,3,3,0}, 0, "", nil}
				c.initialize()
				c.run()

				if c.strip == "2,4,1,5,7,5,1,6,4,2,5,5,0,3,3,0," {
			 		fmt.Printf("Found one such a")
					fmt.Printf("A = %d\n", a)
					found = true
					break
				}
			}
	}

	if found {
		fmt.Printf("We finished")
	}
}
//
// 3 mod 8 x 12 = 36 mod 8 = 4 mod 8
// 4 mod 8 x 12 = 48 mod 8 = 0 mod 8
// 0 mod 8 x 12 = 96 mod 8 = 0 mod 8  

// 
// 
//
// We will create reverse instructions to find the value of A:
// 
// program: 0,3,5,4,3,0 
//              ^
// strip: 0,3,5,4,3,0 
// 
// Note that the input o   
// 
// 
