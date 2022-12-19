package main

import (
	"log"
)

func part1(factory Factory) {
	quality := 0

	for _, blueprint := range factory.Blueprints {
		ResetCaches()

		geode := blueprint.Run(NewState(24))

		log.Printf("Blueprint # %d: %d => %d", blueprint.Id, geode, blueprint.Id*geode)

		quality += blueprint.Id * geode
	}

	log.Printf("Quality: %d", quality) // 1 306
}

func part2(factory Factory) {
	result := 1

	for i, blueprint := range factory.Blueprints {
		if i == 3 {
			break
		}

		ResetCaches()

		geode := blueprint.Run(NewState(32))

		log.Printf("Blueprint # %d: %d", blueprint.Id, geode)

		result *= geode
	}

	log.Printf("Geodes: %d", result) // 37 604
}

func main() {
	factory := ParseFactory("input.txt")

	part1(factory)
	part2(factory)
}
