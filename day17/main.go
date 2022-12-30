package main

import (
	"aoc2022"
	"log"
)

func main() {
	//NewChamber("test.txt").Iterate(10, true)

	for _, iteration := range []int{2022, 10_000_000} {
		log.Printf("Iteration # %s: height = %s",
			aoc2022.PrettyFormat(iteration),
			aoc2022.PrettyFormat(NewChamber("input.txt").Iterate(iteration, false)),
		)
	}

	// Iteration = 2 022-> Height = 3 191
}
