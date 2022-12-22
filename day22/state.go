package main

import (
	"fmt"
	"log"
)

type Direction int

const (
	right Direction = iota
	down
	left
	up
)

func (d Direction) String() string {
	switch d {
	case up:
		return "up"
	case right:
		return "right"
	case down:
		return "down"
	case left:
		return "left"
	}

	return "invalid"
}

type Point struct {
	x, y int
}

func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.x, p.y)
}

func (p Point) Move(d Direction) Point {
	switch d {
	case up:
		return Point{p.x, p.y - 1}
	case right:
		return Point{p.x + 1, p.y}
	case down:
		return Point{p.x, p.y + 1}
	case left:
		return Point{p.x - 1, p.y}
	}

	log.Fatalf("Invalid direction: %v", d)

	return Point{}
}

type State struct {
	pos  Point
	face Direction
}

func NewState(pos Point) State {
	return State{pos, right}
}

func (s State) TurnLeft() State {
	result := s

	switch s.face {
	case up:
		result.face = left
	case right:
		result.face = up
	case down:
		result.face = right
	case left:
		result.face = down
	}

	return result
}

func (s State) TurnRight() State {
	result := s

	switch s.face {
	case up:
		result.face = right
	case right:
		result.face = down
	case down:
		result.face = left
	case left:
		result.face = up
	}

	return result
}

func (s State) Password() int {
	return 1000*(s.pos.y+1) + 4*(s.pos.x+1) + int(s.face)
}

func (s State) String() string {
	return fmt.Sprintf("At %v facing %v", s.pos, s.face)
}
