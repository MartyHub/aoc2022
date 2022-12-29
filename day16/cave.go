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
		} else {
			for _, link := range valve.links {
				move := s.Move(pressure, link)

				if move.Iteration <= minutes && move.ShouldExplore() {
					queue = append(queue, move)
				}
			}
		}
	}

	return result
}

func (c Cave) Iterate2(minutes int) aoc2022.Pair[State] {
	queue := []aoc2022.Pair[State]{aoc2022.NewPair(Start(), Start())}
	visited := map[aoc2022.Pair[StateKey]]bool{}
	result := aoc2022.NewPair(State{}, State{})

	for len(queue) > 0 {
		pair := queue[0]
		queue = queue[1:]

		//log.Printf("%v / %v", pair.First, pair.Second)
		//log.Printf("Queue size: %d", len(queue))

		key1 := aoc2022.NewPair(pair.First.Key(), pair.Second.Key())

		if visited[key1] {
			continue
		}

		visited[key1] = true
		visited[aoc2022.NewPair(pair.Second.Key(), pair.First.Key())] = true

		if pair.First.Iteration == minutes && pair.Second.Iteration == minutes {
			if pair.First.TotalPressure()+pair.Second.TotalPressure() >
				result.First.TotalPressure()+result.Second.TotalPressure() {
				result = pair
			}

			continue
		}

		pressure1 := pair.First.ComputePressure(c)
		pressure2 := pair.Second.ComputePressure(c)
		valve1 := c.Valve(pair.First.Position())
		valve2 := c.Valve(pair.Second.Position())
		queue1 := make([]State, 0)
		queue2 := make([]State, 0)

		currentBest := 2213
		topRemainingPressure := c.TopRemainingPressure(
			aoc2022.CopyAndAppends(pair.First.OpenValves, pair.Second.OpenValves),
			(minutes-aoc2022.Min(pair.First.Iteration, pair.Second.Iteration))/2,
		)
		totalPressure := pair.First.TotalPressure() + pair.Second.TotalPressure() +
			pressure1*(minutes-pair.First.Iteration) + pressure2*(minutes-pair.Second.Iteration) +
			topRemainingPressure

		if totalPressure < currentBest {
			continue
		}

		if pair.First.Iteration < minutes {
			if valve1.FlowRate != 0 && !pair.First.IsOpen(valve1.Name) && !pair.Second.IsOpen(valve1.Name) {
				queue1 = append(queue1, pair.First.Open(pressure1))
			}

			for _, link := range valve1.links {
				move := pair.First.Move(pressure1, link)

				if move.Iteration <= minutes {
					queue1 = append(queue1, move)
				}
			}

			if len(queue1) == 0 {
				queue1 = append(queue1, pair.First.Fill(minutes, pressure1))
			}
		} else {
			queue1 = append(queue1, pair.First)
		}

		if pair.Second.Iteration < minutes {
			if valve2.FlowRate != 0 && !pair.First.IsOpen(valve2.Name) && !pair.Second.IsOpen(valve2.Name) {
				queue2 = append(queue2, pair.Second.Open(pressure2))
			}

			for _, link := range valve2.links {
				move := pair.Second.Move(pressure2, link)

				if move.Iteration <= minutes {
					queue2 = append(queue2, move)
				}
			}

			if len(queue2) == 0 {
				queue2 = append(queue2, pair.Second.Fill(minutes, pressure2))
			}
		} else {
			queue2 = append(queue2, pair.Second)
		}

		for _, state1 := range queue1 {
			for _, state2 := range queue2 {
				if aoc2022.Contains(state1.OpenValves, state2.LastOpenValve()) || aoc2022.Contains(state2.OpenValves, state1.LastOpenValve()) {
					continue
				}

				queue = append(queue, aoc2022.NewPair(state1, state2))
			}
		}

		//queue3 := make([]aoc2022.Pair[State], 0)
		//
		//if pair.First.Iteration < minutes {
		//	if valve1.FlowRate != 0 && !pair.First.IsOpen(valve1.Name) && !pair.Second.IsOpen(valve1.Name) {
		//		queue1 = append(queue1, pair.First.Open(pressure1))
		//	}
		//}
		//
		//if pair.Second.Iteration < minutes {
		//	if valve2.FlowRate != 0 && !pair.First.IsOpen(valve2.Name) && !pair.Second.IsOpen(valve2.Name) {
		//		queue2 = append(queue2, pair.Second.Open(pressure2))
		//	}
		//}
		//
		//if len(queue1) == 1 {
		//	if len(queue2) == 1 {
		//		if valve1.Name == valve2.Name {
		//			queue2 = queue2[1:]
		//		} else {
		//			queue3 = append(queue3, aoc2022.NewPair(queue1[0], queue2[0]))
		//		}
		//	}
		//}
		//
		//if pair.First.Iteration < minutes && len(queue1) == 0 {
		//	for _, link := range valve1.links {
		//		move := pair.First.Move(pressure1, link)
		//
		//		if move.Iteration <= minutes && move.ShouldExplore() {
		//			queue1 = append(queue1, move)
		//		}
		//	}
		//}
		//
		//if pair.Second.Iteration < minutes && len(queue2) == 0 {
		//	for _, link := range valve2.links {
		//		move := pair.Second.Move(pressure2, link)
		//
		//		if move.Iteration <= minutes && move.ShouldExplore() {
		//			queue2 = append(queue2, move)
		//		}
		//	}
		//}
		//
		//if len(queue3) == 0 {
		//	for _, state1 := range queue1 {
		//		for _, state2 := range queue2 {
		//			if !state2.Path.Contains(state1.Position()) && !state1.Path.Contains(state2.Position()) {
		//				queue3 = append(queue3, aoc2022.NewPair(state1, state2))
		//			}
		//		}
		//	}
		//}
		//
		//for _, p := range queue3 {
		//	if len(p.First.OpenValves)+len(p.Second.OpenValves) == c.MaxOpenValves() {
		//		log.Printf("Filling %v / %v", p.First, p.Second)
		//		queue = append(queue, aoc2022.NewPair(
		//			p.First.Fill(minutes, p.First.ComputePressure(c)),
		//			p.Second.Fill(minutes, p.Second.ComputePressure(c)),
		//		))
		//	} else {
		//		queue = append(queue, p)
		//	}
		//}
	}

	return result
}

func (c Cave) MaxOpenValves() int {
	return len(c.valves) - 1
}

func (c Cave) TopRemainingPressure(openValves []string, top int) int {
	flowRates := make([]int, 0)

	for _, valve := range c.valves {
		if valve.FlowRate > 0 && !aoc2022.Contains(openValves, valve.Name) {
			flowRates = append(flowRates, valve.FlowRate)
		}
	}

	sort.Ints(flowRates)

	result := 0

	for i := 0; i < top; i++ {
		index := len(flowRates) - 1 - i

		if index < 0 {
			break
		}

		result += flowRates[index] * (top - i) * 2
	}

	return result
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
