package main

import (
	"log"
)

func part1() {
	grove := parseGrove("input.txt")

	for i := 0; i < 10; i++ {
		grove.round()
	}

	log.Printf("Empty ground tiles after 10 rounds: %d", grove.count()) // 4 172
}

func part2() {
	grove := parseGrove("input.txt")
	round := 1

	for !grove.round() {
		round++
	}

	log.Printf("First round without move: %d", round) // 942
}

func main() {
	part1()
	part2()
}
