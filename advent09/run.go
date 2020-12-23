package advent09

import (
	"fmt"
	"log"

	"github.com/igilham/aoc2020go/util"
)

// Run runs the ninth problem
func Run(lines []string) {
	nums, err := util.ToNumbers(lines)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	firstBadValue := findFirstInvalid(nums, 25) // use 5 for test input
	fmt.Printf("first bad value: %v\n", firstBadValue)
	contig := findContiguousSum(nums, firstBadValue)
	fmt.Printf("contiguous sum: (%v)\n", contig)
	smallest := min(contig)
	largest := max(contig)
	fmt.Printf("smallest=%v, largest=%v, weakness=%v", smallest, largest, smallest+largest)
}

func findFirstInvalid(nums []int, preambleLength int) int {
	preamble := nums[0:preambleLength]
	for _, n := range nums[preambleLength:] {
		sums := validSums(preamble)
		if !contains(sums, n) {
			return n
		}
		preamble = append(preamble[1:], n)
	}
	return 0
}

func findContiguousSum(nums []int, target int) []int {
	for i := range nums {
		for j := range nums[i:] {
			contig := nums[i : i+j]
			s := sum(contig)
			// fmt.Printf("sum=%v\n", s)
			if s == target {
				return contig
			}
		}
	}
	return []int{}
}

func validSums(nums []int) []int {
	sums := []int{}
	for i, n := range nums {
		for j, m := range nums {
			if i != j {
				sums = append(sums, n+m)
			}
		}
	}
	return sums
}

func contains(nums []int, n int) bool {
	for _, m := range nums {
		if n == m {
			return true
		}
	}
	return false
}

func sum(nums []int) int {
	acc := 0
	for _, n := range nums {
		acc += n
	}
	return acc
}

func min(nums []int) int {
	res := nums[0]
	for _, n := range nums {
		if n < res {
			res = n
		}
	}
	return res
}

func max(nums []int) int {
	res := nums[0]
	for _, n := range nums {
		if n > res {
			res = n
		}
	}
	return res
}
