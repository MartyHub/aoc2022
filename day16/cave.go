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

const StartValve = "AA"

type Link struct {
	Cost   int
	Target string
}

func defaultLinks(targets []string) []Link {
	result := make([]Link, len(targets))

	for i, target := range targets {
		result[i] = newLink(1, target)
	}

	return result
}

func newLink(cost int, target string) Link {
	return Link{
		Cost:   cost,
		Target: target,
	}
}

func (l Link) String() string {
	return fmt.Sprintf("cost %d to %s", l.Cost, l.Target)
}

type Valve struct {
	Name     string
	FlowRate int
	links    []Link
}

func newValve(name string, flowRate int, links []Link) Valve {
	return Valve{
		Name:     name,
		FlowRate: flowRate,
		links:    links,
	}
}

var valveRegex = regexp.MustCompile(`Valve (\w+) has flow rate=(\d+); tunnels? leads? to valves? (.+)`)

func parseValve(s string) Valve {
	tokens := valveRegex.FindStringSubmatch(s)

	if len(tokens) < 4 {
		log.Fatalf("Failed to parse valve: %v", s)
	}

	tunnels := strings.Split(tokens[3], ", ")

	sort.Strings(tunnels)

	return newValve(
		tokens[1],
		aoc2022.Must(strconv.Atoi(tokens[2])),
		defaultLinks(tunnels),
	)
}

func (v Valve) Link(target string) Link {
	for _, link := range v.links {
		if link.Target == target {
			return link
		}
	}

	return Link{}
}

func (v Valve) String() string {
	return fmt.Sprintf("%s +%2d: %v", v.Name, v.FlowRate, v.links)
}

type Cave struct {
	valves map[string]Valve
}

func newCave() Cave {
	return Cave{
		valves: make(map[string]Valve),
	}
}

func (c Cave) reduceLink(source string, link Link) []Link {
	targetValve := c.valves[link.Target]
	result := make([]Link, 0, len(targetValve.links))

	for _, subLink := range targetValve.links {
		if subLink.Target == source {
			continue
		}

		result = append(result, newLink(link.Cost+subLink.Cost, subLink.Target))
	}

	return result
}

func (c Cave) reduceValve(valve Valve) Valve {
	newLinks := make([]Link, 0, len(valve.links))
	queue := valve.links
	visited := make(map[string]bool)

	for len(queue) > 0 {
		link := queue[0]
		queue = queue[1:]

		if visited[link.Target] {
			continue
		}

		visited[link.Target] = true

		targetValve := c.valves[link.Target]

		if targetValve.Name != StartValve && targetValve.FlowRate == 0 {
			queue = append(queue, c.reduceLink(valve.Name, link)...)
		} else {
			newLinks = append(newLinks, link)
		}
	}

	return newValve(valve.Name, valve.FlowRate, newLinks)
}

func (c Cave) reduce() Cave {
	result := newCave()
	queue := []string{StartValve}

	for len(queue) > 0 {
		name := queue[0]
		queue = queue[1:]

		if _, found := result.valves[name]; found {
			continue
		}

		newValve := c.reduceValve(c.valves[name])

		result.valves[name] = newValve

		for _, link := range newValve.links {
			queue = append(queue, link.Target)
		}
	}

	return result
}

func (c Cave) Iterate(minutes int) State {
	queue := []State{Start()}
	visited := map[StateKey]bool{}
	result := State{}

	for len(queue) > 0 {
		s := queue[0]
		queue = queue[1:]

		key := s.Key()

		if visited[key] {
			continue
		}

		visited[key] = true

		if s.Iteration == minutes {
			if s.TotalPressure() > result.TotalPressure() {
				result = s
			}

			continue
		}

		pressure := s.ComputePressure(c)
		valve := c.Valve(s.Position())

		if valve.FlowRate != 0 && !s.IsOpen(valve.Name) {
			open := s.Open(pressure)

			if len(open.OpenValves) == c.MaxOpenValves() {
				open = open.Fill(minutes, open.ComputePressure(c))
			}

			queue = append(queue, open)
		}

		for _, link := range valve.links {
			move := s.Move(pressure, link)

			if move.Iteration <= minutes && move.ShouldExplore() {
				queue = append(queue, move)
			}
		}
	}

	return result
}

func (c Cave) MaxOpenValves() int {
	return len(c.valves) - 1
}

func (c Cave) Valve(name string) Valve {
	return c.valves[name]
}

func (c Cave) String() string {
	sb := strings.Builder{}

	sb.WriteString("Cave:")

	queue := []string{StartValve}
	visited := make(map[string]bool)

	for len(queue) > 0 {
		name := queue[0]
		queue = queue[1:]

		if visited[name] {
			continue
		}

		visited[name] = true

		valve := c.valves[name]

		sb.WriteString(fmt.Sprintf("\n%v", valve))

		for _, link := range valve.links {
			queue = append(queue, link.Target)
		}
	}

	return sb.String()
}

func ParseCave(fileName string) Cave {
	lr := aoc2022.NewLineReader(fileName)
	defer aoc2022.Close(lr)

	result := newCave()

	for lr.HasNext() {
		valve := parseValve(lr.Text())

		result.valves[valve.Name] = valve
	}

	result = result.reduce()

	log.Println(result)

	return result
}
