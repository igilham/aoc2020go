package advent04

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	colourCodeRegexp = regexp.MustCompile("^#[0-9a-f]{6}$")
	pidRegexp        = regexp.MustCompile("^\\d{9}$")
)

// Passport has all the expected fields, countryID is ignored
type Passport struct {
	birthYear      string
	issueYear      string
	expirationYear string
	height         string
	hairColour     string
	eyeColour      string
	passportID     string
	countryID      string
}

// String converts to a string
func (p *Passport) String() string {
	return fmt.Sprintf("byr:%v iyr:%v eyr:%v hgt:%v hcl:%v ecl:%v pid:%v cid:%v", p.birthYear, p.issueYear, p.expirationYear, p.height, p.hairColour, p.eyeColour, p.passportID, p.countryID)
}

func (p *Passport) hasAllFields() bool {
	// countryID is ignored
	return p.birthYear != "" && p.issueYear != "" && p.expirationYear != "" && p.height != "" && p.hairColour != "" && p.eyeColour != "" && p.passportID != ""
}

func (p *Passport) isValid() bool {
	return p.hasAllFields() &&
		validateYearBetween(p.birthYear, 1920, 2002) &&
		validateYearBetween(p.issueYear, 2010, 2020) &&
		validateYearBetween(p.expirationYear, 2020, 2030) &&
		validateHeight(p.height) &&
		validateHairColour(p.hairColour) &&
		validateEyeColour(p.eyeColour) &&
		validatePassportID(p.passportID)
}

// Run runs the fourth problem
func Run(input string) {
	// blank lines separate passports
	inputBlocks := strings.Split(input, "\n\n")
	passports := []*Passport{}
	for _, block := range inputBlocks {
		passports = append(passports, parsePassport(block))
	}

	validSimpleCount := 0
	validCount := 0
	for _, p := range passports {
		if p.hasAllFields() {
			validSimpleCount++
		}
		if p.isValid() {
			validCount++
			fmt.Printf("valid: %v\n", p)
		}
	}
	fmt.Printf("passports with all fields: %v\n", validSimpleCount)
	fmt.Printf("valid passports: %v\n", validCount)
}

func parsePassport(input string) *Passport {
	passport := &Passport{}
	fields := strings.Fields(input)
	for _, f := range fields {
		kvp := strings.Split(f, ":")
		k, v := kvp[0], kvp[1]
		switch k {
		case "byr":
			passport.birthYear = v
		case "iyr":
			passport.issueYear = v
		case "eyr":
			passport.expirationYear = v
		case "hgt":
			passport.height = v
		case "hcl":
			passport.hairColour = v
		case "ecl":
			passport.eyeColour = v
		case "pid":
			passport.passportID = v
		case "cid":
			passport.countryID = v
		default:
			break
		}
	}

	return passport
}

func extractNumber(height string) (int, error) {
	re := regexp.MustCompile(`[-]?\d+`)
	sub := re.FindString(height)
	return strconv.Atoi(sub)
}

func isBetween(i int, min int, max int) bool {
	return i >= min && i <= max
}

func validateYearBetween(s string, min int, max int) bool {
	year, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	return isBetween(year, min, max)
}

func validateHeight(s string) bool {
	height, err := extractNumber(s)
	if err != nil {
		return false
	}

	l := len(s)
	if l < 3 {
		return false
	}

	unit := s[l-2:]
	switch unit {
	case "cm":
		return isBetween(height, 150, 193)
	case "in":
		return isBetween(height, 59, 76)
	default:
		return false
	}
}

func validateHairColour(s string) bool {
	return colourCodeRegexp.MatchString(s)
}

func validateEyeColour(s string) bool {
	return s == "amb" ||
		s == "blu" ||
		s == "brn" ||
		s == "gry" ||
		s == "grn" ||
		s == "hzl" ||
		s == "oth"
}

func validatePassportID(s string) bool {
	return pidRegexp.MatchString(s)
}
