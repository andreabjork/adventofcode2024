package day2

import (
	"adventofcode/m/v2/util"
	"fmt"
	"regexp"
	"strings"
	"strconv"
)

func Day2(inputFile string, part int) {
	if part == 0 {
		fmt.Printf("Sum of IDs: %d\n", validateGames(inputFile))
	} else {
		fmt.Printf("Power of Minimums: %d\n", powerOfGames(inputFile))
	}
}

func validateGames(inputFile string) int {
	ls := util.LineScanner(inputFile)
	line, ok := util.Read(ls)

	bag := map[string]int{"red": 12, "green": 13, "blue": 14}
	parseGame := regexp.MustCompile(`Game (\d+):(( \d+ (blue|red|green),*)*;*)*`)
	parseRounds := regexp.MustCompile(`(\d+)\s?(blue|red|green)+,*\s?(\d*)\s?(blue|red|green)*,*\s?(\d*)\s?(blue|red|green)*`)
	sum := 0
	for ok {
		game := parseGame.FindAllStringSubmatch(line, 10)
		id, _ := strconv.Atoi(game[0][1])

		invalid := false
		for _, round := range strings.Split(game[0][0], ";") {
			actions := parseRounds.FindAllStringSubmatch(round, 10)
		  for i := 1; i < len(actions[0]); i+=2 {
				count, _ := strconv.Atoi(actions[0][i])
				color := actions[0][i+1]
			  	
				if count > bag[color] {
					invalid = true
					break	
				}
			}	

			if invalid {
				break
			} 
		}

		if !invalid {
			sum += id 		
		}

		line, ok = util.Read(ls)
	}

	return sum
}

func powerOfGames(inputFile string) int {
	ls := util.LineScanner(inputFile)
	line, ok := util.Read(ls)

	parseGame := regexp.MustCompile(`Game (\d+):(( \d+ (blue|red|green),*)*;*)*`)
	parseRounds := regexp.MustCompile(`(\d+)\s?(blue|red|green)+,*\s?(\d*)\s?(blue|red|green)*,*\s?(\d*)\s?(blue|red|green)*`)
	sum := 0
	for ok {
		game := parseGame.FindAllStringSubmatch(line, 10)

		minBag := map[string]int{"red": 0, "green": 0, "blue": 0}
		for _, round := range strings.Split(game[0][0], ";") {
			actions := parseRounds.FindAllStringSubmatch(round, 10)
		  for i := 1; i < len(actions[0]); i+=2 {
				count, _ := strconv.Atoi(actions[0][i])
				color := actions[0][i+1]
			  	
				if count > minBag[color] {
					minBag[color] = count
				}
			}	
		}

		sum += minBag["red"]*minBag["green"]*minBag["blue"]

		line, ok = util.Read(ls)
	}

	return sum
}
