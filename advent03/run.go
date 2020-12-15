package advent03

import (
	"fmt"

	"github.com/igilham/aoc2020go/util"
)

const treeChar = "#"

// Run calculates the answer to the third advent problem
func Run(lines []string, startCol int, rightMove int, downMove int) {
	fmt.Println("input:")
	util.PrintLines(lines)

	count := 0
	for row, col := downMove, startCol+rightMove; row < len(lines); row, col = row+downMove, col+rightMove {
		nextChar := getChar(lines[row], col)
		if nextChar == treeChar {
			// fmt.Printf("found a tree on line %v:%v (%v)\n", row, col, lines[row])
			count++
		}
	}
	fmt.Printf("encountered %v trees\n", count)
}

func getChar(input string, index int) string {
	return string([]byte{input[index%len(input)]})
}
