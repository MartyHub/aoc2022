package main

import (
	"aoc2022"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var source = Point{
	X: 500,
	Y: 0,
}

type Point struct {
	X, Y int
}

func (p Point) Delta(other Point) Point {
	return Point{
		X: p.X + p.DeltaX(other),
		Y: p.Y + p.DeltaY(other),
	}
}

func (p Point) DeltaX(other Point) int {
	if p.X < other.X {
		return 1
	} else if p.X > other.X {
		return -1
	}

	return 0
}

func (p Point) DeltaY(other Point) int {
	if p.Y < other.Y {
		return 1
	} else if p.Y > other.Y {
		return -1
	}

	return 0
}

func (p Point) Max(other Point) Point {
	return Point{
		X: aoc2022.Max(p.X, other.X),
		Y: aoc2022.Max(p.Y, other.Y),
	}
}

func (p Point) Min(other Point) Point {
	return Point{
		X: aoc2022.Min(p.X, other.X),
		Y: aoc2022.Min(p.Y, other.Y),
	}
}

func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

type TileType int

const (
	Air TileType = iota
	Rock
	Sand
	Source
)

func (tt TileType) String() string {
	switch tt {
	case Air:
		return "."
	case Rock:
		return "#"
	case Sand:
		return "o"
	case Source:
		return "+"
	default:
		return "?"
	}
}

type Tile struct {
	Pos  Point
	Type TileType
}

func ParseTiles(text string) []Tile {
	result := make([]Tile, 0)

	for _, s := range strings.Split(text, " -> ") {
		parts := strings.Split(s, ",")
		pos := Point{
			X: aoc2022.Must(strconv.Atoi(parts[0])),
			Y: aoc2022.Must(strconv.Atoi(parts[1])),
		}
		tile := Tile{Pos: pos, Type: Rock}

		if len(result) > 0 {
			for step := result[len(result)-1]; step != tile; {
				step = Tile{
					Pos:  step.Pos.Delta(tile.Pos),
					Type: Rock,
				}

				result = append(result, step)
			}
		} else {
			result = append(result, tile)
		}
	}

	return result
}

func (t Tile) String() string {
	return fmt.Sprintf("%v: %v", t.Pos, t.Type)
}

type Cave struct {
	Floor                bool
	Tiles                []Tile
	TopLeft, BottomRight Point
}

func ParseCave(fileName string, floor bool) *Cave {
	lr := aoc2022.NewLineReader(fileName)

	defer aoc2022.Close(lr)

	result := &Cave{
		Floor:       floor,
		Tiles:       []Tile{{Pos: source, Type: Source}},
		TopLeft:     Point{X: 500, Y: 0},
		BottomRight: Point{X: 500, Y: 0},
	}

	for lr.HasNext() {
		text := lr.Text()

		for _, tile := range ParseTiles(text) {
			result.Tiles = append(result.Tiles, tile)
			result.TopLeft = result.TopLeft.Min(tile.Pos)
			result.BottomRight = result.BottomRight.Max(tile.Pos)
		}
	}

	sort.Slice(result.Tiles, func(i, j int) bool {
		return result.Tiles[i].Pos.Y < result.Tiles[j].Pos.Y ||
			result.Tiles[i].Pos.Y == result.Tiles[j].Pos.Y &&
				result.Tiles[i].Pos.X < result.Tiles[j].Pos.X
	})

	return result
}

func (c *Cave) Draw() string {
	sb := strings.Builder{}

	for y := c.TopLeft.Y; y <= c.BottomRight.Y; y++ {
		for x := c.TopLeft.X; x <= c.BottomRight.X; x++ {
			sb.WriteString(c.Get(x, y).String())
		}

		sb.WriteString("\n")
	}

	return sb.String()
}

func (c *Cave) Find(x, y int) (int, bool) {
	return sort.Find(len(c.Tiles), func(i int) int {
		tile := c.Tiles[i]
		result := y - tile.Pos.Y

		if result == 0 {
			result = x - tile.Pos.X
		}

		return result
	})
}

func (c *Cave) Get(x, y int) TileType {
	if y == c.BottomRight.Y+2 {
		return Rock
	}

	i, found := c.Find(x, y)

	if found {
		return c.Tiles[i].Type
	}

	return Air
}

func (c *Cave) Produce() bool {
	if c.Get(source.X, source.Y) == Sand {
		return false
	}

	pos := source

	for {
		if c.Get(pos.X, pos.Y+1) == Air {
			pos.Y++
		} else if c.Get(pos.X-1, pos.Y+1) == Air {
			pos.X--
			pos.Y++
		} else if c.Get(pos.X+1, pos.Y+1) == Air {
			pos.X++
			pos.Y++
		} else {
			break
		}

		if !c.Floor && (pos.X < c.TopLeft.X || pos.X > c.BottomRight.X || pos.Y > c.BottomRight.Y) {
			return false
		}
	}

	c.Set(pos, Sand)

	return true
}

func (c *Cave) Set(pos Point, tt TileType) {
	i, found := c.Find(pos.X, pos.Y)

	if found {
		c.Tiles[i].Type = tt
	} else {
		c.Tiles = append(c.Tiles[:i+1], c.Tiles[i:]...)
		c.Tiles[i] = Tile{Pos: pos, Type: tt}
	}
}

func (c *Cave) String() string {
	return fmt.Sprintf("Cave: %v -> %v:\n%v", c.TopLeft, c.BottomRight, c.Draw())
}
