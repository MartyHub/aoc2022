package main

import (
	"aoc2022"
	"fmt"
	"log"
	"regexp"
	"strconv"
)

type Section int

type Assignment struct {
	Start Section
	End   Section
}

func (a Assignment) contains(other Assignment) bool {
	return a.Start <= other.Start && a.End >= other.End
}

func (a Assignment) containsSection(section Section) bool {
	return a.Start <= section && section <= a.End
}

func (a Assignment) overlap(other Assignment) bool {
	return a.containsSection(other.Start) || a.containsSection(other.End) || other.containsSection(a.Start) || other.containsSection(a.End)
}

func (a Assignment) String() string {
	return fmt.Sprintf("[%2v, %2v]", a.Start, a.End)
}

type Pair struct {
	First  Assignment
	Second Assignment
}

func (p Pair) String() string {
	return fmt.Sprintf("(%v, %v)", p.First, p.Second)
}

func (p Pair) FullOverlap() bool {
	return p.First.contains(p.Second) || p.Second.contains(p.First)
}

func (p Pair) Overlap() bool {
	return p.First.overlap(p.Second)
}

var lineRegex = regexp.MustCompile(`(\d+)-(\d+),(\d+)-(\d+)`)

func Parse(s string) Pair {
	tokens := lineRegex.FindStringSubmatch(s)

	if len(tokens) != 5 {
		log.Fatalf("Failed to parse pair: %v", s)
	}

	return Pair{
		First: Assignment{
			Start: Section(aoc2022.Must(strconv.Atoi(tokens[1]))),
			End:   Section(aoc2022.Must(strconv.Atoi(tokens[2]))),
		},
		Second: Assignment{
			Start: Section(aoc2022.Must(strconv.Atoi(tokens[3]))),
			End:   Section(aoc2022.Must(strconv.Atoi(tokens[4]))),
		},
	}
}
