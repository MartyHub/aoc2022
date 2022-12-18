package main

import "log"

func main() {
	chamber := NewChamber("input.txt")

	for chamber.iteration != 2022 {
		//for chamber.iteration != 10_000_000 {
		//for chamber.iteration != 1_000_000_000_000 {
		chamber.iterate()
	}

	log.Printf("Highest Rock: %v", chamber.height+chamber.highestRock()+1) // 3 191
}
