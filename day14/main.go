package main

import "log"

func part1(fileName string) {
	cave := ParseCave(fileName, false)
	unit := 0

	for ; cave.Produce(); unit++ {
	}

	log.Printf("After %v units without floor:\n%v", unit, cave) // 964
}

func part2(fileName string) {
	cave := ParseCave(fileName, true)
	unit := 0

	for ; cave.Produce(); unit++ {
	}

	log.Printf("After %v units with floor:\n%v", unit, cave) // 32 041
}

func main() {
	fileName := "input.txt"

	part1(fileName)
	part2(fileName)
}
