package main

import "log"

func part1() {
	v := parseValley("input.txt")

	log.Printf("cycle: %d", v.cycle())

	minutes := v.iterate(state{v.blizzards, 0, firstStart, firstStart, v.target}, 9)

	log.Printf("Part 1: %d", minutes) // 232
}

func part2() {
	v := parseValley("input.txt")
	minutes1 := v.iterate(state{v.blizzards, 0, firstStart, firstStart, v.target}, 9)
	minutes2 := v.iterate(state{v.blizzardsAfter(minutes1), 0, v.target, v.target, firstStart}, 15)             // 255
	minutes3 := v.iterate(state{v.blizzardsAfter(minutes1 + minutes2), 0, firstStart, firstStart, v.target}, 5) // 228

	log.Printf("Part 2: %d + %d + %d = %d", minutes1, minutes2, minutes3, minutes1+minutes2+minutes3) // 715
}

func main() {
	//part1()
	part2()
}
