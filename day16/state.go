package main

import (
	"aoc2022"
	"fmt"
	"strings"
)

type StateKey struct {
	Iteration  int
	OpenValves string
	Position   string
}

func newStateKey(iteration int, openValves string, position string) StateKey {
	return StateKey{
		Iteration:  iteration,
		OpenValves: openValves,
		Position:   position,
	}
}

type State struct {
	Iteration  int
	OpenValves []string
	Path       aoc2022.Path[string]
	Pressures  []int
}

func newState(iteration int, openValves []string, path aoc2022.Path[string], pressures []int) State {
	return State{
		Iteration:  iteration,
		OpenValves: openValves,
		Path:       path,
		Pressures:  pressures,
	}
}

func Start() State {
	return newState(0, []string{}, aoc2022.NewPath(StartValve), []int{0})
}

func (s State) Fill(minutes int, pressure int) State {
	for i := s.Iteration; i < minutes; i++ {
		s.Iteration++
		s.Pressures = append(s.Pressures, pressure)
	}

	return s
}

func (s State) Key() StateKey {
	return newStateKey(s.Iteration, strings.Join(s.OpenValves, ","), s.Position())
}

func (s State) IsOpen(valve string) bool {
	for _, name := range s.OpenValves {
		if name == valve {
			return true
		}
	}

	return false
}

func (s State) Move(pressure int, link Link) State {
	return newState(
		s.Iteration+link.Cost,
		aoc2022.Copy(s.OpenValves),
		s.Path.Extend(link.Target),
		aoc2022.CopyAndAppend(s.Pressures, pressure*link.Cost),
	)
}

func (s State) Open(pressure int) State {
	valve := s.Position()

	return newState(
		s.Iteration+1,
		aoc2022.CopyAndAppend(s.OpenValves, valve),
		s.Path.Extend(valve),
		aoc2022.CopyAndAppend(s.Pressures, pressure),
	)
}

func (s State) Position() string {
	return s.Path.Last()
}

func (s State) ComputePressure(c Cave) int {
	result := 0

	for _, name := range s.OpenValves {
		result += c.Valve(name).FlowRate
	}

	return result
}

func (s State) TotalPressure() int {
	result := 0

	for _, pressure := range s.Pressures {
		result += pressure
	}

	return result
}

func (s State) ShouldExplore() bool {
	valve := s.Position()

	for i := s.Path.Length() - 3; i >= 0; i-- {
		if s.Path.Steps[i] == valve {
			return false
		} else if s.Path.Steps[i] == s.Path.Steps[i+1] {
			return true
		}
	}

	return true
}

func (s State) Debug(cave Cave) string {
	sb := strings.Builder{}
	minute := 0

	for i := 1; i < s.Path.Length(); i++ {
		valve := s.Path.Steps[i]

		if i > 1 {
			sb.WriteString("\n")
		}

		if valve == s.Path.Steps[i-1] {
			minute++
			sb.WriteString(fmt.Sprintf("Minute %2d:   opening %s +%3d", minute, valve, s.Pressures[i]))
		} else {
			link := cave.Valve(s.Path.Steps[i-1]).Link(valve)
			minute += link.Cost
			sb.WriteString(fmt.Sprintf("Minute %2d: moving to %s +%3d", minute, valve, s.Pressures[i]))
		}
	}

	for i := s.Path.Length(); i < len(s.Pressures); i++ {
		minute++
		sb.WriteString(fmt.Sprintf("\nMinute %2d: waiting      +%3d", minute, s.Pressures[i]))
	}

	return sb.String()
}

func (s State) String() string {
	sb := strings.Builder{}

	for i, name := range s.Path.Steps {
		if i > 0 {
			sb.WriteString(", ")
		}

		sb.WriteString(fmt.Sprintf("%s+%d", name, s.Pressures[i]))
	}

	return sb.String()
}
