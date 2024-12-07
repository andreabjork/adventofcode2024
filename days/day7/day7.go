package day7

import (
	"adventofcode/m/v2/util"
	"fmt"
	"strings"
	"strconv"
)

func Day7(inputFile string, part int) {
	if part == 0 {
		fmt.Printf("Sum of equations: %d\n", search(inputFile, false))
	} else {
		fmt.Printf("Sum of equations: %d\n", search(inputFile, true))
	}
}


func search(inputFile string, withConcat bool) int {
	ls := util.LineScanner(inputFile)
	line, ok := util.Read(ls)

	var parts []string
	var t []int
	var n *Node 
	var sum, m int
	for ok {
		parts = strings.Split(line, ": ")
		m, _ = strconv.Atoi(parts[0])
		t = asIntArray(parts[1])
		n = &Node{
			t[0], 
			t,
			0, 
			nil,
			nil,
			nil,
			nil,
		}
		if !withConcat && n.binary(m) {
			sum += m
		} else if withConcat && n.ternary(m) {
			sum += m
		}

		line, ok = util.Read(ls)
	}
	return sum
}


type Node struct {
	val int 
	ops []int 
	level int
	left *Node
	center *Node
	right *Node
	parent *Node
}

func (n *Node) l() *Node {
	if n.level == len(n.ops)-1 {
		panic("tried to reach left node of a leaf")
	}
	if n.left == nil {
		n.left = &Node{
			n.val + n.ops[n.level+1],
			n.ops,
			n.level+1,
			nil,
			nil,
			nil,
			n,
		}
	}
	return n.left
}

func (n *Node) c() *Node {
	if n.level == len(n.ops)-1 {
		panic("tried to reach left node of a leaf")
	}
	if n.center == nil {
		x, _ := strconv.Atoi(fmt.Sprintf("%d%d", n.val, n.ops[n.level+1]))
		n.center = &Node{
			x,
			n.ops,
			n.level+1,
			nil,
			nil,
			nil,
			n,
		}
	}
	return n.center
}

func (n *Node) r() *Node {
	if n.level == len(n.ops)-1 {
		panic("tried to reach right node of a leaf")
	}
	if n.right == nil {
		n.right = &Node{
			n.val * n.ops[n.level+1],
			n.ops,
			n.level+1,
			nil,
			nil,
			nil,
			n,
		}
	}
	return n.right
}

func (n *Node) leaf() bool {
	return n.level == len(n.ops)-1
}

func (n *Node) binary(m int) bool {
	if n.val == m {
		return true
	} else if n.val > m || n.leaf() {
		return false
	}
	return n.r().binary(m) || n.l().binary(m)
}

func (n *Node) ternary(m int) bool {
	if n.val == m && n.leaf() {
		return true
	} else if n.val > m || n.leaf() {
		return false
	} 
	return n.r().ternary(m) || n.c().ternary(m) || n.l().ternary(m)
}

func asIntArray(line string) []int {
	a := []int{}
	for _, x := range strings.Split(line, " ") {
		xx, _ := strconv.Atoi(x)
		a = append(a, xx)
	}
	return a
}