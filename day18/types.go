package main

import (
	"aoc2022"
	"fmt"
	"log"
	"math"
)

type Cube struct {
	x, y, z int
}

func (c Cube) Neighbors() []Cube {
	return []Cube{
		{x: c.x - 1, y: c.y, z: c.z},
		{x: c.x + 1, y: c.y, z: c.z},
		{x: c.x, y: c.y - 1, z: c.z},
		{x: c.x, y: c.y + 1, z: c.z},
		{x: c.x, y: c.y, z: c.z - 1},
		{x: c.x, y: c.y, z: c.z + 1},
	}
}

func (c Cube) String() string {
	return fmt.Sprintf("%v,%v,%v", c.x, c.y, c.z)
}

func ParseCube(s string) Cube {
	result := Cube{}

	if _, err := fmt.Sscanf(s, "%d,%d,%d", &result.x, &result.y, &result.z); err != nil {
		log.Fatalf("Failed to parse cube %s", s)
	}

	return result
}

type Lava struct {
	min, max Cube
	cubes    map[Cube]Cube
}

func NewLava() *Lava {
	return &Lava{
		cubes: make(map[Cube]Cube),
	}
}

func (l *Lava) AddCube(c Cube) {
	l.cubes[c] = c
}

func (l *Lava) Exists(c Cube) bool {
	_, result := l.cubes[c]

	return result
}

func (l *Lava) Surface() int {
	result := 0

	for c := range l.cubes {
		for _, ac := range c.Neighbors() {
			if !l.Exists(ac) {
				result++
			}
		}
	}

	return result
}

func (l *Lava) ExteriorSurface() int {
	result := 0
	queue := []Cube{{}}
	visited := make(map[Cube]bool)

	for len(queue) > 0 {
		c := queue[0]
		queue = queue[1:]

		if _, found := visited[c]; found {
			continue
		}

		visited[c] = true

		for _, ac := range c.Neighbors() {
			if ac.x >= l.min.x && ac.x <= l.max.x && ac.y >= l.min.y && ac.y <= l.max.y && ac.z >= l.min.z && ac.z <= l.max.z {
				if l.Exists(ac) {
					result++
				} else if _, found := visited[ac]; !found {
					queue = append(queue, ac)
				}
			}
		}
	}

	return result
}

func (l *Lava) Max() Cube {
	max := Cube{x: math.MinInt, y: math.MinInt, z: math.MinInt}

	for c := range l.cubes {
		if c.x > max.x {
			max.x = c.x
		}

		if c.y > max.y {
			max.y = c.y
		}

		if c.z > max.z {
			max.z = c.z
		}
	}

	return Cube{max.x + 1, max.y + 1, max.z + 1}
}

func (l *Lava) Min() Cube {
	min := Cube{x: math.MaxInt, y: math.MaxInt, z: math.MaxInt}

	for c := range l.cubes {
		if c.x < min.x {
			min.x = c.x
		}

		if c.y < min.y {
			min.y = c.y
		}

		if c.z < min.z {
			min.z = c.z
		}
	}

	return Cube{min.x - 1, min.y - 1, min.z - 1}
}

func ParseLava(fileName string) *Lava {
	lr := aoc2022.NewLineReader(fileName)

	defer aoc2022.Close(lr)

	result := NewLava()

	for lr.HasNext() {
		result.AddCube(ParseCube(lr.Text()))
	}

	result.max = result.Max()
	result.min = result.Min()

	fmt.Printf("%v cubes\n", len(result.cubes))

	return result
}
