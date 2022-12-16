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

type State struct {
	Minute     int
	OpenValves string
	Position   string
}

var cache = make(map[State]int)

type Cave struct {
	Iteration  int
	Minutes    int
	OpenValves []string
	Position   string
	Valves     map[string]Valve
}

func (c Cave) IsOpen(valve string) bool {
	for _, name := range c.OpenValves {
		if name == valve {
			return true
		}
	}

	return false
}

func (c Cave) Names() []string {
	result := make([]string, 0)

	for name := range c.Valves {
		result = append(result, name)
	}

	sort.Strings(result)

	return result
}

func (c Cave) Move(position string) Cave {
	return Cave{
		Iteration:  c.Iteration + 1,
		Minutes:    c.Minutes,
		OpenValves: c.OpenValves,
		Position:   position,
		Valves:     c.Valves,
	}
}

func (c Cave) Open() Cave {
	openValves := c.OpenValves

	openValves = append(openValves, c.Position)

	return Cave{
		Iteration:  c.Iteration + 1,
		Minutes:    c.Minutes,
		OpenValves: openValves,
		Position:   c.Position,
		Valves:     c.Valves,
	}
}

func (c Cave) Pressure() int {
	result := 0

	for _, name := range c.OpenValves {
		result += c.Valves[name].FlowRate
	}

	return result
}

func (c Cave) doReleaseMax() int {
	if c.Iteration == c.Minutes {
		return 0
	}

	result := 0
	valve := c.Valves[c.Position]

	if valve.FlowRate != 0 && !c.IsOpen(c.Position) {
		result = c.Open().ReleaseMax()
	}

	for _, position := range valve.Tunnels {
		result = aoc2022.Max(result, c.Move(position).ReleaseMax())
	}

	return result + c.Pressure()
}

func (c Cave) ReleaseMax() int {
	state := c.State()
	result, exist := cache[state]

	if exist {
		return result
	}

	result = c.doReleaseMax()

	cache[state] = result

	return result
}

func (c Cave) State() State {
	return State{
		Minute:     c.Iteration,
		OpenValves: strings.Join(c.OpenValves, ","),
		Position:   c.Position,
	}
}

func (c Cave) Status() string {
	sb := strings.Builder{}

	for _, name := range c.Names() {
		sb.WriteString(c.Valves[name].String())
		sb.WriteString("\n")
	}

	return sb.String()
}

func (c Cave) String() string {
	sb := strings.Builder{}

	sb.WriteString(fmt.Sprintf("== Iteration %d ==\n", c.Iteration))

	switch len(c.OpenValves) {
	case 0:
		sb.WriteString("No valves are open.\n")
	case 1:
		sb.WriteString(fmt.Sprintf("Valve %v is open, releasing %v pressure.\n", c.OpenValves[0], c.Pressure()))
	default:
		sb.WriteString(fmt.Sprintf("Valves %v are open, releasing %v pressure.\n", c.OpenValves, c.Pressure()))
	}

	sb.WriteString(fmt.Sprintf("Position: %v\n\n", c.Position))

	return sb.String()
}

func ParseCave(fileName string, minutes int) Cave {
	result := Cave{
		Minutes:    minutes,
		OpenValves: make([]string, 0),
		Position:   "AA",
		Valves:     make(map[string]Valve),
	}

	lr := aoc2022.NewLineReader(fileName)

	defer aoc2022.Close(lr)

	for lr.HasNext() {
		valve := ParseValve(lr.Text())

		result.Valves[valve.Name] = valve
	}

	return result
}
