package main

import (
	"aoc2022"
	"log"
)

func findDuplicate(runes []rune) Supply {
	mid := len(runes) / 2

	for i := 0; i < mid; i++ {
		for j := 0; j < mid; j++ {
			if runes[i] == runes[mid+j] {
				return Supply(runes[i])
			}
		}
	}

	log.Fatalf("Failed to find duplicate in %v", string(runes))

	return -1
}

func findBadge(rucksacks [3][]rune) Supply {
	for i := 0; i < len(rucksacks[0]); i++ {
		for j := 0; j < len(rucksacks[1]); j++ {
			for k := 0; k < len(rucksacks[2]); k++ {
				if rucksacks[0][i] == rucksacks[1][j] && rucksacks[1][j] == rucksacks[2][k] {
					return Supply(rucksacks[0][i])
				}
			}
		}
	}

	log.Fatalf("Failed to find badge in %v", rucksacks)

	return -1
}

func main() {
	lr := aoc2022.NewLineReader("input.txt")

	defer aoc2022.Close(lr)

	badges := 0
	duplicates := 0
	rucksacks := [3][]rune{}

	for lr.HasNext() {
		runes := []rune(lr.Text())
		duplicates += findDuplicate(runes).Priority()

		i := (lr.Count() - 1) % 3
		rucksacks[i] = runes

		if i == 2 {
			badges += findBadge(rucksacks).Priority()
		}
	}

	log.Printf("Duplicates: %v", aoc2022.PrettyFormat(duplicates)) // 7 727
	log.Printf("Badges: %v", aoc2022.PrettyFormat(badges))         // 2 609
}
