package advent07

import (
	"fmt"
	"regexp"
	"strings"
	"text/scanner"

	"github.com/igilham/aoc2020go/util"
)

var isNumberRegexp = regexp.MustCompile("^\\d+$")

// Run runs the seventh problem
func Run(lines []string) {
	bags := parseBagTypes(lines)
	util.PrintLines(bags)
	fmt.Printf("total: %v", len(bags))
}

// pull out all types of bags, including duplicates
// TODO: remove duplicates
// TODO: encode container rules and numbers in response
func parseBagTypes(lines []string) []string {
	bags := []string{}
	for _, line := range lines {
		var s scanner.Scanner
		s.Init(strings.NewReader(line))
		w1 := ""
		for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
			t := s.TokenText()
			switch t {
			case "no":
				fallthrough
			case "other":
				fallthrough
			case "bag":
				fallthrough
			case "bags":
				fallthrough
			case "contain":
				fallthrough
			case ",":
				continue
			default:
				isNumber := isNum(t)
				if isNumber {
					continue
				}
				if w1 == "" {
					w1 = t
					continue
				} else {
					bags = append(bags, fmt.Sprintf("%s %s", w1, t))
					w1 = ""
				}
			}
		}
	}
	return bags
}

func isNum(s string) bool {
	return isNumberRegexp.MatchString(s)
}
