package main

import (
	"fmt"
	"log"
)

type Cube struct {
	x, y, z int
}

func (c Cube) Adjacent() []Cube {
	return []Cube{
		{x: c.x - 1, y: c.y, z: c.z},
		{x: c.x + 1, y: c.y, z: c.z},
		{x: c.x, y: c.y - 1, z: c.z},
		{x: c.x, y: c.y + 1, z: c.z},
		{x: c.x, y: c.y, z: c.z - 1},
		{x: c.x, y: c.y, z: c.z + 1},
	}
}

func (c Cube) Compare(other Cube) int {
	result := c.x - other.x

	if result == 0 {
		result = c.y - other.y

		if result == 0 {
			result = c.z - other.y
		}
	}

	return result
}

func (c Cube) String() string {
	return fmt.Sprintf("Cube (%v, %v, %v)", c.x, c.y, c.z)
}

func ParseCube(s string) Cube {
	result := Cube{}

	if _, err := fmt.Sscanf(s, "%d,%d,%d", &result.x, &result.y, &result.z); err != nil {
		log.Fatalf("Failed to parse cube %s", s)
	}

	return result
}
