package main

import (
	"log"
)

func part1() {
	m := ParseMap("input.txt")

	log.Print(m)

	best := m.Path(Path{positions: []Position{m.Start}})

	log.Printf("Best %v", best.String(m)) // 394
}

func part2() {
	m := ParseMap("input.txt")
	best := Path{}

	for y, row := range m.heights {
		for x, h := range row {
			if h == 0 {
				p := m.Path(Path{positions: []Position{{
					x: x,
					y: y,
				}}})

				if best.Len() == 0 || p.Len() != 0 && p.Len() < best.Len() {
					best = p
				}
			}
		}
	}

	log.Printf("Best %v", best.String(m)) // 394
}

func main() {
	part1()
	part2()
}
