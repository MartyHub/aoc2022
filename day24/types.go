package main

import (
	"aoc2022"
	"fmt"
	"log"
	"math"
	"reflect"
	"sort"
	"strings"
)

type direction rune

const (
	up    direction = '^'
	right direction = '>'
	down  direction = 'v'
	left  direction = '<'
)

var directions = []direction{down, right, up, left}

func (d direction) String() string {
	switch d {
	case up:
		return "up"
	case right:
		return "right"
	case down:
		return "down"
	case left:
		return "left"
	default:
		return "unknown"
	}
}

type position struct {
	x, y int
}

var firstStart = position{1, 0}

func (p position) move(d direction) position {
	switch d {
	case up:
		return position{p.x, p.y - 1}
	case right:
		return position{p.x + 1, p.y}
	case down:
		return position{p.x, p.y + 1}
	case left:
		return position{p.x - 1, p.y}
	default:
		return p
	}
}

func (p position) String() string {
	return fmt.Sprintf("(%d, %d)", p.x, p.y)
}

type state struct {
	blizzards          map[position][]direction
	iteration          int
	start, pos, target position
}

func (s state) equals(other state) bool {
	return s.pos == other.pos && reflect.DeepEqual(s.blizzards, other.blizzards)
}

func (s state) String() string {
	return fmt.Sprintf("Iteration # %d: %v", s.iteration, s.pos)
}

type visit struct {
	states []state
}

func (v *visit) contains(s state) bool {
	for _, vs := range v.states {
		if s.equals(vs) {
			return true
		}
	}

	return false
}

func (v *visit) add(s state) {
	v.states = append(v.states, s)
}

type valley struct {
	blizzards     map[position][]direction
	height, width int
	pos           position
	target        position
}

func (v valley) draw(s state) string {
	sb := strings.Builder{}

	for y := 0; y < v.height; y++ {
		if y > 0 {
			sb.WriteString("\n")
		}

		for x := 0; x < v.width; x++ {
			sb.WriteRune(v.get(s, position{x, y}))
		}
	}

	return sb.String()
}

func (v valley) get(s state, pos position) rune {
	if pos == s.pos {
		return 'E'
	}

	if pos == s.start || pos == s.target {
		return '.'
	}

	if pos.x == 0 || pos.x == v.width-1 || pos.y == 0 || pos.y == v.height-1 {
		return '#'
	}

	blizzards := s.blizzards[pos]

	switch len(blizzards) {
	case 0:
		return '.'
	case 1:
		return rune(blizzards[0])
	default:
		return rune(len(blizzards) + '0')
	}
}

type shortestPaths struct {
	iterations map[position]int
}

func newShortestPaths(v valley, start, target position) shortestPaths {
	result := shortestPaths{make(map[position]int)}

	result.iterations[start] = 0
	result.iterations[target] = math.MaxInt

	for y := 1; y < v.height-1; y++ {
		for x := 1; x < v.width-1; x++ {
			result.iterations[position{x, y}] = math.MaxInt
		}
	}

	return result
}

func (v valley) cycle() int {
	s := state{v.blizzards, 0, firstStart, firstStart, v.target}

	for i := 1; ; i++ {
		s = v.moveBlizzards(s)

		if reflect.DeepEqual(s.blizzards, v.blizzards) {
			return i
		}
	}
}

func (v valley) iterate(start state, maxWait int) int {
	paths := newShortestPaths(v, start.start, start.target)
	queue := []state{start}
	visited := visit{}

	for len(queue) > 0 {
		s := queue[0]
		queue = queue[1:]

		if visited.contains(s) {
			continue
		}

		visited.add(s)

		if paths.iterations[s.pos] != math.MaxInt && s.iteration > paths.iterations[s.pos]+maxWait {
			continue
		}

		if paths.iterations[s.pos] == math.MaxInt {
			paths.iterations[s.pos] = s.iteration
		}

		if s.pos == s.target {
			continue
		}

		mbs := v.moveBlizzards(s)

		for _, move := range v.moves(mbs) {
			if paths.iterations[move.pos] == math.MaxInt || move.iteration <= paths.iterations[move.pos]+maxWait {
				queue = append(queue, move)
			}
		}
	}

	return paths.iterations[start.target]
}

func (v valley) moves(s state) []state {
	result := make([]state, 0, 5)

	for _, d := range directions {
		pos := s.pos.move(d)

		if v.valid(s, pos) {
			result = append(result, state{s.blizzards, s.iteration, s.start, pos, s.target})
		}
	}

	if v.valid(s, s.pos) {
		result = append(result, state{s.blizzards, s.iteration, s.start, s.pos, s.target})
	}

	return result
}

func (v valley) blizzardsAfter(minutes int) map[position][]direction {
	s := state{v.blizzards, 0, firstStart, firstStart, v.target}

	for i := 0; i < minutes; i++ {
		s = v.moveBlizzards(s)
	}

	return s.blizzards
}

func (v valley) moveBlizzards(s state) state {
	newBlizzards := make(map[position][]direction)

	for pos, directions := range s.blizzards {
		for _, d := range directions {
			newPos := v.moveBlizzard(pos, d)
			blizzards := append(newBlizzards[newPos], d)

			if len(blizzards) > 1 {
				sort.Slice(blizzards, func(i, j int) bool {
					return blizzards[i] < blizzards[j]
				})
			}

			newBlizzards[newPos] = blizzards
		}
	}

	return state{
		blizzards: newBlizzards,
		iteration: s.iteration + 1,
		start:     s.start,
		pos:       s.pos,
		target:    s.target,
	}
}

func (v valley) moveBlizzard(pos position, d direction) position {
	result := pos.move(d)

	if result.x == 0 {
		result.x = v.width - 2
	} else if result.x == v.width-1 {
		result.x = 1
	} else if result.y == 0 {
		result.y = v.height - 2
	} else if result.y == v.height-1 {
		result.y = 1
	}

	return result
}

func (v valley) valid(s state, pos position) bool {
	return pos == v.pos || pos == s.start || pos == s.target || pos.x > 0 && pos.x < v.width-1 && pos.y > 0 && pos.y < v.height-1 && len(s.blizzards[pos]) == 0
}

func (v valley) String() string {
	return fmt.Sprintf("valley %dx%d with %d blizzards -> %v", v.width, v.height, len(v.blizzards), v.target)
}

func parseValley(fileName string) valley {
	lr := aoc2022.NewLineReader(fileName)

	defer aoc2022.Close(lr)

	result := valley{
		blizzards: make(map[position][]direction),
		pos:       position{1, 0},
	}

	for lr.HasNext() {
		text := lr.Text()

		if result.width == 0 {
			result.width = len(text)
		}

		for x, r := range text {
			if r != '#' && r != '.' {
				key := position{x, lr.Count() - 1}

				result.blizzards[key] = append(result.blizzards[key], direction(r))
			}
		}
	}

	result.height = lr.Count()
	result.target = position{result.width - 2, result.height - 1}

	log.Println(result)

	return result
}
