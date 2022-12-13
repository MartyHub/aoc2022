package main

import (
	"aoc2022"
	"fmt"
	"math"
	"strings"
)

type Map struct {
	heights [][]int
	paths   *paths
	Start   Position
	Target  Position
}

func (m Map) Height(p Position) int {
	if p.y < 0 || p.y >= len(m.heights) ||
		p.x < 0 || p.x >= len(m.heights[0]) {
		return math.MaxInt
	}

	return m.heights[p.y][p.x]
}

func (m Map) MovesUp(p Position) []Position {
	result := make([]Position, 0, 4)

	for _, s := range []Position{
		{x: p.x, y: p.y - 1},
		{x: p.x + 1, y: p.y},
		{x: p.x, y: p.y + 1},
		{x: p.x - 1, y: p.y},
	} {
		diff := m.Height(s) - m.Height(p)

		if diff <= 1 {
			result = append(result, s)
		}
	}

	return result
}

func (m Map) MovesDown(p Position) []Position {
	result := make([]Position, 0, 4)

	for _, s := range []Position{
		{x: p.x, y: p.y - 1},
		{x: p.x + 1, y: p.y},
		{x: p.x, y: p.y + 1},
		{x: p.x - 1, y: p.y},
	} {
		if m.Height(s) != math.MaxInt && m.Height(s)-m.Height(p) >= -1 {
			result = append(result, s)
		}
	}

	return result
}

func (m Map) Path(p Path) Path {
	result := Path{}

	for _, s := range m.MovesUp(p.Last()) {
		np := p.Extend(s)

		if s == m.Target {
			return np
		}

		if m.paths.best(np) {
			fp := m.Path(np)

			if result.Len() == 0 || fp.Len() != 0 && fp.Len() < result.Len() {
				result = fp
			}
		}
	}

	return result
}

func (m Map) PathToHeight(p Path, height int) Path {
	result := Path{}

	for _, s := range m.MovesDown(p.Last()) {
		np := p.Extend(s)

		if m.Height(s) == height {
			return np
		}

		if m.paths.best(np) {
			fp := m.PathToHeight(np, height)

			if result.Len() == 0 || fp.Len() != 0 && fp.Len() < result.Len() {
				result = fp
			}
		}
	}

	return result
}

func (m Map) String() string {
	sb := strings.Builder{}

	sb.WriteString("Map:")

	for i, row := range m.heights {
		if i == 0 {
			sb.WriteString("\n[")
		} else {
			sb.WriteString("\n ")
		}

		for _, h := range row {
			sb.WriteString(fmt.Sprintf(" %c", rune(h)+'a'))
		}
	}

	sb.WriteString(" ]\nStart:  ")
	sb.WriteString(m.Start.String())

	sb.WriteString("\nTarget: ")
	sb.WriteString(m.Target.String())

	return sb.String()
}

type Position struct {
	x, y int
}

func (p Position) String() string {
	return fmt.Sprintf("(%d, %d)", p.x, p.y)
}

func ParseMap(file string) Map {
	lr := aoc2022.NewLineReader(file)

	defer aoc2022.Close(lr)

	var start, target Position
	heights := make([][]int, 0)

	for lr.HasNext() {
		text := lr.Text()
		row := make([]int, 0)

		for i, c := range text {
			switch c {
			case 'S':
				row = append(row, 0)
				start = Position{x: i, y: lr.Count() - 1}
			case 'E':
				row = append(row, 25)
				target = Position{x: i, y: lr.Count() - 1}
			default:
				row = append(row, int(c-'a'))
			}
		}

		heights = append(heights, row)
	}

	return Map{
		heights: heights,
		paths:   newPaths(len(heights), len(heights[0])),
		Start:   start,
		Target:  target,
	}
}

type Path struct {
	positions []Position
}

func (p Path) Extend(pos Position) Path {
	result := make([]Position, p.Len()+1)

	copy(result, p.positions)

	result[p.Len()] = pos

	return Path{result}
}

func (p Path) Last() Position {
	return p.positions[p.Len()-1]
}

func (p Path) Len() int {
	return len(p.positions)
}

func (p Path) String(m Map) string {
	sb := strings.Builder{}

	sb.WriteString("Path:")

	for i, pos := range p.positions {
		sb.WriteString("\n")
		sb.WriteString(fmt.Sprintf("%2d: ", i+1))
		sb.WriteString(pos.String())
		sb.WriteString(" = ")
		sb.WriteString(fmt.Sprintf("%v", m.Height(pos)))
	}

	return sb.String()
}

type paths struct {
	paths [][]Path
}

func (p *paths) best(path Path) bool {
	pos := path.Last()
	current := p.paths[pos.y][pos.x].Len()
	result := current == 0 || path.Len() < current

	if result {
		p.paths[pos.y][pos.x] = path
	}

	return result
}

func newPaths(height, width int) *paths {
	p := make([][]Path, height)

	for i := 0; i < height; i++ {
		p[i] = make([]Path, width)
	}

	return &paths{paths: p}
}
