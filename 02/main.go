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
	nbuf = 1024 * 4
)

type record struct {
	min      int
	max      int
	testChar string
	password string
}

func (r *record) String() string {
	return fmt.Sprintf("min: %v, max: %v, testChar: %v, password: %v", r.min, r.max, r.testChar, r.password)
}

// must be between min and max testChar instances in password
func (r *record) isValid1() bool {
	occurances := strings.Count(r.password, r.testChar)
	return occurances >= r.min && occurances <= r.max
}

// testChar must be at one of min or max indexes (1-indexed) only
func (r *record) isValid2() bool {
	n := len(r.password)
	// fmt.Printf("len: %v\n", n)
	return n > r.min && n >= r.max &&
		(r.password[r.min-1] == r.testChar[0] && r.password[r.max-1] != r.testChar[0]) ||
		(r.password[r.min-1] != r.testChar[0] && r.password[r.max-1] == r.testChar[0])

}

func main() {
	reader := os.Stdin
	var writer strings.Builder
	if error := Cat(reader, &writer); error != nil {
		log.Fatalf("failed to read input: %v\n", error)
	}

	input := writer.String()
	lines := strings.Split(input, "\n")
	validCount1 := 0
	validCount2 := 0

	for _, line := range lines {
		if line != "" {
			r, err := parseLine(line)
			if err != nil {
				log.Fatalf("%v\n", err)
			}
			// fmt.Printf("got record: %v\n", r)
			if r.isValid1() {
				validCount1++
			}
			if r.isValid2() {
				validCount2++
			}
		}
	}
	fmt.Printf("valid lines (part 1): %v\n", validCount1)
	fmt.Printf("valid lines (part 2): %v\n", validCount2)
}

func parseLine(line string) (*record, error) {
	// format: "1-3 a: abcdea"
	fields := strings.Split(line, " ")
	min, max, e1 := parseMinMax(fields[0])
	if e1 != nil {
		return nil, e1
	}
	return &record{
		min:      min,
		max:      max,
		testChar: strings.Trim(fields[1], ":"),
		password: fields[2],
	}, nil
}

func parseMinMax(text string) (int, int, error) {
	fields := strings.Split(text, "-")
	smin := fields[0]
	smax := fields[1]
	min, e1 := strconv.Atoi(smin)
	if e1 != nil {
		return -1, -1, e1
	}
	max, e2 := strconv.Atoi(smax)
	if e2 != nil {
		return -1, -1, e2
	}
	return min, max, nil
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
