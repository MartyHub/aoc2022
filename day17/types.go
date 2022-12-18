package main

import (
	"aoc2022"
	"fmt"
	"log"
	"strings"
)

type Push rune

func (p Push) String() string {
	return fmt.Sprintf("%c", p)
}

const (
	Left  Push = '<'
	Right Push = '>'
	width      = 7
)

type Jet[T any] struct {
	index   int
	pattern []T
}

func (j *Jet[T]) Next() T {
	result := j.pattern[j.index%len(j.pattern)]

	j.index++

	return result
}

func (j *Jet[T]) String() string {
	return fmt.Sprint(j.pattern[j.index%len(j.pattern)])
}

func ParseJet(fileName string) *Jet[Push] {
	lr := aoc2022.NewLineReader(fileName)

	defer aoc2022.Close(lr)

	if !lr.HasNext() {
		log.Fatalf("Failed to parse jet from %v", fileName)
	}

	return &Jet[Push]{pattern: []Push(lr.Text())}
}

type Rock [][]rune

func (r Rock) String() string {
	sb := strings.Builder{}

	for row := len(r) - 1; row >= 0; row-- {
		for col := 0; col < len(r[row]); col++ {
			sb.WriteRune(r[row][col])
		}

		sb.WriteRune('\n')
	}

	return sb.String()
}

var Rock1 = Rock{
	{'.', '.', '@', '@', '@', '@', '.'},
}

var Rock2 = Rock{
	{'.', '.', '.', '@', '.', '.', '.'},
	{'.', '.', '@', '@', '@', '.', '.'},
	{'.', '.', '.', '@', '.', '.', '.'},
}

var Rock3 = Rock{
	{'.', '.', '@', '@', '@', '.', '.'},
	{'.', '.', '.', '.', '@', '.', '.'},
	{'.', '.', '.', '.', '@', '.', '.'},
}

var Rock4 = Rock{
	{'.', '.', '@', '.', '.', '.', '.'},
	{'.', '.', '@', '.', '.', '.', '.'},
	{'.', '.', '@', '.', '.', '.', '.'},
	{'.', '.', '@', '.', '.', '.', '.'},
}

var Rock5 = Rock{
	{'.', '.', '@', '@', '.', '.', '.'},
	{'.', '.', '@', '@', '.', '.', '.'},
}

func NewRockJet() *Jet[Rock] {
	return &Jet[Rock]{pattern: []Rock{Rock1, Rock2, Rock3, Rock4, Rock5}}
}

type Chamber struct {
	data              [][]rune
	height            int
	iteration         int
	currentRockHeight int
	currentRockY      int
	jet               *Jet[Push]
	rocks             *Jet[Rock]
}

func (c *Chamber) addRock() {
	empty := aoc2022.Max(0, len(c.data)-1-c.highestRock())

	if empty > 3 {
		c.data = c.data[:len(c.data)-(empty-3)]
	} else {
		for i := empty; i < 3; i++ {
			c.data = append(c.data, []rune{'.', '.', '.', '.', '.', '.', '.'})
		}
	}

	rock := c.rocks.Next()

	c.currentRockHeight = len(rock)
	c.currentRockY = len(c.data)

	for i := 0; i < c.currentRockHeight; i++ {
		c.data = append(c.data, []rune(string(rock[i])))
	}
}

func (c *Chamber) applyJet() {
	push := c.jet.Next()

	switch push {
	case Left:
		c.moveLeft()
	case Right:
		c.moveRight()
	}
}

func (c *Chamber) highestRock() int {
	for y := len(c.data) - 1; y >= 0; y-- {
		if c.hasRock(y) {
			return y
		}
	}

	return 0
}

func (c *Chamber) hasRock(row int) bool {
	for _, r := range c.data[row] {
		if r == '#' {
			return true
		}
	}

	return false
}

func (c *Chamber) iterate() {
	c.addRock()

	for {
		c.applyJet()

		if !c.moveDown() {
			c.rest()
			c.compress()

			break
		}
	}

	c.iteration++
}

func (c *Chamber) compress() {
	for i := c.currentRockHeight; i >= 0; i-- {
		y := c.currentRockY + i

		if c.rowFull(y) {
			c.data = c.data[y+1:]
			c.height += y + 1

			break
		}
	}
}

func (c *Chamber) rowFull(y int) bool {
	for x := 0; x < width; x++ {
		if c.data[y][x] != '#' {
			return false
		}
	}

	return true
}

func (c *Chamber) canMoveDown() bool {
	for y := 0; y < c.currentRockHeight; y++ {
		if !c.canRowMoveDown(c.currentRockY + y) {
			return false
		}
	}

	return true
}

func (c *Chamber) canRowMoveDown(y int) bool {
	if y == 0 {
		return false
	}

	for x := 0; x < width; x++ {
		if c.data[y][x] == '@' && c.data[y-1][x] == '#' {
			return false
		}
	}

	return true
}

func (c *Chamber) moveDown() bool {
	if !c.canMoveDown() {
		return false
	}

	for y := 0; y < c.currentRockHeight; y++ {
		c.moveRowDown(c.currentRockY + y)
	}

	c.currentRockY--

	return true
}

func (c *Chamber) moveRowDown(y int) {
	for x := 0; x < width; x++ {
		if c.data[y][x] == '@' {
			c.data[y-1][x] = '@'
			c.data[y][x] = '.'
		}
	}
}

func (c *Chamber) canMoveLeft() bool {
	for y := 0; y < c.currentRockHeight; y++ {
		if !c.canRowMoveLeft(c.currentRockY + y) {
			return false
		}
	}

	return true
}

func (c *Chamber) canRowMoveLeft(y int) bool {
	for x := width - 1; x > 0; x-- {
		if c.data[y][x] == '@' && c.data[y][x-1] == '#' {
			return false
		}
	}

	return c.data[y][0] != '@'
}

func (c *Chamber) moveLeft() {
	if c.canMoveLeft() {
		for y := 0; y < c.currentRockHeight; y++ {
			c.moveRowLeft(c.currentRockY + y)
		}
	}
}

func (c *Chamber) moveRowLeft(y int) {
	for x := 1; x < width; x++ {
		if c.data[y][x] == '@' {
			c.data[y][x] = '.'
			c.data[y][x-1] = '@'
		}
	}
}

func (c *Chamber) canMoveRight() bool {
	for y := 0; y < c.currentRockHeight; y++ {
		if !c.canRowMoveRight(c.currentRockY + y) {
			return false
		}
	}

	return true
}

func (c *Chamber) canRowMoveRight(y int) bool {
	for x := 0; x < width-1; x++ {
		if c.data[y][x] == '@' && c.data[y][x+1] == '#' {
			return false
		}
	}

	return c.data[y][width-1] != '@'
}

func (c *Chamber) moveRight() {
	if c.canMoveRight() {
		for y := 0; y < c.currentRockHeight; y++ {
			c.moveRowRight(c.currentRockY + y)
		}
	}
}

func (c *Chamber) moveRowRight(y int) {
	for x := 5; x >= 0; x-- {
		if c.data[y][x] == '@' {
			c.data[y][x] = '.'
			c.data[y][x+1] = '@'
		}
	}
}

func (c *Chamber) rest() {
	for y := 0; y < c.currentRockHeight; y++ {
		c.restRow(c.currentRockY + y)
	}
}

func (c *Chamber) restRow(y int) {
	for x := 0; x < width; x++ {
		if c.data[y][x] == '@' {
			c.data[y][x] = '#'
		}
	}
}

func (c *Chamber) String() string {
	sb := strings.Builder{}

	for row := len(c.data) - 1; row >= 0; row-- {
		for col := 0; col < len(c.data[row]); col++ {
			sb.WriteRune(c.data[row][col])
		}

		sb.WriteRune('\n')
	}

	sb.WriteString(fmt.Sprintf("Current Rock Height: %v\n", c.currentRockHeight))
	sb.WriteString(fmt.Sprintf("Current Rock Y:      %v\n", c.currentRockY))

	return sb.String()
}

func NewChamber(fileName string) *Chamber {
	return &Chamber{
		data:  [][]rune{},
		jet:   ParseJet(fileName),
		rocks: NewRockJet(),
	}
}
