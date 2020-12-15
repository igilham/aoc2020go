package advent05

import (
	"fmt"
	"sort"
)

// Run runs the fifth problem
func Run(lines []string) {
	ids := []int{}
	for _, s := range lines {
		row, col := findSeat(s)
		ids = append(ids, seatID(row, col))
	}
	max := maxInSlice(ids)
	fmt.Printf("max: %v\n", max)
	mySeat := findMySeat(ids)
	fmt.Printf("my seat is %v\n", mySeat)
}

func findSeat(address string) (int, int) {
	var row uint = 0
	var col uint = 0

	for _, c := range address {
		switch c {
		case 'F':
			row = row << 1
		case 'B':
			row = row<<1 + 1
		case 'L':
			col = col << 1
		case 'R':
			col = col<<1 + 1
		default:
			break
		}
	}
	return int(row), int(col)
}

func seatID(row int, col int) int {
	return row*8 + col
}

func findMySeat(ids []int) int {
	sort.Ints(ids)
	for i := 1; i < maxInSlice(ids); i++ {
		if !contains(ids, i) && contains(ids, i-1) && contains(ids, i+1) {
			return i
		}
	}
	return -1
}

func maxInSlice(a []int) int {
	m := 0
	for i, e := range a {
		if i == 0 || e > m {
			m = e
		}
	}
	return m
}

func contains(a []int, n int) bool {
	i := sort.SearchInts(a, n)
	if i < len(a) && a[i] == n {
		return true
	}
	return false
}
