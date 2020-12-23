package util

import (
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const (
	nbuf = 1024 * 4
)

// PrintLines prints lines in an easy-to-read format
func PrintLines(rows []string) {
	for _, r := range rows {
		fmt.Printf("\"%v\"\n", r)
	}
}

// StringToGroupedLines groups parses the input into a list of groups where each group is a list of lines
func StringToGroupedLines(input string, dropEmptyLine bool) [][]string {
	groupedLines := [][]string{}
	blocks := StringToBlocks(input)
	for _, block := range blocks {
		groupedLines = append(groupedLines, StringToLines(block, dropEmptyLine))
	}
	return groupedLines
}

// StringToBlocks splits a string into groups separated by a blank line
func StringToBlocks(input string) []string {
	return strings.Split(input, "\n\n")
}

// StringToLines splits a string into lines, optionally dropping a final empty line
func StringToLines(input string, dropEmptyLine bool) []string {
	rows := strings.Split(input, "\n")
	if dropEmptyLine {
		maxIndex := len(rows) - 1
		if rows[maxIndex] == "" {
			return rows[:maxIndex]
		}
	}
	return rows
}

// ReadToLines reads to an array of lines, optionally dropping a final empty line
func ReadToLines(reader io.Reader, dropEmptyLine bool) ([]string, error) {
	input, err := ReadToString(reader)
	if err != nil {
		return nil, err
	}

	return StringToLines(input, dropEmptyLine), nil
}

// ReadToString reads to a string
func ReadToString(reader io.Reader) (string, error) {
	var writer strings.Builder
	if error := Cat(reader, &writer); error != nil {
		return "", error
	}

	return writer.String(), nil
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

// ToNumbers converst a slice of strings to a slice of ints
func ToNumbers(records []string) ([]int, error) {
	var nums []int
	for _, s := range records {
		if s == "" {
			break
		}

		n, err := strconv.Atoi(s)
		if err != nil {
			return nil, errors.New("Failed to convert to numbers")
		}
		nums = append(nums, n)
	}
	return nums, nil
}
