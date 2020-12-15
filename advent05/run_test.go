package advent05

import "testing"

type Case struct {
	s      string
	row    int
	column int
	id     int
}

var data = []Case{
	{
		s:      "FFFFFFFLLL",
		row:    0,
		column: 0,
		id:     0,
	},
	{
		s:      "FFFFFFBLLL",
		row:    1,
		column: 0,
		id:     8,
	},
	{
		s:      "FBFBBFFRLR",
		row:    44,
		column: 5,
		id:     357,
	},
	{
		s:      "BFFFBBFRRR",
		row:    70,
		column: 7,
		id:     567,
	},
	{
		s:      "FFFBBBFRRR",
		row:    14,
		column: 7,
		id:     119,
	},
	{
		s:      "BBFFBBFRLL",
		row:    102,
		column: 4,
		id:     820,
	},
}

func TestFindSeat(t *testing.T) {
	for _, c := range data {
		row, col := findSeat(c.s)
		if row != c.row {
			t.Errorf("wrong row %v:%v, got %v", c.s, c.row, row)
		}
		if col != c.column {
			t.Errorf("wrong column %v:%v, got %v", c.s, c.column, col)
		}
	}
}

func TestSeatID(t *testing.T) {
	for _, c := range data {
		result := seatID(c.row, c.column)
		if result != c.id {
			t.Errorf("wrong seat ID: %v:%v=%v, got %v", c.row, c.column, c.id, result)
		}
	}
}
