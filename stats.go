package aoc2022

import (
	"log"
	"time"
)

var durations = make(map[string]time.Duration)

func Print() {
	for k, v := range durations {
		log.Printf("%s took %v", k, v)
	}
}

func Timer(name string) func() {
	start := time.Now()

	return func() {
		duration := time.Since(start)

		if _, found := durations[name]; found {
			durations[name] += duration
		} else {
			durations[name] = duration
		}
	}
}
