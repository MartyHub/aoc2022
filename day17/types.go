package main

import (
	"aoc2022"
	"fmt"
	"log"
	"strings"
)

type push rune

func (p push) String() string {
	return fmt.Sprintf("%c", p)
}

const (
	left  push = '<'
	right push = '>'
)

type loop[T any] struct {
	index   int
	pattern []T
}

func newLoop[T any](pattern []T) *loop[T] {
	return &loop[T]{pattern: pattern}
}

func (l *loop[T]) next() T {
	result := l.pattern[l.index%len(l.pattern)]

	l.index++

	return result
}

func (l *loop[T]) String() string {
	return fmt.Sprint(l.pattern[l.index%len(l.pattern)])
}

func parseJet(fileName string) *loop[push] {
	lr := aoc2022.NewLineReader(fileName)

	defer aoc2022.Close(lr)

	if !lr.HasNext() {
		log.Fatalf("Failed to parse jet from %v", fileName)
	}

	return newLoop([]push(lr.Text()))
}

type row byte

func (r row) String() string {
	sb := strings.Builder{}

	for mask := byte(0b1000000); mask != 0; mask >>= 1 {
		if byte(r)&mask == 0 {
			sb.WriteRune('.')
		} else {
			sb.WriteRune('#')
		}
	}

	return sb.String()
}

type rock []row

func (r rock) String() string {
	sb := strings.Builder{}

	for y := len(r) - 1; y >= 0; y-- {
		sb.WriteString(r[y].String())

		if y > 0 {
			sb.WriteRune('\n')
		}
	}

	return sb.String()
}

var rock1 = rock{
	0b0011110,
}

var rock2 = rock{
	0b0001000,
	0b0011100,
	0b0001000,
}

var rock3 = rock{
	0b0011100,
	0b0000100,
	0b0000100,
}

var rock4 = rock{
	0b0010000,
	0b0010000,
	0b0010000,
	0b0010000,
}

var rock5 = rock{
	0b0011000,
	0b0011000,
}

func newRocksLoop() *loop[rock] {
	return newLoop([]rock{rock1, rock2, rock3, rock4, rock5})
}

type Chamber struct {
	data         []row
	cycle        []row
	height       int
	currentRock  rock
	currentRockY int
	jet          *loop[push]
	rocks        *loop[rock]
}

func (c *Chamber) addRock() {
	c.currentRock = aoc2022.Copy(c.rocks.next())
	c.currentRockY = len(c.data) + 3
}

func (c *Chamber) applyJet() {
	switch c.jet.next() {
	case left:
		c.moveLeft()
	case right:
		c.moveRight()
	}
}

func (c *Chamber) Height() int {
	return c.height + len(c.data)
}

func (c *Chamber) Iterate(iteration int, debug bool) int {
	previousCycleStart := 0

	for i := 0; i < iteration; i++ {
		c.addRock()

		for {
			c.applyJet()

			if !c.moveDown() {
				c.rest()

				break
			}
		}

		if debug {
			fmt.Printf("Iteration # %d, height = %s:\n%s\n",
				i+1,
				aoc2022.PrettyFormat(c.Height()),
				c.String(),
			)
		}

		if c.cycle == nil {
			if y := c.findCycle(); y > 0 {
				c.cycle = aoc2022.Copy(c.data[y+1:])
				previousCycleStart = i

				fmt.Printf("Iteration # %d: found cycle of length %s\n",
					i+1,
					aoc2022.PrettyFormat(len(c.cycle)),
				)
			}
		} else if c.hasCycle() {
			cycleIterations := i - previousCycleStart
			cycles := (iteration - 1 - i) / cycleIterations

			fmt.Printf("Iteration # %d: cycle in %d iterations\n",
				i+1,
				cycleIterations,
			)

			i += cycles * cycleIterations
			c.height += cycles * len(c.cycle)
		}
	}

	return c.Height()
}

func (c *Chamber) canMoveDown() bool {
	if c.currentRockY == 0 {
		return false
	}

	for i, rr := range c.currentRock {
		if c.row(c.currentRockY+i-1)&rr != 0 {
			return false
		}
	}

	return true
}

func (c *Chamber) moveDown() bool {
	if !c.canMoveDown() {
		return false
	}

	c.currentRockY--

	return true
}

func (c *Chamber) canMoveLeft() bool {
	for i, rr := range c.currentRock {
		if rr&0b1000000 != 0 {
			return false
		}

		if c.row(c.currentRockY+i)&(rr<<1) != 0 {
			return false
		}
	}

	return true
}

func (c *Chamber) moveLeft() {
	if c.canMoveLeft() {
		for i, rr := range c.currentRock {
			c.currentRock[i] = rr << 1
		}
	}
}

func (c *Chamber) canMoveRight() bool {
	for i, rr := range c.currentRock {
		if rr&0b0000001 != 0 {
			return false
		}

		if c.row(c.currentRockY+i)&(rr>>1) != 0 {
			return false
		}
	}

	return true
}

func (c *Chamber) moveRight() {
	if c.canMoveRight() {
		for i, rr := range c.currentRock {
			c.currentRock[i] = rr >> 1
		}
	}
}

func (c *Chamber) rest() {
	for i, rr := range c.currentRock {
		if len(c.data) > c.currentRockY+i {
			c.data[c.currentRockY+i] |= rr
		} else {
			c.data = append(c.data, rr)
		}
	}
}

func (c *Chamber) row(y int) row {
	if y >= len(c.data) {
		return 0
	}

	return c.data[y]
}

func (c *Chamber) findCycle() int {
	l := len(c.data)
	minCycleLength := 10

	for y := l - 1 - minCycleLength; y > l-y; y-- {
		if c.isCycle(y) {
			return y
		}
	}

	return -1
}

func (c *Chamber) isCycle(y int) bool {
	l := len(c.data)

	for i := l - 1; i > y; i-- {
		if c.data[i] != c.data[y-(l-1-i)] {
			return false
		}
	}

	return true
}

func (c *Chamber) hasCycle() bool {
	j := 0

	for i := len(c.cycle) - 1; i >= 0; i-- {
		if c.data[len(c.data)-1-j] != c.cycle[i] {
			return false
		}

		j++
	}

	return true
}

func (c *Chamber) String() string {
	sb := strings.Builder{}

	for y := len(c.data) - 1; y >= 0; y-- {
		sb.WriteString(c.data[y].String())

		if y > 0 {
			sb.WriteRune('\n')
		}
	}

	return sb.String()
}

func NewChamber(fileName string) *Chamber {
	return &Chamber{
		jet:   parseJet(fileName),
		rocks: newRocksLoop(),
	}
}
