package advent06

import "fmt"

// Run runs the sixth problem
func Run(groupedLines [][]string) {
	// each line in a group is a person
	// each char in a line is a question they answer "yes" to
	totalAny := 0
	totalEvery := 0
	for i, group := range groupedLines {
		n, m := countAnswers(group)
		totalAny = totalAny + n
		totalEvery = totalEvery + m
		fmt.Printf("group %v: anyone=%v,  everyone=%v\n", i, n, m)
	}
	fmt.Printf("total for which anyone answered 'yes': %v\n", totalAny)
	fmt.Printf("total for which everyone answered 'yes': %v\n", totalEvery)
}

func loadAnswers(group []string) map[rune]int {
	m := make(map[rune]int, 26)
	for _, line := range group {
		for _, b := range line {
			n := m[b]
			if n == 0 {
				m[b] = 1
			} else {
				m[b] = n + 1
			}
		}
	}
	return m
}

func countAnswers(group []string) (anyone int, everyone int) {
	m := loadAnswers(group)

	nPeople := len(group)
	anyone = len(m)
	everyone = 0
	for _, n := range m {
		if n == nPeople {
			everyone++
		}
	}
	return
}
