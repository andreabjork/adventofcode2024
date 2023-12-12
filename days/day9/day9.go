package day9

import (
	"adventofcode/m/v2/util"
	"fmt"
	"strconv"
	"strings"
)

func Day9(inputFile string, part int) {
	if part == 0 {
		fmt.Printf("Sum of extrapolated: %d\n", solve(inputFile, false))
	} else {
		fmt.Printf("Sum of extrapolated: %d\n", solve(inputFile, true))
	}
}

func solve(inputFile string, reverse bool) int {
	ls := util.LineScanner(inputFile)
	sum := 0
	sequence, ok := util.Read(ls)
	for ok {
		
		var prev, next *Node
		if reverse {
			prev = left(sequence)
			next = &Node{-1, prev, nil, false, -1}
			sum += next.value(neg)
		} else {
			prev = right(sequence)
			next = &Node{-1, prev, nil, false, -1}
			sum += next.value(pos)
		}
	  
		sequence, ok = util.Read(ls)
	}

	return sum
}

func right(sequence string) *Node {
	var row int 
	var prev *Node
	for _, s := range strings.Split(sequence, " ") {
		d, _ := strconv.Atoi(s)
		if prev != nil {
			row = prev.row
		}
		prev = &Node{d, prev, nil, true, row + abs(d)}
	}

	return prev
}

func left(sequence string) *Node {
	var row int 
	var prev *Node
	seq := strings.Split(sequence, " ")
	for i := len(seq)-1; i >= 0; i--  {
		d, _ := strconv.Atoi(seq[i])
		if prev != nil {
			row = prev.row
		}
		prev = &Node{d, prev, nil, true, row + abs(d)}
	}

	return prev
}

func (n *Node) bot(op func(int) int) *Node {
	if n.bottom != nil {
		return n.bottom
	} else if n.left == nil { 
		return nil // No bottom unless both parents exist
	}

	d := op(n.val - n.left.value(op))
	l := n.left.bot(op)
	var row int
	if l != nil {
		row = l.row
	}
	n.bottom = &Node{d, l, nil, n.isSet, row+abs(d)}

	// Force 0 if the entire row is 0
	if !n.bottom.isSet && n.bottom.left.row == 0 {
		n.bottom.val = 0
		n.bottom.isSet = true
	}

	return n.bottom
}

func (n *Node) value(op func(int) int) int {
	if n.isSet {
		return n.val
	} 

	n.val = n.left.value(op) + op(n.bot(op).value(op))
	n.isSet = true
	return n.val
}

type Node struct {
	val        int
	left 			 *Node
	bottom 		 *Node
	isSet 		 bool
	row        int // Sum of absolute values, to detect 0 rows
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func pos(x int) int {
	return x
}

func neg(x int) int {
	return -x
}