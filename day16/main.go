package main

import (
	"log"
	"time"
)

func part1() {
	cave := ParseCave("input.txt")
	start := time.Now()
	state := cave.Iterate(30)

	log.Printf("Most Pressure %d in %v:\n%s", state.TotalPressure(), time.Since(start), state.Debug(cave)) // 1 580
}

func part2() {
	cave := ParseCave("input.txt")
	start := time.Now()
	pair := cave.Iterate2(26)

	log.Printf("Most Pressure %d in %v:\n%s\n%s",
		pair.First.TotalPressure()+pair.Second.TotalPressure(),
		time.Since(start),
		pair.First.Debug(cave),
		pair.Second.Debug(cave),
	) // 2 213
}

func main() {
	part1()
	part2()
}
