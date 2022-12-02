package main

import (
	"aoc2022"
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer aoc2022.Close(file)

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	line := 0
	score1 := 0
	score2 := 0

	for scanner.Scan() {
		text := scanner.Text()
		line += 1

		if len(text) != 3 {
			log.Fatalf("Invalid line %v: %v", line, text)
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
