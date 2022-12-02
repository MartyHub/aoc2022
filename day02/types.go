package main

import "log"

type Result int

const (
	Loose Result = 0
	Draw         = 3
	Win          = 6
)

func MustToResult(c rune) Result {
	switch {
	case c == 'X':
		return Loose
	case c == 'Y':
		return Draw
	case c == 'Z':
		return Win
	}

	log.Fatalf("Unknown result: %c", c)

	return -1
}

func (r Result) Play(other Shape) Shape {
	for _, s := range []Shape{Rock, Paper, Scissors} {
		if s.Play(other) == r {
			return s
		}
	}

	log.Fatalf("Failed to play against %v to %v", other, r)

	return -1
}

func (r Result) Score() int {
	return int(r)
}

func (r Result) String() string {
	switch r {
	case Loose:
		return "loose"
	case Draw:
		return "draw"
	case Win:
		return "win"
	}

	return "unknown"
}

type Shape int

const (
	Rock     Shape = 1
	Paper          = 2
	Scissors       = 3
)

func MustToShape(c rune) Shape {
	switch {
	case c == 'A' || c == 'X':
		return Rock
	case c == 'B' || c == 'Y':
		return Paper
	case c == 'C' || c == 'Z':
		return Scissors
	}

	log.Fatalf("Unknown shape: %c", c)

	return -1
}

func (s Shape) Play(other Shape) Result {
	if s == other {
		return Draw
	}

	if other == Rock && s == Paper || other == Paper && s == Scissors || other == Scissors && s == Rock {
		return Win
	}

	return Loose
}

func (s Shape) Score() int {
	return int(s)
}

func (s Shape) String() string {
	switch s {
	case Rock:
		return "rock"
	case Paper:
		return "paper"
	case Scissors:
		return "scissors"
	}

	return "unknown"
}
