package main

import (
	"aoc2022"
	"log"
)

func main() {
	lr := aoc2022.NewLineReader("input.txt")

	defer aoc2022.Close(lr)

	fullOverlaps := 0
	overlaps := 0

	for lr.HasNext() {
		pair := Parse(lr.Text())

		if pair.FullOverlap() {
			fullOverlaps += 1
		}

		if pair.Overlap() {
			overlaps += 1
		}
	}

	log.Printf("Full Overlaps: %v", aoc2022.PrettyFormat(fullOverlaps)) // 573
	log.Printf("Overlaps: %v", aoc2022.PrettyFormat(overlaps))          // 867
}
