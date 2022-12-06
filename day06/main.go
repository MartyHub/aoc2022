package main

import (
	"aoc2022"
	"log"
)

func uniq(r []rune, end, size int) bool {
	if end < size-1 {
		return false
	}

	for i := 0; i < size; i++ {
		start := end - size + 1 + i
		value := r[start]

		for j := start + 1; j <= end; j++ {
			if value == r[j] {
				return false
			}
		}
	}

	return true
}

func main() {
	lr := aoc2022.NewLineReader("input.txt")

	defer aoc2022.Close(lr)

	startOfPacket := -1
	startOfMessage := -1

	for lr.HasNext() {
		runes := []rune(lr.Text())

		for i := 0; i < len(runes); i++ {
			if startOfPacket == -1 && uniq(runes, i, 4) {
				startOfPacket = i + 1
			}

			if uniq(runes, i, 14) {
				startOfMessage = i + 1

				break
			}
		}
	}

	log.Printf("Start of Packet Marker: %v", startOfPacket)   // 1042
	log.Printf("Start of Message Marker: %v", startOfMessage) // 2980
}
