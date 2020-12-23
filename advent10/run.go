package advent10

import (
	"fmt"
	"log"
	"math"
	"sort"

	"github.com/igilham/aoc2020go/util"
)

// Run runs the tenth problem
func Run(lines []string) {
	joltages, err := util.ToNumbers(lines)
	if err != nil {
		log.Fatalln(err)
	}
	builtin := util.Max(joltages) + 3
	outlet := 0
	joltages = append(joltages, builtin)
	sort.Ints(joltages)

	// fmt.Println(joltages)

	dist := getDistribution(outlet, joltages)
	// fmt.Printf("distribution: %v\n", dist)
	fmt.Printf("the product of 1-jolt and 3-jolt differences is %v\n", dist[1]*dist[3])

	// part 2 solution stolen from https://davidlozzi.com/2020/12/10/advent-of-code-day-10-check-back-in-629-days/
	// I did not fully understand the solution at time of writing
	variants := countVariations(outlet, joltages)
	// fmt.Printf("variants: %v\n", variants)
	nGroups := len(variants)
	combos := getPossibleCombos(nGroups)
	total := int64(1)
	for i := 0; i < nGroups; i++ {
		total = total * int64((math.Pow(float64(combos[i]), float64(variants[i]))))
	}
	fmt.Printf("total variations: %v\n", total)
}

// Distribution models a table of joltage differences
type Distribution map[int]int

func getDistribution(outlet int, joltages []int) Distribution {
	dist := make(Distribution)
	base := outlet
	for _, n := range joltages {
		difference := n - base
		if difference < 1 || difference > 3 {
			log.Fatalf("difference too great: base=%v, n=%v, diff=%v\n", base, n, difference)
		}
		dist[difference]++
		base = n
	}
	return dist
}

func countVariations(outlet int, joltages []int) Distribution {
	iterationCount := 0
	base := outlet
	dist := make(Distribution)

	for i, jolt := range joltages {
		difference := jolt - base

		if difference == 1 {
			iterationCount++
		} else {
			dist[iterationCount]++
			iterationCount = 0
		}
		base = jolt

		// last row
		if i == len(joltages)-1 {
			dist[iterationCount]++
		}
	}

	return dist
}

func getPossibleCombos(count int) Distribution {
	response := make(Distribution)
	lastThreeResponses := []int{0, 0, 0}
	minValue := 1
	for i := 0; i <= count; i++ {
		possible := util.Sum(lastThreeResponses)
		lastThreeResponses = lastThreeResponses[1:] // shift
		if possible < minValue {
			possible = minValue
		}
		lastThreeResponses = append(lastThreeResponses, possible)
		response[i] = possible
	}
	return response
}
