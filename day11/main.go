package main

import (
	"aoc2022"
	"log"
)

func parse() *Monkeys {
	lr := aoc2022.NewLineReader("input.txt")

	defer aoc2022.Close(lr)

	result := &Monkeys{}
	lines := make([]string, 0, 6)

	for lr.HasNext() {
		text := lr.Text()

		if text != "" {
			lines = append(lines, text)
		}

		if len(lines) == 6 {
			result.Add(ParseMonkey(lines))
			lines = lines[:0]
		}
	}

	return result
}

func part1() {
	monkeys := parse()

	for round := 1; round <= 20; round++ {
		monkeys.Inspect(true)

	}

	log.Printf("Business level: %v", monkeys.BusinessLevel()) // 57 348
}

func part2() {
	monkeys := parse()

	for round := 1; round <= 10_000; round++ {
		monkeys.Inspect(false)
	}

	log.Printf("Business level: %v", monkeys.BusinessLevel()) // 14 106 266 886
}

func main() {
	part1()
	part2()
}
