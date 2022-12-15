package main

import (
	"aoc2022"
	"log"
)

func main() {
	//fileName := "test.txt"
	//y := 10

	fileName := "input.txt"
	y := 2000000

	area := ParseArea(fileName)

	log.Printf("%v", area)

	positions := 0

	for x := area.TopLeft.X; x <= area.BottomRight.X; x++ {
		if area.NotBeacon(x, y) {
			positions++
		}
	}

	log.Printf("Positions: %v", positions) // 4 876 693

	beacon := aoc2022.Point{}
	size := 4000000

	for y := 0; y <= size; y++ {
		exclusions := make([]aoc2022.Interval, 0)

		for _, sensor := range area.Sensors {
			exclusions = append(exclusions, sensor.Exclusion(y))
		}

		x := 0

		for x <= size {
			exclude := false

			for _, e := range exclusions {
				if e.Contains(x) {
					x = e.To
					exclude = true

					break
				}
			}

			if !exclude {
				beacon = aoc2022.Point{X: x, Y: y}

				break
			}
		}
	}

	log.Printf("Beacon: %v", beacon)
	log.Printf("Tuning frequency: %v", beacon.X*size+beacon.Y)
}
