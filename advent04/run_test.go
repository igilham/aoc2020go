package advent04

import "testing"

func TestValidateHeight(t *testing.T) {
	passData := []string{
		"150cm",
		"151cm",
		"180cm",
		"192cm",
		"193cm",
		"59in",
		"60in",
		"70in",
		"75in",
		"76in",
	}
	for _, s := range passData {
		if !validateHeight(s) {
			t.Errorf("should be valid: %v\n", s)
		}
	}

	failData := []string{
		"149cm",
		"151",
		"58in",
		"58",
		"194cm",
		"194",
		"77in",
		"77",
	}
	for _, s := range failData {
		if validateHeight(s) {
			t.Errorf("should not be valid: %v\n", s)
		}
	}
}

func TestNegativeValidity(t *testing.T) {
	data := []string{
		"eyr:1972 cid:100\nhcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926",
		"iyr:2019\nhcl:#602927 eyr:1967 hgt:170cm\necl:grn pid:012533040 byr:1946",
		"hcl:dab227 iyr:2012\necl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277",
		"hgt:59cm ecl:zzz\neyr:2038 hcl:74454a iyr:2023\npid:3556412378 byr:2007",
	}
	for _, input := range data {
		p := parsePassport(input)
		if p.isValid() {
			t.Errorf("should not be valid: %v", input)
		}
	}
}

func TestPositiveValidity(t *testing.T) {
	data := []string{
		"pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980\nhcl:#623a2f",
		"eyr:2029 ecl:blu cid:129 byr:1989\niyr:2014 pid:896056539 hcl:#a97842 hgt:165cm",
		"hcl:#888785\nhgt:164cm byr:2001 iyr:2015 cid:88\npid:545766238 ecl:hzl\neyr:2022",
		"iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719",
	}
	for _, input := range data {
		p := parsePassport(input)
		if !p.isValid() {
			t.Errorf("should be valid: %v", input)
		}
	}
}
