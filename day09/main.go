package main

import (
	"aoc2022"
	"log"
)

func main() {
	lr := aoc2022.NewLineReader("input.txt")

	defer aoc2022.Close(lr)

	rope2 := NewRope(2)
	rope10 := NewRope(10)

	for lr.HasNext() {
		motion := ParseMotion(lr.Text())

		rope2.Move(motion)
		rope10.Move(motion)
	}

	log.Printf("Rope with  2 knots: tail visit %v position(s) at least once", len(rope2.tailPositions))  // 6 498
	log.Printf("Rope with 10 knots: tail visit %v position(s) at least once", len(rope10.tailPositions)) // 2 531
}
