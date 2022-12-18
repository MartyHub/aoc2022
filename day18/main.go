package main

import (
	"aoc2022"
	"log"
)

func main() {
	lr := aoc2022.NewLineReader("input.txt")

	defer aoc2022.Close(lr)

	cubes := make(map[Cube]Cube, 0)

	for lr.HasNext() {
		cube := ParseCube(lr.Text())

		cubes[cube] = cube
	}

	log.Printf("%v cubes", len(cubes))

	surface := 0

	for _, c := range cubes {
		for _, ac := range c.Adjacent() {
			if _, exist := cubes[ac]; !exist {
				surface++
			}
		}
	}

	log.Printf("Surface: %v", surface) // 4 390
}
