package advent07

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var isNumberRegexp = regexp.MustCompile("^\\d+$")

// Run runs the seventh problem
func Run(lines []string) {
	bags := parseRules(lines)

	// fmt.Printf("Rules:\n%v", bags)
	// fmt.Printf("total: %v\n", len(bags))

	// scenario 1: you have a shiny gold bag
	have1 := "shiny gold"
	directHolders := bags.whatCanHold(have1)
	fmt.Printf("%v bags can directly hold %v\n", len(directHolders), have1)

	allHolders := bags.allHolders(have1)
	fmt.Printf("%v bags can hold %v\n", len(allHolders), have1)
}

// BagRules is a map of parsed inpot data
type BagRules map[string]Contents

func (r BagRules) String() string {
	var sb strings.Builder
	for k, v := range r {
		contents := []string{}
		for k1, v1 := range v {
			contents = append(contents, fmt.Sprintf("%v %v bags", v1, k1))
		}
		contentsStr := strings.Join(contents, ", ")
		if contentsStr == "" {
			contentsStr = "no other bags."
		}

		sb.WriteString(fmt.Sprintf("%v bags contain %v.\n", k, contentsStr))
	}
	return sb.String()
}

func (r BagRules) allHolders(s string) map[string]bool {
	direct := r.whatCanHold(s)
	holders := merge(direct, r.allHoldersMap(direct))
	return holders
}

func (r BagRules) allHoldersMap(s map[string]bool) map[string]bool {
	holders := make(map[string]bool)
	for k := range s {
		holders = merge(holders, r.allHolders(k))
	}
	return holders
}

func (r BagRules) whatCanHold(s string) map[string]bool {
	holders := make(map[string]bool)
	for k, v := range r {
		if v[s] > 0 {
			holders[k] = true
		}
	}
	return holders
}

// Contents is a convenience alias for what's in the top level map
// of the data model
type Contents map[string]int

func parseRules(lines []string) BagRules {
	bags := make(BagRules)
	for _, line := range lines {
		topParts := strings.Split(line, " bags contain ")
		s := topParts[0]
		if bags[s] == nil {
			bags[s] = parseContents(topParts[1])
		}
	}
	return bags
}

func parseContents(s string) Contents {
	contents := make(Contents)

	containedStr := strings.TrimSuffix(s, ".")
	contained := strings.Split(containedStr, ", ")
	for _, b := range contained {
		words := strings.Split(b, " ")
		count, err := getNum(words[0])
		if err == nil {
			containedBag := strings.Join(words[1:len(words)-1], " ")
			contents[containedBag] = count
		}
	}

	return contents
}

func getNum(s string) (int, error) {
	m := isNumberRegexp.FindString(s)
	if m == "" {
		return 0, errors.New("no number found")
	}
	i, err2 := strconv.Atoi(m)
	if err2 != nil {
		return 0, err2
	}
	return i, nil
}

func merge(ms ...map[string]bool) map[string]bool {
	res := make(map[string]bool)
	for _, m := range ms {
		for k, v := range m {
			res[k] = v
		}
	}
	return res
}
