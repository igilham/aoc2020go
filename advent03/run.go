package advent03

import (
	"fmt"
)

const treeChar = "#"

// Run calculates the answer to the third advent problem
func Run(lines []string) {
	// fmt.Println("input:")
	// util.PrintLines(lines)

	count1 := countTrees(lines, 0, 1, 1)
	count2 := countTrees(lines, 0, 3, 1)
	count3 := countTrees(lines, 0, 5, 1)
	count4 := countTrees(lines, 0, 7, 1)
	count5 := countTrees(lines, 0, 1, 2)
	fmt.Printf("move(1,1) encountered %v trees\n", count1)
	fmt.Printf("move(3,1) encountered %v trees\n", count2)
	fmt.Printf("move(5,1) encountered %v trees\n", count3)
	fmt.Printf("move(7,1) encountered %v trees\n", count4)
	fmt.Printf("move(1,2) encountered %v trees\n", count5)
	fmt.Printf("the product of the above is %v trees\n", count1*count2*count3*count4*count5)
}

func countTrees(lines []string, startCol int, rightMove int, downMove int) int {
	count := 0
	for row, col := downMove, startCol+rightMove; row < len(lines); row, col = row+downMove, col+rightMove {
		nextChar := getChar(lines[row], col)
		if nextChar == treeChar {
			// fmt.Printf("found a tree on line %v:%v (%v)\n", row, col, lines[row])
			count++
		}
	}
	return count
}

func getChar(input string, index int) string {
	return string([]byte{input[index%len(input)]})
}
