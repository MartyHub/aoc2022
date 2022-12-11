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

type Monkey struct {
	Index      int
	Items      []int
	Operation  func(int) int
	Test       int
	IfTrue     int
	IfFalse    int
	Inspection int
}

func (m *Monkey) AddItem(item int) {
	m.Items = append(m.Items, item)
}

func (m *Monkey) ClearItems() {
	m.Items = []int{}
}

func (m *Monkey) Inspect(relief bool, magic int, item int) (int, int) {
	m.Inspection++

	worryLevel := m.Operation(item)

	if relief {
		worryLevel /= 3
	}

	worryLevel = worryLevel % magic

	if worryLevel%m.Test == 0 {
		return m.IfTrue, worryLevel
	} else {
		return m.IfFalse, worryLevel
	}
}

func (m *Monkey) String() string {
	return fmt.Sprintf(
		"Monkey # %d:\n- Items:     %v\n- Operation: %v\n- Test:      %v\n- If True:   %d\n- If False:  %d",
		m.Index,
		m.Items,
		m.Operation(1),
		m.Test,
		m.IfTrue,
		m.IfFalse,
	)
}

func ParseMonkey(lines []string) *Monkey {
	return &Monkey{
		Index:     parseIndex(lines[0]),
		Items:     parseItems(lines[1]),
		Operation: parseOperation(lines[2]),
		Test:      parseTest(lines[3]),
		IfTrue:    parseIfTrue(lines[4]),
		IfFalse:   parseIfFalse(lines[5]),
	}
}

func parseIndex(s string) int {
	return int(s[7] - '0')
}

func parseItems(s string) []int {
	items := strings.Split(s[18:], ", ")
	result := make([]int, len(items))

	for i, item := range items {
		result[i] = aoc2022.Must(strconv.Atoi(item))
	}

	return result
}

func parseOperation(s string) func(int) int {
	tokens := strings.Split(s[19:], " ")
	value := math.MinInt

	if tokens[2] != "old" {
		value = aoc2022.Must(strconv.Atoi(tokens[2]))
	}

	return func(old int) int {
		switch tokens[1] {
		case "+":
			if value == math.MinInt {
				return old * 2
			} else {
				return old + value
			}
		case "*":
			if value == math.MinInt {
				return old * old
			} else {
				return old * value
			}
		}

		log.Fatalf("Don't know how to handle operation %v", s)

		return 0
	}
}

func parseTest(s string) int {
	return aoc2022.Must(strconv.Atoi(s[21:]))
}

func parseIfTrue(s string) int {
	return aoc2022.Must(strconv.Atoi(s[29:]))
}

func parseIfFalse(s string) int {
	return aoc2022.Must(strconv.Atoi(s[30:]))
}

type Monkeys struct {
	monkeys []*Monkey
}

func (m *Monkeys) Add(monkey *Monkey) {
	m.monkeys = append(m.monkeys, monkey)
}

func (m *Monkeys) BusinessLevel() int {
	var inspections []int

	for _, monkey := range m.monkeys {
		inspections = append(inspections, monkey.Inspection)
	}

	sort.Ints(inspections)

	return inspections[len(inspections)-1] * inspections[len(inspections)-2]
}

func (m *Monkeys) Inspect(relief bool) {
	var tests []int

	for _, monkey := range m.monkeys {
		tests = append(tests, monkey.Test)
	}

	lcm := LCM(tests[0], tests[1], tests[2:]...)

	for _, monkey := range m.monkeys {
		for _, item := range monkey.Items {
			i, worryLevel := monkey.Inspect(relief, lcm, item)

			m.monkeys[i].AddItem(worryLevel)
		}

		monkey.ClearItems()
	}
}

func (m *Monkeys) String() string {
	result := strings.Builder{}

	for _, monkey := range m.monkeys {
		result.WriteString(monkey.String())
		result.WriteString("\n")
	}

	return result.String()
}
