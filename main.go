package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/igilham/aoc2020go/advent01"
	"github.com/igilham/aoc2020go/advent02"
	"github.com/igilham/aoc2020go/advent03"
	"github.com/igilham/aoc2020go/advent04"
	"github.com/igilham/aoc2020go/advent05"
	"github.com/igilham/aoc2020go/advent06"
	"github.com/igilham/aoc2020go/advent07"
	"github.com/igilham/aoc2020go/advent08"
	"github.com/igilham/aoc2020go/advent09"
	"github.com/igilham/aoc2020go/util"
)

const nbuf = 1024 * 4

var advent = flag.Int("a", 1, "which advent to run")

func main() {
	flag.Parse()
	input, err := util.ReadToString(os.Stdin)
	if err != nil {
		log.Fatalf("%v", err)
	}

	fmt.Printf("running problem %v\n", *advent)

	switch *advent {
	case 1:
		advent01.Run(util.StringToLines(input, true))
	case 2:
		advent02.Run(util.StringToLines(input, true))
	case 3:
		advent03.Run(util.StringToLines(input, true))
	case 4:
		advent04.Run(util.StringToBlocks(input))
	case 5:
		advent05.Run(util.StringToLines(input, true))
	case 6:
		advent06.Run(util.StringToGroupedLines(input, true))
	case 7:
		advent07.Run(util.StringToLines(input, true))
	case 8:
		advent08.Run(util.StringToLines(input, true))
	case 9:
		advent09.Run(util.StringToLines(input, true))
	default:
		fmt.Println("no problem specified")
	}
}
