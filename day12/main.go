package main

import (
	"log"
)

func part1() {
	m := ParseMap("input.txt")

	best := m.Path(Path{positions: []Position{m.Start}})

	log.Printf("Best: %v", best.Len()-1) // 394
}

func part2() {
	m := ParseMap("input.txt")
	best := m.PathToHeight(Path{positions: []Position{m.Target}}, 0)

	log.Printf("Best %v", best.Len()-1) // 388
}

func main() {
	part1()
	part2()
}
