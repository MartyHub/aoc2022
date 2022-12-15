package main

import (
	"aoc2022"
	"fmt"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"
)

type Sensor struct {
	Distance      int
	Point         aoc2022.Point
	ClosestBeacon aoc2022.Point
}

func (s Sensor) Exclusion(y int) aoc2022.Interval {
	dy := aoc2022.Abs(s.Point.Y - y)

	if dy > s.Distance {
		return aoc2022.Interval{}
	}

	dx := s.Distance - dy

	return aoc2022.Interval{From: s.Point.X - dx, To: s.Point.X + dx + 1}
}

func (s Sensor) String() string {
	return fmt.Sprintf("Sensor @ %s, closest beason @ %s", s.Point, s.ClosestBeacon)
}

func ParseSensor(s string) Sensor {
	result := Sensor{}

	if _, err := fmt.Sscanf(
		s,
		"Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d",
		&result.Point.X,
		&result.Point.Y,
		&result.ClosestBeacon.X,
		&result.ClosestBeacon.Y,
	); err != nil {
		log.Fatalf("Failed to parse sensor %v: %v", s, err)
	}

	result.Distance = result.Point.Distance(result.ClosestBeacon)

	return result
}

type TileType string

const (
	Unknown    TileType = "."
	BeaconType TileType = "B"
	SensorType TileType = "S"
)

type Tile struct {
	Point aoc2022.Point
	Type  TileType
}

type Area struct {
	Sensors              []Sensor
	Tiles                []Tile
	TopLeft, BottomRight aoc2022.Point
}

func (a Area) NotBeacon(x, y int) bool {
	if a.Type(x, y) != Unknown {
		return false
	}

	point := aoc2022.Point{X: x, Y: y}

	for _, sensor := range a.Sensors {
		if sensor.Point.Distance(point) <= sensor.Distance {
			return true
		}
	}

	return false
}

func (a Area) Draw() string {
	yLen := aoc2022.Max(len(strconv.Itoa(a.TopLeft.Y)), len(strconv.Itoa(a.BottomRight.Y)))
	yFormat := fmt.Sprintf("%%%dd ", yLen)
	sb := strings.Builder{}

	for y := a.TopLeft.Y; y <= a.BottomRight.Y; y++ {
		sb.WriteString(fmt.Sprintf(yFormat, y))

		for x := a.TopLeft.X; x <= a.BottomRight.X; x++ {
			sb.WriteString(string(a.Type(x, y)))
		}

		sb.WriteString("\n")
	}

	return sb.String()
}

func (a Area) Find(x, y int) (int, bool) {
	return sort.Find(len(a.Tiles), func(i int) int {
		tile := a.Tiles[i]
		result := y - tile.Point.Y

		if result == 0 {
			result = x - tile.Point.X
		}

		return result
	})
}

func (a Area) Type(x, y int) TileType {
	i, found := a.Find(x, y)

	if found {
		return a.Tiles[i].Type
	}

	return Unknown
}

func (a Area) String() string {
	return fmt.Sprintf("Area: %v sensors in %v -> %v", len(a.Sensors), a.TopLeft, a.BottomRight)
}

func ParseArea(fileName string) Area {
	lr := aoc2022.NewLineReader(fileName)

	defer aoc2022.Close(lr)

	result := Area{
		TopLeft:     aoc2022.Point{X: math.MaxInt, Y: math.MaxInt},
		BottomRight: aoc2022.Point{X: math.MinInt, Y: math.MinInt},
	}

	for lr.HasNext() {
		sensor := ParseSensor(lr.Text())

		result.Sensors = append(result.Sensors, sensor)
		result.Tiles = append(result.Tiles, Tile{Point: sensor.Point, Type: SensorType})
		result.Tiles = append(result.Tiles, Tile{Point: sensor.ClosestBeacon, Type: BeaconType})

		result.TopLeft = result.TopLeft.Min(aoc2022.Point{X: sensor.Point.X - sensor.Distance, Y: sensor.Point.Y - sensor.Distance})
		result.BottomRight = result.BottomRight.Max(aoc2022.Point{X: sensor.Point.X + sensor.Distance, Y: sensor.Point.Y + sensor.Distance})
	}

	sort.Slice(result.Tiles, func(i, j int) bool {
		return result.Tiles[i].Point.Less(result.Tiles[j].Point)
	})

	return result
}
