package main

import (
	"aoc2022"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type Notes struct {
	board [][]rune
	cube  bool
	path  []string
}

func (n Notes) Draw() string {
	sb := strings.Builder{}

	for i, row := range n.board {
		if i > 0 {
			sb.WriteString("\n")
		}

		sb.WriteString(string(row))
	}

	return sb.String()
}

func (n Notes) Get(p Point) rune {
	if p.y < 0 || p.y >= len(n.board) {
		return ' '
	}

	row := n.board[p.y]

	if p.x < 0 || p.x >= len(row) {
		return ' '
	}

	return row[p.x]
}

func (n Notes) GetStartingPointOfColumn(x int) Point {
	for y := 0; y < len(n.board); y++ {
		pt := Point{x, y}

		if n.Get(pt) != ' ' {
			return pt
		}
	}

	log.Fatalf("No starting point found in column %d", x)

	return Point{}
}

func (n Notes) GetEndingPointOfColumn(x int) Point {
	for y := len(n.board) - 1; y >= 0; y-- {
		pt := Point{x, y}

		if n.Get(pt) != ' ' {
			return pt
		}
	}

	log.Fatalf("No ending point found in column %d", x)

	return Point{}
}

func (n Notes) GetStartingPointOfRow(y int) Point {
	row := n.board[y]

	for x, c := range row {
		if c != ' ' {
			return Point{x, y}
		}
	}

	log.Fatalf("No starting point found in row %d: %v", y, row)

	return Point{}
}

func (n Notes) GetEndingPointOfRow(y int) Point {
	row := n.board[y]

	for x := len(row) - 1; x >= 0; x-- {
		if row[x] != ' ' {
			return Point{x, y}
		}
	}

	log.Fatalf("No ending point found in row %d: %v", y, row)

	return Point{}
}

func (n Notes) Password1() int {
	state := NewState(n.Start())

	for _, action := range n.path {
		switch action {
		case "L":
			state = state.TurnLeft()
		case "R":
			state = state.TurnRight()
		default:
			for i := 0; i < aoc2022.Must(strconv.Atoi(action)); i++ {
				state = n.Move(state)
			}
		}
	}

	return state.Password()
}

func (n Notes) Start() Point {
	return n.GetStartingPointOfRow(0)
}

func (n Notes) Move(state State) State {
	result := state

	result.pos = state.pos.Move(state.face)

	if n.Get(result.pos) == ' ' {
		result = n.Wrap(result)
	}

	if n.Get(result.pos) == '.' {
		return result
	}

	return state
}

func (n Notes) Wrap(state State) State {
	result := state

	switch state.face {
	case left:
		result.pos = n.GetEndingPointOfRow(state.pos.y)
	case right:
		result.pos = n.GetStartingPointOfRow(state.pos.y)
	case up:
		result.pos = n.GetEndingPointOfColumn(state.pos.x)
	case down:
		result.pos = n.GetStartingPointOfColumn(state.pos.x)
	}

	return result
}

func (n Notes) String() string {
	return fmt.Sprintf("Board of height %d", len(n.board))
}

func ParseNotes(fileName string, cube bool) Notes {
	lr := aoc2022.NewLineReader(fileName)

	defer aoc2022.Close(lr)

	result := Notes{
		board: [][]rune{},
		cube:  cube,
	}

	for lr.HasNext() {
		text := lr.Text()

		if text == "" {
			break
		}

		result.board = append(result.board, []rune(text))
	}

	if lr.HasNext() {
		result.path = parseActions(lr.Text())
	} else {
		log.Fatalf("No path found in %v", fileName)
	}

	return result
}

func parseActions(s string) []string {
	re := regexp.MustCompile(`(\d+|L|R)`)

	return re.FindAllString(s, -1)
}
