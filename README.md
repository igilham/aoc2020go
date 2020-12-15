# Advent of Code 2020 in Go

An implementation of the [Advent of Code](https://adventofcode.com/2020) for 2020.

## Usage

There is an application in the root directory for running the scenarios. The `-a` flag is used to select a scenario in the range of `1` to `12`. The application reads from standard input for test data, copies of which are provided in subdirectories.

From the root you can run a scenario using `go run main.go -a SCENARIO <DIRECTORY/input.txt`. For example, to run the third scenario, use `go run main.go -a 3 <advent03/input.text`.
