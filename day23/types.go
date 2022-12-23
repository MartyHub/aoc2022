package main

import (
	"aoc2022"
	"fmt"
	"log"
	"math"
	"strings"
)

type direction struct {
	name   string
	dx, dy int
}

func (d direction) String() string {
	return fmt.Sprintf("%s (%d, %d)", d.name, d.dx, d.dy)
}

var (
	N          = direction{"N", 0, -1}
	NE         = direction{"NE", 1, -1}
	E          = direction{"E", 1, 0}
	SE         = direction{"SE", 1, 1}
	S          = direction{"S", 0, 1}
	SW         = direction{"SW", -1, 1}
	W          = direction{"W", -1, 0}
	NW         = direction{"NW", -1, -1}
	directions = []direction{N, NE, E, SE, S, SW, W, NW}
)

type position struct {
	x, y int
}

func (p position) move(d direction) position {
	return position{p.x + d.dx, p.y + d.dy}
}

func (p position) String() string {
	return fmt.Sprintf("(%d, %d)", p.x, p.y)
}

type proposal struct {
	checks []direction
	move   direction
}

func (p proposal) String() string {
	return fmt.Sprintf("If %v are empty then move %v", p.checks, p.move)
}

type strategy struct {
	proposals []proposal
}

func newStrategy() *strategy {
	return &strategy{
		proposals: []proposal{
			{[]direction{NW, N, NE}, N},
			{[]direction{SW, S, SE}, S},
			{[]direction{NW, W, SW}, W},
			{[]direction{NE, E, SE}, E},
		},
	}
}

func (s *strategy) next() {
	s.proposals = append(s.proposals[1:], s.proposals[0])
}

type elf struct {
	id       int
	proposal *position
	strategy *strategy
}

func newElf(id int) *elf {
	return &elf{
		id:       id,
		strategy: newStrategy(),
	}
}

func (e *elf) String() string {
	if e.proposal == nil {
		return fmt.Sprintf("Elf # %d", e.id)
	}

	return fmt.Sprintf("Elf # %d wants to move to %v", e.id, e.proposal)
}

type grove struct {
	elves map[position]*elf
}

func parseGrove(fileName string) *grove {
	lr := aoc2022.NewLineReader(fileName)

	defer aoc2022.Close(lr)

	result := &grove{elves: make(map[position]*elf)}

	for lr.HasNext() {
		for x, r := range lr.Text() {
			if r == '#' {
				key := position{x, lr.Count() - 1}

				result.elves[key] = newElf(len(result.elves) + 1)
			}
		}
	}

	log.Println(result)

	return result
}

func (g *grove) count() int {
	max := g.max()
	min := g.min()
	result := 0

	for y := min.y; y <= max.y; y++ {
		for x := min.x; x <= max.x; x++ {
			if !g.elf(position{x, y}) {
				result++
			}
		}
	}

	return result
}

func (g *grove) draw() string {
	sb := strings.Builder{}
	max := g.max()
	min := g.min()

	for y := min.y; y <= max.y; y++ {
		if y > min.y {
			sb.WriteRune('\n')
		}

		for x := min.x; x <= max.x; x++ {
			if g.elf(position{x, y}) {
				sb.WriteRune('#')
			} else {
				sb.WriteRune('.')
			}
		}
	}

	return sb.String()
}

func (g *grove) elf(pos position) bool {
	if _, found := g.elves[pos]; found {
		return true
	}

	return false
}

func (g *grove) empty(pos position, directions []direction) bool {
	for _, d := range directions {
		if g.elf(pos.move(d)) {
			return false
		}
	}

	return true
}

func (g *grove) max() position {
	result := position{math.MinInt, math.MinInt}

	for pos := range g.elves {
		if pos.x > result.x {
			result.x = pos.x
		}

		if pos.y > result.y {
			result.y = pos.y
		}
	}

	return result
}

func (g *grove) min() position {
	result := position{math.MaxInt, math.MaxInt}

	for pos := range g.elves {
		if pos.x < result.x {
			result.x = pos.x
		}

		if pos.y < result.y {
			result.y = pos.y
		}
	}

	return result
}

func (g *grove) round() bool {
	proposals := make(map[position]int)

	for pos, elf := range g.elves {
		elf.proposal = nil

		if !g.empty(pos, directions) {
			for _, proposal := range elf.strategy.proposals {
				if g.empty(pos, proposal.checks) {
					move := pos.move(proposal.move)

					elf.proposal = &move

					proposals[move]++

					break
				}
			}

		}

		elf.strategy.next()
	}

	elves := make(map[position]*elf)

	for pos, elf := range g.elves {
		if elf.proposal == nil || proposals[*elf.proposal] > 1 {
			if e, found := elves[pos]; found {
				panic(fmt.Sprintf("Conflict between:\n - %v\n - %v", elf, e))
			}

			elves[pos] = elf
		} else {
			if e, found := elves[*elf.proposal]; found {
				panic(fmt.Sprintf("Conflict between:\n - %v\n - %v", elf, e))
			}

			elves[*elf.proposal] = elf
		}
	}

	if len(elves) != len(g.elves) {
		panic(fmt.Sprintf("Lost elves: %d != %d", len(elves), len(g.elves)))
	}

	g.elves = elves

	return len(proposals) == 0
}

func (g *grove) String() string {
	return fmt.Sprintf("Grove %v -> %v with %d elves", g.min(), g.max(), len(g.elves))
}
