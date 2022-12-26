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

//func part2() {
//	cave := ParseCave("input.txt", 20)
//	mp := cave.Iterate2(false)
//
//	log.Printf("Max Pressure %d: %v", mp.Sum(), mp)
//}

func main() {
	part1()
	//part2()
}
