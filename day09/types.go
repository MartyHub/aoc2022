package main

import (
	"aoc2022"
	"fmt"
	"log"
	"strconv"
)

type Direction rune

const (
	Up    Direction = 'U'
	Right Direction = 'R'
	Down  Direction = 'D'
	Left  Direction = 'L'
)

func (d Direction) String() string {
	switch d {
	case Up:
		return "Up"
	case Right:
		return "Right"
	case Down:
		return "Down"
	case Left:
		return "Left"
	}

	log.Fatalf("Unknown direction %#v", d)

	return ""
}

type Motion struct {
	Direction Direction
	Steps     int
}

func ParseMotion(s string) Motion {
	return Motion{
		Direction: Direction(s[0]),
		Steps:     aoc2022.Must(strconv.Atoi(s[2:])),
	}
}

func (m Motion) String() string {
	return fmt.Sprintf("%s %d", m.Direction, m.Steps)
}

type Position struct {
	X int
	Y int
}

func (p *Position) CheckX(other *Position) Direction {
	if p.X < other.X {
		return Right
	}

	if p.X > other.X {
		return Left
	}

	return 0
}

func (p *Position) CheckY(other *Position) Direction {
	if p.Y < other.Y {
		return Up
	}

	if p.Y > other.Y {
		return Down
	}

	return 0
}

func (p *Position) Move(d Direction) {
	switch d {
	case Up:
		p.Y += 1
	case Right:
		p.X += 1
	case Down:
		p.Y -= 1
	case Left:
		p.X -= 1
	}
}

func (p *Position) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

func (p *Position) Touch(other *Position) bool {
	return aoc2022.Abs(p.X-other.X) <= 1 && aoc2022.Abs(p.Y-other.Y) <= 1
}

type Rope struct {
	knots         []*Position
	tailPositions map[Position]bool
}

func NewRope(size int) *Rope {
	result := &Rope{
		knots:         make([]*Position, size),
		tailPositions: make(map[Position]bool),
	}

	for i := 0; i < size; i++ {
		result.knots[i] = &Position{}
	}

	result.tailPositions[*result.Tail()] = true

	return result
}

func (r *Rope) Head() *Position {
	return r.knots[0]
}

func (r *Rope) Move(m Motion) {
	for i := 0; i < m.Steps; i++ {
		r.Head().Move(m.Direction)

		for i := 0; i < len(r.knots)-1; i++ {
			if !r.knots[i].Touch(r.knots[i+1]) {
				r.knots[i+1].Move(r.knots[i+1].CheckX(r.knots[i]))
				r.knots[i+1].Move(r.knots[i+1].CheckY(r.knots[i]))
			}
		}

		r.tailPositions[*r.Tail()] = true
	}
}

func (r *Rope) String() string {
	return fmt.Sprintf("Head %v -> Tail %v", r.Head(), r.Tail())
}

func (r *Rope) Tail() *Position {
	return r.knots[len(r.knots)-1]
}
