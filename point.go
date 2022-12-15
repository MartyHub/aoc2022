package aoc2022

import "fmt"

type Point struct {
	X, Y int
}

func (p Point) Distance(other Point) int {
	return Abs(p.X-other.X) + Abs(p.Y-other.Y)
}

func (p Point) Equals(x, y int) bool {
	return p.X == x && p.Y == y
}

func (p Point) Less(other Point) bool {
	if p.Y < other.Y {
		return true
	} else if p.Y > other.Y {
		return false
	}

	return p.X < other.X
}

func (p Point) Max(other Point) Point {
	return Point{X: Max(p.X, other.X), Y: Max(p.Y, other.Y)}
}

func (p Point) Min(other Point) Point {
	return Point{X: Min(p.X, other.X), Y: Min(p.Y, other.Y)}
}

func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}
