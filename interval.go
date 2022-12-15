package aoc2022

import "fmt"

type Interval struct {
	From, To int
}

func (i Interval) Contains(x int) bool {
	return i.From <= x && x < i.To
}

func (i Interval) Empty() bool {
	return i.From == i.To
}

func (i Interval) String() string {
	return fmt.Sprintf("[%v, %v[", i.From, i.To)
}
