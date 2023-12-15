package day14

import (
	"adventofcode/m/v2/util"
	"fmt"
)

func Day14(inputFile string, part int) {
	switch part {
	case 0:
		fmt.Printf("Weight: %d\n", tilt(inputFile))
	case 1:
		fmt.Printf("Grains of sand: %d\n", tiltOften(inputFile))
	}
}

func tilt(inputFile string) int {
	ls := util.LineScanner(inputFile)
	line, ok := util.Read(ls)

	platform := [][]rune{}
	for ok {
		platform = append(platform, []rune(line))

		line, ok = util.Read(ls)
	}
	
	north(platform)
	return load(platform)
}

func tiltOften(inputFile string) int {
	ls := util.LineScanner(inputFile)
	line, ok := util.Read(ls)

	platform := [][]rune{}
	for ok {
		platform = append(platform, []rune(line))

		line, ok = util.Read(ls)
	}
	
	
	// Cycling 1000 times shows a repeating cycle of length 22,
	// which starts occurring from (at least) round 186.
	// We have (1000-186)%22 == 0
	// as well as (1000000000 - 186)%22 == 0, so the load will be 
	// the same after both number of cycles
	for i := 1; i <= 1000; i++ {
		load := cycle(platform)
		if (i - 186) % 22 == 0 {
			fmt.Println(i, load)
		}
	}
	return load(platform)
}

func cycle(p [][]rune) int {
	north(p)
	west(p)
	south(p)
	east(p)

	return load(p)
}

func load(p [][]rune) int {
	weight := 0
	for i := range p {
		for j := range p {
			//fmt.Printf("%s", string(p[i][j]))
			if p[i][j] == 'O' {
				weight += len(p)-i
			}
		}
		//fmt.Println()
	}

	return weight 
}

func north(platform [][]rune) {
	for i := range platform {
		for j := range platform[i] {
			for k := 1; i-k >= 0; k++ {
				if platform[i-k+1][j] == 'O' && platform[i-k][j] == '.' {
					platform[i-k+1][j] = '.'
					platform[i-k][j] = 'O'
				} else {
					break
				}
			}
		}
	}
}

func south(platform [][]rune) {
	for i := len(platform)-1; i >= 0; i-- {
		for j := range platform[i] {
			for k := 0; i+k+1 < len(platform); k++ {
				if platform[i+k][j] == 'O' && platform[i+k+1][j] == '.' {
					platform[i+k][j] = '.'
					platform[i+k+1][j] = 'O'
				} else {
					break
				}
			}
		}
	}
}

func west(platform [][]rune) {
	for j := range platform[0] {
		for i := range platform {
			for k := 1; j-k >= 0; k++ {
				if platform[i][j-k+1] == 'O' && platform[i][j-k] == '.' {
					platform[i][j-k+1] = '.'
					platform[i][j-k] = 'O'
				} else {
					break
				}
			}
		}
	}
}

func east(platform [][]rune) {
	for j := len(platform[0])-1; j >= 0; j-- {
		for i := range platform {
			for k := 0; j+k+1 < len(platform[0]); k++ {
				if platform[i][j+k] == 'O' && platform[i][j+k+1] == '.' {
					platform[i][j+k] = '.'
					platform[i][j+k+1] = 'O'
				} else {
					break
				}
			}
		}
	}
}