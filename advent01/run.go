package advent01

import (
	"fmt"
	"log"

	"github.com/igilham/aoc2020go/util"
)

const (
	// buffer size in bytes
	target int = 2020
)

// Run calculates the answer to the first advent problem
func Run(lines []string) {
	nums, err := util.ToNumbers(lines)
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	a, b := find2Nums(nums)
	fmt.Printf("2 numbers: a = %v, b = %v, product = %v\n", a, b, a*b)
	x, y, z := find3Nums(nums)
	fmt.Printf("3 numbers: x = %v, y = %v, z = %v, product = %v\n", x, y, z, x*y*z)
}

func find2Nums(nums []int) (int, int) {
	for i, n := range nums {
		for j, m := range nums {
			if i != j {
				if n+m == target {
					return n, m
				}
			}
		}
	}
	return -1, -1
}

func find3Nums(nums []int) (int, int, int) {
	for i, n := range nums {
		for j, m := range nums {
			if i != j {
				for l, o := range nums {
					if i != l && j != l {
						if n+m+o == target {
							return n, m, o
						}
					}
				}
			}
		}
	}
	return -1, -1, -1
}
