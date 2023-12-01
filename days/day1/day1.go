package day1

import (
	"adventofcode/m/v2/util"
	"fmt"
	"strconv"
)

func Day1(inputFile string, part int) {
	if part == 0 {
		fmt.Printf("Calibration sum: %d\n", calibrate(inputFile, false))
	} else {
		fmt.Printf("Calibration sum: %d\n", calibrate(inputFile, true))
	}
}

var lookup = map[string]rune{
	"one": '1',
	"two": '2',
	"three": '3',
	"four": '4',
	"five": '5',
	"six": '6',
	"seven": '7',
	"eight": '8',
	"nine": '9',
}

func calibrate(inputFile string, allowWords bool) int {
	ls := util.LineScanner(inputFile)
	line, ok := util.Read(ls)

	sum := 0
	for ok {
		var first, last rune
		first = 'u'
		runes := []rune(line)
		for i, r := range runes {
			// Check if integer
			_, err := strconv.Atoi(string(r))
			if err == nil && first == 'u' {
				first = r
			}

			if err == nil {
				last = r
				continue
			}

			if allowWords {
				// Check 3-size window
				if r, ok := lookup[string(runes[i:util.Min(i+3, len(runes))])]; ok {
					if first == 'u' {
						first = r
					}	
					last = r
					continue
				}
				// Check 4-size window
				if r, ok := lookup[string(runes[i:util.Min(i+4, len(runes))])]; ok {
					if first == 'u' {
						first = r
					}	
					last = r
					continue
				}
				// Check 5-size window
				if r, ok := lookup[string(runes[i:util.Min(i+5, len(runes))])]; ok {
					if first == 'u' {
						first = r
					}	
					last = r
					continue
				}
			}			
		}

		val, err := strconv.Atoi(string([]rune{first, last}))
		if err == nil {
			sum += val
		}

		line, ok = util.Read(ls)
	}

	return sum
}
