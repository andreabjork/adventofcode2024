package day6

import (
	"adventofcode/m/v2/util"
	"fmt"
	"math"
	"regexp"
	"strings"
)

func Day6(inputFile string, part int) {
	if part == 0 {
		fmt.Printf("Ways to win: %d\n", sol(inputFile, false))
	} else {
		fmt.Printf("Ways to win: %d\n", sol(inputFile, true))
	}
}

func sol(inputFile string, ignoreSpaces bool) int {
	ls := util.LineScanner(inputFile)
	whitespace := regexp.MustCompile(`\s+`)
	time, _ := util.Read(ls) 
	distance, _ := util.Read(ls) 

	var t, d []float64
	if ignoreSpaces {
		t = util.AsFloats(strings.Split(whitespace.ReplaceAllString(time, ""), "Time:")[1:])
		d = util.AsFloats(strings.Split(whitespace.ReplaceAllString(distance, ""), "Distance:")[1:])
	} else {
		t = util.AsFloats(strings.Split(time, " ")[1:])
		d = util.AsFloats(strings.Split(distance, " ")[1:])
	}

  fmt.Printf("%+v\n", t)
  fmt.Printf("%+v\n", d)
	// Let v := velocity, seconds holding the button
	//     t := total time of race
	//     d := minimum distance to cover
	// 
	// Then (t-v) is the remaining time after pressing button and
	// we need v(t-v) > d
	// i.e.    -v^2 + v*t - d > 0
	// and thus
	var mult int = 1
	var v1, v2 float64
	for i, _ := range t {
		v1 = -0.5*(-t[i] - math.Sqrt(t[i]*t[i] - 4*d[i]))
 	  v2 = -0.5*(-t[i] + math.Sqrt(t[i]*t[i] - 4*d[i]))
		
		var iv1, iv2 int = int(v1), int(v2)
		if float64(iv1) == v1 {
			iv1--
		}
		mult *= (iv1-iv2)
	}

  return mult 
}
