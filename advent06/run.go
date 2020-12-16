package advent06

import "fmt"

// Run runs the sixth problem
func Run(groupedLines [][]string) {
	// each line in a group is a person
	// each char in a line is a question they answer "yes" to
	total := 0
	for i, group := range groupedLines {
		n := countAnswers(group)
		total = total + n
		fmt.Printf("group %v: %v common answers\n", i, n)
	}
	fmt.Printf("total: %v\n", total)
}

func countAnswers(group []string) int {
	m := make(map[rune]bool, 26)
	for _, line := range group {
		for _, b := range line {
			m[b] = true
		}
	}
	return len(m)
}
