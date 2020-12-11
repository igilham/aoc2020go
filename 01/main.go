package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	// buffer size in bytes
	nbuf       = 1024 * 4
	target int = 2020
)

func main() {
	reader := os.Stdin
	var writer strings.Builder

	if error := Cat(reader, &writer); error != nil {
		log.Fatalf("failed to read input: %v\n", error)
	}
	input := writer.String()
	strRecords := strings.Split(input, "\n")
	nums, err := toNumbers(strRecords)
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

// Cat reads a single file and writes to os.Stdout.
func Cat(reader io.Reader, writer io.Writer) error {
	var buf [nbuf]byte
	for {
		nr, er := reader.Read(buf[:])
		switch {
		case nr < 0:
			return errors.New("read error: " + er.Error())
		case nr == 0: // EOF
			return nil
		case nr > 0:
			if nw, ew := writer.Write(buf[0:nr]); nw != nr {
				return errors.New("write error: " + ew.Error())
			}
		}
	}
}
