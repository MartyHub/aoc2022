package main

import (
	"aoc2022"
	"log"
)

func main() {
	lr := aoc2022.NewLineReader("input.txt")

	defer aoc2022.Close(lr)

	sum := 0

	for lr.HasNext() {
		sum += toDecimal(lr.Text())
	}

	log.Printf("sum: %d = %s", sum, toSnafu(sum)) // 20=212=1-12=200=00-1
}
