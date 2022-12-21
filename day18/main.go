package main

import (
	"fmt"
)

func main() {
	lava := ParseLava("input.txt")

	fmt.Printf("Surface: %v\n", lava.Surface())                  // 4 390
	fmt.Printf("Exterior Surface: %v\n", lava.ExteriorSurface()) // 2 534
}
