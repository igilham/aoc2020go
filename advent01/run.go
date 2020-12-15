package advent01

import (
	"errors"
	"fmt"
	"log"
	"strconv"
)

const (
	// buffer size in bytes
	target int = 2020
)

// Run calculates the answer to the first advent problem
func Run(lines []string) {
	nums, err := toNumbers(lines)
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

func toNumbers(records []string) ([]int, error) {
	var nums []int
	for _, s := range records {
		if s == "" {
			break
		}

		n, err := strconv.Atoi(s)
		if err != nil {
			return nil, errors.New("Failed to convert to numbers")
		}
		nums = append(nums, n)
	}
	return nums, nil
}
