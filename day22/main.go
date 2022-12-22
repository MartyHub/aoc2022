package main

import (
	"log"
)

func part1() {
	notes := ParseNotes("input.txt", false)

	log.Printf("Password 1: %v", notes.Password1()) // 88 226
}

func part2() {
	notes := ParseNotes("test.txt", true)

	log.Printf("Password 1: %v", notes.Password1())
}

func main() {
	part1()
	part2()
}
