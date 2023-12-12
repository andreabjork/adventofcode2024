package day9

import (
	"adventofcode/m/v2/util"
	"fmt"
	"strconv"
	"strings"
)

func Day9(inputFile string, part int) {
	if part == 0 {
		fmt.Printf("Sum of extrapolated: %d\n", solve(inputFile))
	} else {
		fmt.Printf("Sum of extrapolated: %d\n", solve(inputFile))
	}
}

func solve(inputFile string) int {
	ls := util.LineScanner(inputFile)
	sum := 0
	sequence, ok := util.Read(ls)
	for ok {
		var prev *Node
		for _, s := range strings.Split(sequence, " ") {
			d, _ := strconv.Atoi(s)
			prev = &Node{d, prev, nil, nil, nil, true}
		}

		next := &Node{-1, prev, nil, nil, nil, false}
		sum += next.value()
		fmt.Println(next.value())

		sequence, ok = util.Read(ls)
	}

	return sum
}

func (n *Node) value() int {
	// n's value is known
	if n.val != -1 {
		return n.val
	}

	// Evaluate based on parents, only if both exist with values:
	if n.topLeft != nil && n.topRight != nil && n.topLeft.isSet && n.topRight.isSet {
		n.val = n.topRight.val - n.topLeft.val
	// Add left parent if missing, then evaluate based on parents:
	} else if n.topLeft == nil && n.topRight != nil && n.topRight.isSet {
		n.topLeft = &Node{-1, nil, n.topRight.topLeft.left, n.topRight.topLeft, nil, false} 
		n.topLeft.val = n.topLeft.value()
		n.val = n.topRight.val - n.topLeft.val
	// Otherwise, evaluate based on left and bottom left, adding both if needed:
	} else {
		if n.left == nil {
			n.left = &Node{-1, nil, n.topLeft.left, n.topLeft, nil, false}
			n.left.val = n.left.value()
		}
		if n.left.val != 0 && n.bottomLeft == nil {
			n.bottomLeft = &Node{-1, nil, n.left, n, nil, false}
			n.bottomLeft.val = n.bottomLeft.value()
			n.val = n.left.val + n.bottomLeft.val
		} else if n.left.val == 0 {
			n.val = 0
		}
	}

	n.isSet = true
	return n.val
}

type Node struct {
	val        int
	left			 *Node
	topLeft  	 *Node
	topRight   *Node
	bottomLeft *Node
	isSet 		 bool
}
