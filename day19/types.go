package main

import (
	"aoc2022"
	"fmt"
	"log"
	"strings"
)

type RobotType string

const (
	Ore      RobotType = "Ore"
	Clay     RobotType = "Clay"
	Obsidian RobotType = "Obsidian"
	Geode    RobotType = "Geode"
)

type Robot struct {
	OreCost      int
	ClayCost     int
	ObsidianCost int
	Type         RobotType
}

func (r Robot) CanBuild(s State) bool {
	return s.Ore >= r.OreCost && s.Clay >= r.ClayCost && s.Obsidian >= r.ObsidianCost
}

type Blueprint struct {
	Id                           int
	OreRobot                     Robot
	ClayRobot                    Robot
	ObsidianRobot                Robot
	GeodeRobot                   Robot
	MaxOre, MaxClay, MaxObsidian int
}

func NewBlueprint() Blueprint {
	return Blueprint{
		OreRobot:      Robot{Type: Ore},
		ClayRobot:     Robot{Type: Clay},
		ObsidianRobot: Robot{Type: Obsidian},
		GeodeRobot:    Robot{Type: Geode},
	}
}

func (bp Blueprint) CanBuildOreRobot(s State) bool {
	return bp.OreRobot.CanBuild(s)
}

func (bp Blueprint) CanBuildClayRobot(s State) bool {
	return bp.ClayRobot.CanBuild(s)
}

func (bp Blueprint) CanBuildObsidianRobot(s State) bool {
	return bp.ObsidianRobot.CanBuild(s)
}

func (bp Blueprint) CanBuildGeodeRobot(s State) bool {
	return bp.GeodeRobot.CanBuild(s)
}

var iterationCache map[int]int
var stateCache map[State]int

func ResetCaches() {
	iterationCache = make(map[int]int)
	stateCache = make(map[State]int)
}

func (bp Blueprint) doRun(s State) int {
	if s.Iteration == s.Duration {
		return s.Geode
	}

	result := 0

	for _, child := range s.Children(bp) {
		result = aoc2022.Max(result, bp.Run(child))
	}

	return result
}

func (bp Blueprint) Run(s State) int {
	if result, exists := stateCache[s]; exists {
		return result
	}

	result := bp.doRun(s)

	stateCache[s] = result

	return result
}

func (bp Blueprint) String() string {
	sb := strings.Builder{}

	sb.WriteString(fmt.Sprintf("Blueprint %d:\n", bp.Id))
	sb.WriteString(
		fmt.Sprintf("\tEach ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.\n",
			bp.OreRobot.OreCost,
			bp.ClayRobot.OreCost,
			bp.ObsidianRobot.OreCost,
			bp.ObsidianRobot.ClayCost,
			bp.GeodeRobot.OreCost,
			bp.GeodeRobot.ObsidianCost,
		),
	)
	sb.WriteString(fmt.Sprintf("\tMax Ore: %d, Max Clay: %d, Max Obsidian: %d", bp.MaxOre, bp.MaxClay, bp.MaxObsidian))

	return sb.String()
}

func ParseBlueprint(s string) Blueprint {
	result := NewBlueprint()

	if _, err := fmt.Sscanf(
		s,
		"Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.",
		&result.Id,
		&result.OreRobot.OreCost,
		&result.ClayRobot.OreCost,
		&result.ObsidianRobot.OreCost,
		&result.ObsidianRobot.ClayCost,
		&result.GeodeRobot.OreCost,
		&result.GeodeRobot.ObsidianCost,
	); err != nil {
		log.Fatalf("Failed to parse blueprint %s", s)
	}

	for _, robot := range []Robot{result.OreRobot, result.ClayRobot, result.ObsidianRobot, result.GeodeRobot} {
		result.MaxOre = aoc2022.Max(result.MaxOre, robot.OreCost)
		result.MaxClay = aoc2022.Max(result.MaxClay, robot.ClayCost)
		result.MaxObsidian = aoc2022.Max(result.MaxObsidian, robot.ObsidianCost)
	}

	return result
}

type Factory struct {
	Blueprints []Blueprint
}

func NewFactory() Factory {
	return Factory{}
}

func (f Factory) String() string {
	sb := strings.Builder{}

	sb.WriteString(fmt.Sprintf("Factory with %d Blueprints:\n", len(f.Blueprints)))

	for _, bp := range f.Blueprints {
		sb.WriteString(fmt.Sprintf("%v\n", bp))
	}

	return sb.String()
}

func ParseFactory(fileName string) Factory {
	lr := aoc2022.NewLineReader(fileName)

	defer aoc2022.Close(lr)

	result := NewFactory()

	for lr.HasNext() {
		result.Blueprints = append(result.Blueprints, ParseBlueprint(lr.Text()))
	}

	log.Printf("Parsed %v", result)

	return result
}

type State struct {
	Duration, Iteration                                int
	OreRobots, ClayRobots, ObsidianRobots, GeodeRobots int
	Ore, Clay, Obsidian, Geode                         int
}

func NewState(duration int) State {
	return State{Duration: duration, OreRobots: 1}
}

func (s State) Build(r Robot) State {
	result := s.Iterate()

	result.Ore -= r.OreCost
	result.Clay -= r.ClayCost
	result.Obsidian -= r.ObsidianCost

	switch r.Type {
	case Ore:
		result.OreRobots++
	case Clay:
		result.ClayRobots++
	case Obsidian:
		result.ObsidianRobots++
	case Geode:
		result.GeodeRobots++
	}

	return result
}

func (s State) Compare(other State) int {
	return s.Geode - other.Geode
}

func (s State) Children(bp Blueprint) []State {
	if bp.CanBuildGeodeRobot(s) {
		child := s.Build(bp.GeodeRobot)

		return []State{child}
	}

	result := make([]State, 0, 4)

	child := s.Iterate()
	if child.Valid(bp) {
		result = append(result, child)
	}

	if bp.CanBuildObsidianRobot(s) {
		child := s.Build(bp.ObsidianRobot)

		if child.Valid(bp) {
			result = append(result, child)
		}
	}

	if bp.CanBuildClayRobot(s) {
		child := s.Build(bp.ClayRobot)

		if child.Valid(bp) {
			result = append(result, child)
		}
	}

	if bp.CanBuildOreRobot(s) {
		child := s.Build(bp.OreRobot)

		if child.Valid(bp) {
			result = append(result, child)
		}
	}

	return result
}

func (s State) Iterate() State {
	result := s

	result.Iteration++

	result.Ore += result.OreRobots
	result.Clay += result.ClayRobots
	result.Obsidian += result.ObsidianRobots
	result.Geode += result.GeodeRobots

	return result
}

func (s State) String() string {
	return fmt.Sprintf("Iteration %d: %d ore robots, %d clay robots, %d obsidian robots, %d geode robots. %d ore, %d clay, %d obsidian, %d geode.",
		s.Iteration,
		s.OreRobots,
		s.ClayRobots,
		s.ObsidianRobots,
		s.GeodeRobots,
		s.Ore,
		s.Clay,
		s.Obsidian,
		s.Geode,
	)
}

func (s State) Valid(bp Blueprint) bool {
	if s.OreRobots > bp.MaxOre || s.ClayRobots > bp.MaxClay || s.ObsidianRobots > bp.MaxObsidian {
		return false
	}

	if max, exists := iterationCache[s.Iteration]; exists {
		if s.Geode > max {
			iterationCache[s.Iteration] = s.Geode

			return true
		}

		return s.Geode-max >= -1
	} else {
		iterationCache[s.Iteration] = s.Geode
	}

	return true
}
