package day8

import (
	"adventofcode/m/v2/util"
	"fmt"
	"regexp"
)

func Day8(inputFile string, part int) {
	if part == 0 {
		fmt.Printf("Steps: %d\n", followPath(inputFile, false))
	} else {
		fmt.Printf("Visible: %d\n", followPath(inputFile, true))
	}
}

func followPath(inputFile string, ghost bool) int {
	ls := util.LineScanner(inputFile)
	directions, ok := util.Read(ls)
	util.Read(ls)	
	
	endInA := []*Node{}
	nodes := make(map[string]*Node, 0)
	line, ok := util.Read(ls)
	for ok {
		parseDirections := regexp.MustCompile(`([A-Z,0-9]{3}) = \(([A-Z,0-9]{3}), ([A-Z,0-9]{3})\)`)
		m := parseDirections.FindAllStringSubmatch(line, 3)
		
		if _, found := nodes[m[0][1]]; !found {
			nodes[m[0][1]] = &Node{m[0][1], nil,nil}
		}
		if _, found := nodes[m[0][2]]; !found {
			nodes[m[0][2]] = &Node{m[0][2], nil,nil}
		}
		if _, found := nodes[m[0][3]]; !found {
			nodes[m[0][3]] = &Node{m[0][3], nil,nil}
		}

		nodes[m[0][1]].left = nodes[m[0][2]]
		nodes[m[0][1]].right = nodes[m[0][3]]

		if m[0][1][2] == 'A' {
			endInA = append(endInA, nodes[m[0][1]])
		}
		line, ok = util.Read(ls)
	}

	
	var starts []*Node
	if ghost {
		starts = endInA
	} else {
		starts = []*Node{nodes["AAA"]}
	}
	print(starts)
	var steps, mult int
	mult = 1
	for i := range starts {
		steps = 0
		for true {
			for _, d := range []rune(directions) {
				switch d {
				case 'L':
					starts[i] = starts[i].left
				case 'R': 
				  starts[i] = starts[i].right
				}
				steps++
				if (ghost && starts[i].name[2] == 'Z') || starts[i].name == "ZZZ" {
					break
				}
			}
			if (ghost && starts[i].name[2] == 'Z') || starts[i].name == "ZZZ" {
				mult = mult*steps/gcd(mult, steps)
				break
			}
		}

		fmt.Printf("%s in %d steps\n", starts[i].name, steps)
	}
		
	return mult
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}
func print(nodes []*Node) {
	for _, n := range nodes { 
		fmt.Printf("  %s |", n.name)
	}
	fmt.Printf("\n->")
}
type Node struct {
	name string
	left *Node
	right *Node
}
