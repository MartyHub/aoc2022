package main

import (
	"aoc2022"
	"fmt"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Valve struct {
	Name     string
	FlowRate int
	Tunnels  []string
}

func (v Valve) String() string {
	return fmt.Sprintf("Valve %v: flowRate=%v, tunnels=%v", v.Name, v.FlowRate, v.Tunnels)
}

var valveRegex = regexp.MustCompile(`Valve (\w+) has flow rate=(\d+); tunnels? leads? to valves? (.+)`)

func ParseValve(s string) Valve {
	tokens := valveRegex.FindStringSubmatch(s)

	if len(tokens) < 4 {
		log.Fatalf("Failed to parse valve: %v", s)
	}

	tunnels := strings.Split(tokens[3], ", ")

	sort.Strings(tunnels)

	return Valve{
		Name:     tokens[1],
		FlowRate: aoc2022.Must(strconv.Atoi(tokens[2])),
		Tunnels:  tunnels,
	}
}

type Cave struct {
	Minutes           int
	Valves            map[string]Valve
	maxReleaseByState map[StateKey]int
}

func (c Cave) Pressure(s State) int {
	result := 0

	for _, name := range s.OpenValves {
		result += c.Valves[name].FlowRate
	}

	return result
}

func (c Cave) doReleaseMax(s State) int {
	if s.Iteration == c.Minutes {
		return 0
	}

	result := 0
	valve := c.Valves[s.Position]

	if valve.FlowRate != 0 && !s.IsOpen(s.Position) {
		result = c.ReleaseMax(s.Open())
	}

	for _, position := range valve.Tunnels {
		result = aoc2022.Max(result, c.ReleaseMax(s.Move(position)))
	}

	return result + c.Pressure(s)
}

func (c Cave) ReleaseMax(s State) int {
	key := s.StateKey()
	result, exist := c.maxReleaseByState[key]

	if exist {
		return result
	}

	result = c.doReleaseMax(s)

	c.maxReleaseByState[key] = result

	return result
}

func (c Cave) String(s State) string {
	sb := strings.Builder{}

	sb.WriteString(fmt.Sprintf("== Iteration %d ==\n", s.Iteration))

	switch len(s.OpenValves) {
	case 0:
		sb.WriteString("No valves are open.\n")
	case 1:
		sb.WriteString(fmt.Sprintf("Valve %v is open, releasing %v pressure.\n", s.OpenValves[0], c.Pressure(s)))
	default:
		sb.WriteString(fmt.Sprintf("Valves %v are open, releasing %v pressure.\n", s.OpenValves, c.Pressure(s)))
	}

	sb.WriteString(fmt.Sprintf("Position: %v\n\n", s.Position))

	return sb.String()
}

type StateKey struct {
	Minute     int
	OpenValves string
	Position   string
}

type State struct {
	Iteration  int
	OpenValves []string
	Position   string
}

func Single() State {
	return State{Position: "AA"}
}

func (s State) IsOpen(valve string) bool {
	for _, name := range s.OpenValves {
		if name == valve {
			return true
		}
	}

	return false
}

func (s State) Move(position string) State {
	return State{
		Iteration:  s.Iteration + 1,
		OpenValves: s.OpenValves,
		Position:   position,
	}
}

func (s State) Open() State {
	openValves := s.OpenValves

	openValves = append(openValves, s.Position)

	return State{
		Iteration:  s.Iteration + 1,
		OpenValves: openValves,
		Position:   s.Position,
	}
}

func (s State) StateKey() StateKey {
	return StateKey{
		Minute:     s.Iteration,
		OpenValves: strings.Join(s.OpenValves, ","),
		Position:   s.Position,
	}
}

func ParseCave(fileName string, minutes int) Cave {
	result := Cave{
		Minutes:           minutes,
		Valves:            make(map[string]Valve),
		maxReleaseByState: make(map[StateKey]int),
	}

	lr := aoc2022.NewLineReader(fileName)

	defer aoc2022.Close(lr)

	for lr.HasNext() {
		valve := ParseValve(lr.Text())

		result.Valves[valve.Name] = valve
	}

	return result
}
