package main

import (
	"aoc2022"
	"log"
)

func main() {
	//NewChamber("test.txt").Iterate(10, true)

	for _, iteration := range []int{2022, 1000000000000} {
		log.Printf("Iteration # %s: height = %s",
			aoc2022.PrettyFormat(iteration),
			aoc2022.PrettyFormat(NewChamber("input.txt").Iterate(iteration, false)),
		)
	}

	// Iteration # 2 022: height = 3 191
	// Iteration # 1 000 000 000 000: height = 1 572 093 023 267
}
