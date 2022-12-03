package main

import (
	"aoc2022"
	"log"
)

func main() {
	lr := aoc2022.NewLineReader("input.txt")

	defer aoc2022.Close(lr)

	score1 := 0
	score2 := 0

	for lr.HasNext() {
		text := lr.Text()

		if len(text) != 3 {
			log.Fatalf("Invalid line %v: %v", lr.Count(), text)
		}

		runes := []rune(text)
		other := MustToShape(runes[0])
		you := MustToShape(runes[2])
		result := MustToResult(runes[2])

		score1 += you.Score() + you.Play(other).Score()
		score2 += result.Score() + result.Play(other).Score()
	}

	log.Printf("Score 1: %v", aoc2022.PrettyFormat(score1)) // 15 572
	log.Printf("Score 2: %v", aoc2022.PrettyFormat(score2)) // 16 098
}
