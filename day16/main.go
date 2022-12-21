package main

import (
	"log"
)

func main() {
	cave := ParseCave("input.txt", 30)

	log.Printf("ReleaseMax: %v", cave.ReleaseMax(Single())) // 1 580
}
